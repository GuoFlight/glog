package glog

import (
	"os"
	"path"
	"runtime"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func NewLogger(logDir, logLevel string, AllLevelReportCaller bool, logFileCount uint) (*logrus.Logger, error) {
	Logger := logrus.New()

	//设置日志级别
	logLevelLogrus, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return nil, err
	}
	Logger.SetLevel(logLevelLogrus)

	//输出Caller
	if AllLevelReportCaller {
		//所有日志等级输出Caller
		Logger.SetReportCaller(true)
	} else {
		//指定日志等级输出Caller
		var printFileAndNumHook PrintFileAndNumHook
		Logger.AddHook(&printFileAndNumHook)
	}

	//取消日志标准输出(终端输出)
	if logLevelLogrus != logrus.DebugLevel {
		nullFile, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			return nil, err
		}
		Logger.SetOutput(nullFile) //将日志输出到此文件
	}

	//多文件输出+日志切割
	logFileTrace, err := GetWriter(path.Join(logDir, "TRACE.log"), logFileCount)
	if err != nil {
		return nil, err
	}
	logFileDebug, err := GetWriter(path.Join(logDir, "DEBUG.log"), logFileCount)
	if err != nil {
		return nil, err
	}
	logFileInfo, err := GetWriter(path.Join(logDir, "INFO.log"), logFileCount)
	if err != nil {
		return nil, err
	}
	logFileWarn, err := GetWriter(path.Join(logDir, "WARN.log"), logFileCount)
	if err != nil {
		return nil, err
	}
	logFileError, err := GetWriter(path.Join(logDir, "ERROR.log"), logFileCount)
	if err != nil {
		return nil, err
	}
	logFileFatal, err := GetWriter(path.Join(logDir, "FATAL.log"), logFileCount)
	if err != nil {
		return nil, err
	}
	logFilePanic, err := GetWriter(path.Join(logDir, "PANIC.log"), logFileCount)
	if err != nil {
		return nil, err
	}
	pathMap := lfshook.WriterMap{
		logrus.TraceLevel: logFileTrace,
		logrus.DebugLevel: logFileDebug,
		logrus.InfoLevel:  logFileInfo,
		logrus.WarnLevel:  logFileWarn,
		logrus.ErrorLevel: logFileError,
		logrus.FatalLevel: logFileFatal,
		logrus.PanicLevel: logFilePanic,
	}
	Logger.AddHook(lfshook.NewHook(
		pathMap,
		&logrus.TextFormatter{}, //普通文本模式
	))

	return Logger, nil
}
func NewLoggerWithConf(conf Conf) (*logrus.Logger, error) {
	Logger := logrus.New()

	//设置日志级别
	Logger.SetLevel(logrus.Level(conf.LogLevel))

	//指定日志等级输出Caller
	var printFileAndNumHook PrintFileAndNumHook
	var levels []logrus.Level
	for _, v := range conf.CallerPrintLevels {
		levels = append(levels, logrus.Level(v))
	}
	printFileAndNumHook.CallerPrintLevels = levels
	Logger.AddHook(&printFileAndNumHook)

	//取消日志标准输出(终端输出)
	if !conf.IsStdoutPrint {
		nullFile, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			return nil, err
		}
		Logger.SetOutput(nullFile) //将日志输出到此文件
	}

	//多文件输出+日志切割
	logFileTrace, err := GetWriter(path.Join(conf.LogDir, "TRACE.log"), conf.LogFileCount)
	if err != nil {
		return nil, err
	}
	logFileDebug, err := GetWriter(path.Join(conf.LogDir, "DEBUG.log"), conf.LogFileCount)
	if err != nil {
		return nil, err
	}
	logFileInfo, err := GetWriter(path.Join(conf.LogDir, "INFO.log"), conf.LogFileCount)
	if err != nil {
		return nil, err
	}
	logFileWarn, err := GetWriter(path.Join(conf.LogDir, "WARN.log"), conf.LogFileCount)
	if err != nil {
		return nil, err
	}
	logFileError, err := GetWriter(path.Join(conf.LogDir, "ERROR.log"), conf.LogFileCount)
	if err != nil {
		return nil, err
	}
	logFileFatal, err := GetWriter(path.Join(conf.LogDir, "FATAL.log"), conf.LogFileCount)
	if err != nil {
		return nil, err
	}
	logFilePanic, err := GetWriter(path.Join(conf.LogDir, "PANIC.log"), conf.LogFileCount)
	if err != nil {
		return nil, err
	}
	pathMap := lfshook.WriterMap{
		logrus.TraceLevel: logFileTrace,
		logrus.DebugLevel: logFileDebug,
		logrus.InfoLevel:  logFileInfo,
		logrus.WarnLevel:  logFileWarn,
		logrus.ErrorLevel: logFileError,
		logrus.FatalLevel: logFileFatal,
		logrus.PanicLevel: logFilePanic,
	}
	Logger.AddHook(lfshook.NewHook(
		pathMap,
		&logrus.TextFormatter{}, //普通文本模式
	))
	return Logger, nil
}

//指定日志等级输出文件名+行号的hook
type PrintFileAndNumHook struct {
	CallerPrintLevels []logrus.Level
}

func (hook *PrintFileAndNumHook) Fire(entry *logrus.Entry) error {
	pcs := make([]uintptr, 10)
	i := runtime.Callers(0, pcs)
	entry.Data["func"] = runtime.FuncForPC(pcs[i-2]).Name()
	file, line := runtime.FuncForPC(pcs[i-2]).FileLine(pcs[i-2])
	entry.Data["filename"] = file
	entry.Data["lineNum"] = line
	return nil
}
func (hook *PrintFileAndNumHook) Levels() []logrus.Level {
	return hook.CallerPrintLevels //在这些Levels上生效
}

//得到日志切割的输出对象
func GetWriter(pathLogFile string, logFileCount uint) (*rotatelogs.RotateLogs, error) {
	writer, err := rotatelogs.New(
		pathLogFile+".%Y%m%d%H",                                 //日志文件后缀：年月日时
		rotatelogs.WithLinkName(pathLogFile),                    //为当前正在输出的日志文件建立软连接
		rotatelogs.WithRotationCount(logFileCount),              //日志文件保存的个数(包括当前正在输出的日志)
		rotatelogs.WithRotationTime(time.Duration(1)*time.Hour), //设置日志分割的时间(隔多久分割一次)
	)
	if err != nil {
		return nil, err
	}
	return writer, nil
}

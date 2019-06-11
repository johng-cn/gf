<<<<<<< HEAD
// Copyright 2017 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.

// 日志模块.
// 直接文件/输出操作，没有异步逻辑，没有使用缓存或者通道
package glog

import (
    "io"
    "sync"
    "gitee.com/johng/gf/g/container/gtype"
)

type Logger struct {
    mu     sync.RWMutex
    io     io.Writer           // 日志内容写入的IO接口
    path   *gtype.String       // 日志写入的目录路径
    debug  *gtype.Bool         // 是否允许输出DEBUG信息
    btSkip *gtype.Int          // 错误产生时的backtrace回调信息skip条数
}

// 默认的日志对象
var logger = New()

// 新建自定义的日志操作对象
func New() *Logger {
    return &Logger {
        io     : nil,
        path   : gtype.NewString(),
        debug  : gtype.NewBool(true),
        btSkip : gtype.NewInt(3),
    }
}

// 日志日志目录绝对路径
=======
// Copyright 2017 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
//
// @author john, zseeker

// Package glog implements powerful and easy-to-use levelled logging functionality.
package glog

import (
	"github.com/gogf/gf/g/internal/cmdenv"
	"github.com/gogf/gf/g/os/grpool"
	"io"
)

const (
    LEVEL_ALL  = LEVEL_DEBU | LEVEL_INFO | LEVEL_NOTI | LEVEL_WARN | LEVEL_ERRO | LEVEL_CRIT
    LEVEL_DEV  = LEVEL_ALL
    LEVEL_PROD = LEVEL_WARN | LEVEL_ERRO | LEVEL_CRIT
    LEVEL_DEBU = 1 << iota
    LEVEL_INFO
    LEVEL_NOTI
    LEVEL_WARN
    LEVEL_ERRO
    LEVEL_CRIT
)

var (
    // Default logger object, for package method usage
    logger = New()
    // Goroutine pool for async logging output.
	asyncPool = grpool.New(1)
)

func init() {
    SetDebug(cmdenv.Get("gf.glog.debug", true).Bool())
}

// SetPath sets the directory path for file logging.
>>>>>>> upstream/master
func SetPath(path string) {
    logger.SetPath(path)
}

<<<<<<< HEAD
// 设置是否允许输出DEBUG信息
func SetDebug(debug bool) {
    logger.SetDebug(debug)
}

// 获取日志目录绝对路径
func GetPath() string {
    return logger.path.Val()
}

func Print(v ...interface{}) {
    logger.Print(v ...)
}

func Printf(format string, v ...interface{}) {
    logger.Printf(format, v ...)
}

func Println(v ...interface{}) {
    logger.Println(v ...)
}

func Printfln(format string, v ...interface{}) {
    logger.Printfln(format, v ...)
}

func Fatal(v ...interface{}) {
    logger.Fatal(v ...)
}

func Fatalf(format string, v ...interface{}) {
    logger.Fatalf(format, v ...)
}

func Fatalln(v ...interface{}) {
    logger.Fatalln(v ...)
}

func Fatalfln(format string, v ...interface{}) {
    logger.Fatalfln(format, v ...)
}

func Panic(v ...interface{}) {
    logger.Panic(v ...)
}

func Panicf(format string, v ...interface{}) {
    logger.Panicf(format, v ...)
}

func Panicln(v ...interface{}) {
    logger.Panicln(v ...)
}

func Panicfln(format string, v ...interface{}) {
    logger.Panicfln(format, v ...)
}

func Info(v ...interface{}) {
    logger.Info(v...)
}

func Debug(v ...interface{}) {
    logger.Debug(v...)
}

func Notice(v ...interface{}) {
    logger.Notice(v...)
}

func Warning(v ...interface{}) {
    logger.Warning(v...)
}

func Error(v ...interface{}) {
    logger.Error(v...)
}

func Critical(v ...interface{}) {
    logger.Critical(v...)
}

func Infof(format string, v ...interface{}) {
    logger.Infof(format, v...)
}

func Debugf(format string, v ...interface{}) {
    logger.Debugf(format, v...)
}

func Noticef(format string, v ...interface{}) {
    logger.Noticef(format, v...)
}

func Warningf(format string, v ...interface{}) {
    logger.Warningf(format, v...)
}

func Errorf(format string, v ...interface{}) {
    logger.Errorf(format, v...)
}

func Criticalf(format string, v ...interface{}) {
    logger.Criticalf(format, v...)
}

func Infofln(format string, v ...interface{}) {
    logger.Infofln(format, v...)
}

func Debugfln(format string, v ...interface{}) {
    logger.Debugfln(format, v...)
}

func Noticefln(format string, v ...interface{}) {
    logger.Noticefln(format, v...)
}

func Warningfln(format string, v ...interface{}) {
    logger.Warningfln(format, v...)
}

func Errorfln(format string, v ...interface{}) {
    logger.Errorfln(format, v...)
}

func Criticalfln(format string, v ...interface{}) {
    logger.Criticalfln(format, v...)
=======
// GetPath returns the logging directory path for file logging.
// It returns empty string if no directory path set.
func GetPath() string {
	return logger.GetPath()
}

// SetFile sets the file name <pattern> for file logging.
// Datetime pattern can be used in <pattern>, eg: access-{Ymd}.log.
// The default file name pattern is: Y-m-d.log, eg: 2018-01-01.log
func SetFile(pattern string) {
    logger.SetFile(pattern)
}

// SetLevel sets the default logging level.
func SetLevel(level int) {
    logger.SetLevel(level)
}

// GetLevel returns the default logging level value.
func GetLevel() int {
	return logger.GetLevel()
}

// SetWriter sets the customized logging <writer> for logging.
// The <writer> object should implements the io.Writer interface.
// Developer can use customized logging <writer> to redirect logging output to another service, 
// eg: kafka, mysql, mongodb, etc.
func SetWriter(writer io.Writer) {
    logger.SetWriter(writer)
}

// GetWriter returns the customized writer object, which implements the io.Writer interface.
// It returns nil if no customized writer set.
func GetWriter() io.Writer {
    return logger.GetWriter()
}

// SetDebug enables/disables the debug level for default logger.
// The debug level is enbaled in default.
func SetDebug(debug bool) {
    logger.SetDebug(debug)
}

// SetAsync enables/disables async logging output feature for default logger.
func SetAsync(enabled bool) {
	logger.SetAsync(enabled)
}

// SetStdoutPrint sets whether ouptput the logging contents to stdout, which is false in default.
func SetStdoutPrint(enabled bool) {
    logger.SetStdoutPrint(enabled)
}

// SetHeaderPrint sets whether output header of the logging contents, which is true in default.
func SetHeaderPrint(enabled bool) {
	logger.SetHeaderPrint(enabled)
}

// SetPrefix sets prefix string for every logging content.
// Prefix is part of header, which means if header output is shut, no prefix will be output.
func SetPrefix(prefix string) {
	logger.SetPrefix(prefix)
}

// SetFlags sets extra flags for logging output features.
func SetFlags(flags int) {
	logger.SetFlags(flags)
}

// GetFlags returns the flags of logger.
func GetFlags() int {
	return logger.GetFlags()
}

// PrintBacktrace prints the caller backtrace, 
// the optional parameter <skip> specify the skipped backtrace offset from the end point.
func PrintBacktrace(skip...int) {
    logger.PrintBacktrace(skip...)
}

// GetBacktrace returns the caller backtrace content, 
// the optional parameter <skip> specify the skipped backtrace offset from the end point.
func GetBacktrace(skip...int) string {
    return logger.GetBacktrace(skip...)
}

// SetBacktrace enables/disables the backtrace feature in failure logging outputs.
func SetBacktrace(enabled bool) {
    logger.SetBacktrace(enabled)
>>>>>>> upstream/master
}

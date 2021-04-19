package logger

import (
	"runtime"
)

func (l *Logger) caller() {
	pc,_,line,_ := runtime.Caller(2)
	pcName := runtime.FuncForPC(pc).Name() //获取函数名
	l.Caller = make(map[string]interface{})
	l.Caller["p-caller"] = pcName
	l.Caller["p-line"] = line
	l.Caller["trace_id"] = l.TraceId
}

func (l *Logger) Debug(format string){
	l.caller()
	if l.TraceId == "" {
		delete(l.Caller, "trace_id")
	}
	l.Logger.WithFields(l.Caller).Debug(format)
}

func (l *Logger) Debugf(format string,args ...interface{}){
	l.caller()
	if l.TraceId == "" {
		delete(l.Caller, "trace_id")
	}
	l.Logger.WithFields(l.Caller).Debugf(format,args)

}

func (l *Logger) Info(format string){
	l.caller()
	if l.TraceId == "" {
		delete(l.Caller, "trace_id")
	}
	l.Logger.WithFields(l.Caller).Info(format)
}

func (l *Logger) Infof(format string,args ...interface{}){
	l.caller()
	if l.TraceId == "" {
		delete(l.Caller, "trace_id")
	}
	l.Logger.WithFields(l.Caller).Infof(format,args)
}

func (l *Logger) Warn(format string){
	l.caller()
	if l.TraceId == "" {
		delete(l.Caller, "trace_id")
	}
	l.Logger.WithFields(l.Caller).Warn(format)
}

func (l *Logger) Warnf(format string,args ...interface{}){
	l.caller()
	if l.TraceId == "" {
		delete(l.Caller, "trace_id")
	}
	l.Logger.WithFields(l.Caller).Warnf(format,args)
}

func (l *Logger) Error(format string){
	l.caller()
	if l.TraceId == "" {
		delete(l.Caller, "trace_id")
	}
	l.Logger.WithFields(l.Caller).Error(format)
}

func (l *Logger) Errorf(format string,args ...interface{}){
	l.caller()
	if l.TraceId == "" {
		delete(l.Caller, "trace_id")
	}
	l.Logger.WithFields(l.Caller).Errorf(format,args)
}

func (l *Logger) Fatal(format string){
	l.caller()
	if l.TraceId == "" {
		delete(l.Caller, "trace_id")
	}
	l.Logger.WithFields(l.Caller).Fatal(format)
}

func (l *Logger) Fatalf(format string,args ...interface{}){
	l.caller()
	if l.TraceId == "" {
		delete(l.Caller, "trace_id")
	}
	l.Logger.WithFields(l.Caller).Fatalf(format,args)
}

func (l *Logger) Panic(format string){
	l.caller()
	if l.TraceId == "" {
		delete(l.Caller, "trace_id")
	}
	l.Logger.WithFields(l.Caller).Panic(format)
}

func (l *Logger) Panicf(format string,args ...interface{}){
	l.caller()
	if l.TraceId == "" {
		delete(l.Caller, "trace_id")
	}
	l.Logger.WithFields(l.Caller).Panicf(format,args)
}

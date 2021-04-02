package logger

import "github.com/sirupsen/logrus"

func (l *Logger) Debug(format string){
	if l.TraceId != "" {
		l.Logger.WithFields(logrus.Fields{
			"trace_id":l.TraceId,
		}).Debug(format)
	} else {
		l.Logger.Debug(format)
	}
}

func (l *Logger) Debugf(format string,args ...interface{}){
	if l.TraceId != "" {
		l.Logger.WithFields(logrus.Fields{
			"trace_id":l.TraceId,
		}).Debugf(format,args...)
	} else {
		l.Logger.Debugf(format,args...)
	}
}

func (l *Logger) Info(format string){
	if l.TraceId != "" {
		l.Logger.WithFields(logrus.Fields{
			"trace_id":l.TraceId,
		}).Info(format)
	} else {
		l.Logger.Info(format)
	}
}

func (l *Logger) Infof(format string,args ...interface{}){
	if l.TraceId != "" {
		l.Logger.WithFields(logrus.Fields{
			"trace_id":l.TraceId,
		}).Infof(format,args...)
	} else {
		l.Logger.Infof(format,args...)
	}
}

func (l *Logger) Warn(format string){
	if l.TraceId != "" {
		l.Logger.WithFields(logrus.Fields{
			"trace_id":l.TraceId,
		}).Warn(format)
	} else {
		l.Logger.Warn(format)
	}
}

func (l *Logger) Warnf(format string,args ...interface{}){
	if l.TraceId != "" {
		l.Logger.WithFields(logrus.Fields{
			"trace_id":l.TraceId,
		}).Warnf(format,args...)
	} else {
		l.Logger.Warnf(format,args...)
	}
}

func (l *Logger) Error(format string){
	if l.TraceId != "" {
		l.Logger.WithFields(logrus.Fields{
			"trace_id":l.TraceId,
		}).Error(format)
	} else {
		l.Logger.Error(format)
	}
}

func (l *Logger) Errorf(format string,args ...interface{}){
	if l.TraceId != "" {
		l.Logger.WithFields(logrus.Fields{
			"trace_id":l.TraceId,
		}).Errorf(format,args...)
	} else {
		l.Logger.Errorf(format,args...)
	}
}

func (l *Logger) Fatal(format string){
	if l.TraceId != "" {
		l.Logger.WithFields(logrus.Fields{
			"trace_id":l.TraceId,
		}).Fatal(format)
	} else {
		l.Logger.Fatal(format)
	}
}

func (l *Logger) Fatalf(format string,args ...interface{}){
	if l.TraceId != "" {
		l.Logger.WithFields(logrus.Fields{
			"trace_id":l.TraceId,
		}).Fatalf(format,args...)
	} else {
		l.Logger.Fatalf(format,args...)
	}
}

func (l *Logger) Panic(format string){
	if l.TraceId != "" {
		l.Logger.WithFields(logrus.Fields{
			"trace_id":l.TraceId,
		}).Panic(format)
	} else {
		l.Logger.Panic(format)
	}
}

func (l *Logger) Panicf(format string,args ...interface{}){
	if l.TraceId != "" {
		l.Logger.WithFields(logrus.Fields{
			"trace_id":l.TraceId,
		}).Panicf(format,args...)
	} else {
		l.Logger.Panicf(format,args...)
	}
}

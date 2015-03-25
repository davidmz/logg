package logg

// TRACE fires log event with TRACE level
func (l *Logger) TRACE(format string, vs ...interface{}) { l.Log(TRACE, format, vs...) }

// DEBUG fires log event with DEBUG level
func (l *Logger) DEBUG(format string, vs ...interface{}) { l.Log(DEBUG, format, vs...) }

// INFO fires log event with INFO level
func (l *Logger) INFO(format string, vs ...interface{}) { l.Log(INFO, format, vs...) }

// WARN fires log event with WARN level
func (l *Logger) WARN(format string, vs ...interface{}) { l.Log(WARN, format, vs...) }

// ERROR fires log event with ERROR level
func (l *Logger) ERROR(format string, vs ...interface{}) { l.Log(ERROR, format, vs...) }

// FATAL fires log event with FATAL level
func (l *Logger) FATAL(format string, vs ...interface{}) { l.Log(FATAL, format, vs...) }

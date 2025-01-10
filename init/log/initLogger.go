package log

// TODO: Check proper error handling in both the loggers

// Uncomment to check if the logger and logger2 deeply equal and have the same address
// var logger = GetCustomZeroLogger()
// var logger2 = GetCustomZapLogger()
// fmt.Println(reflect.DeepEqual(logger, logger2))
// fmt.Println(&logger, &logger2)

// GlobalLogger initialised with custom logger, every import will return the same logger instance
var GlobalLogger = GetCustomZapLogger()

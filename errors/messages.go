package errors

const (
	InputError    = "ERROR - invalid command\nusage:\ngo run . <filename>"
	FileOpenError = "ERROR - can not open file\nis given filename correct?\nusage:\ngo run . <filename>"
	ScannerError  = "ERROR - can not scan file\n"
	InvalidData   = "ERROR: invalid data format\n"
	SameNameError = "ERROR - invalid data format\ntwo rooms have the same names"
)

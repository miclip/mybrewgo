package ui

// UI ...
type UI interface {
	AskForText(label string) (string, error)
	AskForFloat(label string) (float64, error)
	AskForInt(label string) (int, error)

	PrintLinef(pattern string, args ...interface{})
	SystemLinef(pattern string, args ...interface{})
	ErrorLinef(pattern string, args ...interface{})

	Printf(pattern string, args ...interface{})
	Systemf(pattern string, args ...interface{})
	Errorf(pattern string, args ...interface{})

	DisplayColumns(items []string, columns int)

	SetMaxInvalidInput(value int)
}

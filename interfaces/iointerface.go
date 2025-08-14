package interfaces

type IOManagerInt interface {
	ReadLines() ([]string, error)
	WriteResult(data any) error
}

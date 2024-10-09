package structExample

import "fmt"

type DivideError struct{
	Divided float64
	Dividor float64
}

func(e *DivideError) Error() string{
	return fmt.Sprintf("cannnot divide %f by %f", e.Divided, e.Dividor)
}
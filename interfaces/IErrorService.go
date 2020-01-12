package interfaces

type IErrorService interface {
	NewError(err error, status int, detail string) error
}

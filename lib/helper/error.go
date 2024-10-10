package helper

type ErrBadRequest struct {
  Message string
}

func NewErrBadRequest(message string) *ErrBadRequest {
  return &ErrBadRequest{
    Message: message,
  }
}

func (e *ErrBadRequest) Error() string {
  return e.Message
}

type ErrNotFound struct {
  Message string
}

func NewErrNotFound(message string) *ErrNotFound {
  return &ErrNotFound{
    Message: message,
  }
}

func (e *ErrNotFound) Error() string {
  return e.Message
}

type ErrInternalServerError struct {
  Message string
}

func NewErrInternalServerError(message string) *ErrInternalServerError {
  return &ErrInternalServerError{
    Message: message,
  }
}

func (e *ErrInternalServerError) Error() string {
  return e.Message
}

type ErrUnauthorized struct {
  Message string
}

func NewErrUnauthorized(message string) *ErrUnauthorized {
  return &ErrUnauthorized{
    Message: message,
  }
}

func (e *ErrUnauthorized) Error() string {
  return e.Message
}

type ErrForbidden struct {
  Message string
}

func NewErrForbidden(message string) *ErrForbidden {
  return &ErrForbidden{
    Message: message,
  }
}

func (e *ErrForbidden) Error() string {
  return e.Message
}


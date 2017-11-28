package helpers

type UnexpectedStatus struct{
	Msg string
}

func (e *UnexpectedStatus) Error() string {
	return e.Msg
}

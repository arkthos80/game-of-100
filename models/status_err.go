package models

type StatusErr struct {
	GameStatus Status
	Message    string
}

func (se StatusErr) Error() string {
	return se.Message
}
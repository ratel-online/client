package consts

type StateID int

const (
	_ StateID = iota
	StateWelcome
	StateHome
	StateJoin
	StateNew
	StateSetting
	StateWaiting
	StateClassics
	StateLaiZi
)

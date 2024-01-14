package pm

type Typist interface {
	Type(text string, delayMs int) (string, error)
}

type Menu interface {
	Dmenu(menu string) (string, error)
	Message(msg string) error
}

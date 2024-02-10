package pm

type Typist interface {
	Type(text string, delayMs int) (string, error)
}

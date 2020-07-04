package config

const (
	BotToken = "1285255270:AAEjv4wFeOKp08oX5cUYYjoekExBAU6JNfo"
	DevID    = 662834330
)

const (
	DefaultFieldDimension = 9
	PlayersCount          = 3
	Horizon               = 1
)

// Separated in order to make iotas work properly
const (
	//User state constants
	StateNone         = iota // no interaction possible at all
	StateAFK          = iota // passive state for user
	StateChangingName = iota // user is changing nickname
	StateWriting      = iota // user answers the question
	StateWaiting      = iota // user is waiting for his turn
	StateThinking     = iota // there is user's move now
)

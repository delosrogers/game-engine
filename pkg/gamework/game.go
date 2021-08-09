package gamework

type Game struct {
	// default nil
	GameStateReducer GameStateReducerFunc
	// default nil
	InitialState GameState
	// default 2
	MaxPlayers int
	// default ""
	JsLocation string
}

// a string enum that allows the framework to know how the move went
type MoveResult string

const (
	Valid MoveResult = "VALID"
	Invalid MoveResult = "INVALID"
	Error MoveResult = "ERROR"
	Tie MoveResult = "TIE"
	Win MoveResult = "Win"
)

type GameMoveResult struct {
	Type MoveResult
	NewState GameState
}

type GameStateReducerFunc func(string) (GameMoveResult, error)

// your state struct must by JSON serializable
type GameState interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(b []byte) error
}

func NewGame() *Game {
	return &Game{
		GameStateReducer: nil,
		InitialState: nil,
		MaxPlayers: 2,
		JsLocation: "",

	}
}
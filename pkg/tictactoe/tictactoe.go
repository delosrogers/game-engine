package tictactoe

import "encoding/json"

type state [][]string
type move struct {
	x int
	y int
	newVal string
}


func GameStateReducer(stateString string, moveString string)  {
	var thisMove move
	err := json.Unmarshal([]byte(moveString), &thisMove)
	if err != nil {
		return "ERROR"
	}
	var s state
	err = json.Unmarshal([]byte(stateString), &s)
	if err != nil {
		return "ERROR"
	}
	s[thisMove.y][thisMove.x] = thisMove.newVal

	newState, err := json.Marshal(s)
	var statusString string
	if s.checkWin() {
		statusString = "WIN"
	} else {
		statusString = "VALID"
	}
	return statusString + string(newState)
}

func (s state) checkWin() bool {
	win := s[0][0] == s[0][1] && s[0][1] == s[0][2] ||
		 s[1][0] == s[1][1] && s[1][1] == s[1][2] ||
		 s[2][0] == s[2][1] && s[2][1] == s[2][2] ||
		 s[0][0] == s[1][1] && s[1][1] == s[2][2] ||
		 s[0][2] == s[1][1] && s[1][1] == s[2][0]
	return win
}

func GetInitialState(state) ([]byte, error) {
	return json.Marshal(state{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	})
}

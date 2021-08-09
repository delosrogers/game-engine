# Basic Interface:
All game backends must provide the following functions:
- `GameStateReducer(state string, move string) string` State is a text encoded state
that the frameworks maintains for you and is probably a JSON string. Move is a text
  encoded object that represents a move made by a player. This function must return the new state
  prepended by one of the following strings to let the framework know the status of the move
    - `ERROR` used for actual errors in your game logic
    - `INVALID` when the client tries to make an illegal move
    - `VALID`
    - `WIN`
    - `TIE`
    
- `GetInitialState() string` which returns the initial state of your game encoded in which
ever way you decide to encode state
package messaging

type GuessingGameData struct {
	GameID        string `json:"gameid"`
	NbrOfAttempts int    `json:"nbrofattempts"`
	GeneratedNbr  int    `json:"generatednbr"`
	GuessNbr      int    `json:"guessnbr"`
	GameOver      bool   `json:"gameover"`
	Timestamp     string `json:"timestamp"`
	Message       string `json:"message,omitempty"`
}

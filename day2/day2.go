package day2

import (
	"bufio"
	"embed"
	"fmt"
	"strings"
)

const (
	opponentRock     = "A"
	opponentPaper    = "B"
	opponentScissors = "C"
	meRock           = "X"
	mePaper          = "Y"
	meScissors       = "Z"
	needWin          = "Z"
	needLose         = "X"
	needDraw         = "Y"
)

type round struct {
	opponentMove string
	myMove       string
}

func (r round) won() bool {
	rockWin := r.myMove == meRock && r.opponentMove == opponentScissors
	paperWin := r.myMove == mePaper && r.opponentMove == opponentRock
	scissorsWin := r.myMove == meScissors && r.opponentMove == opponentPaper
	return rockWin || paperWin || scissorsWin
}

func (r round) lost() bool {
	rockLost := r.myMove == meRock && r.opponentMove == opponentPaper
	paperLost := r.myMove == mePaper && r.opponentMove == opponentScissors
	scissorsLost := r.myMove == meScissors && r.opponentMove == opponentRock
	return rockLost || paperLost || scissorsLost
}

func (r round) shapeScore() int {
	switch r.myMove {
	case meRock:
		return 1
	case mePaper:
		return 2
	default:
		return 3
	}
}

func (r round) outcomeScore() int {
	if r.lost() {
		return 0
	}
	if r.won() {
		return 6
	}
	return 3
}

func (r round) score() int {
	return r.shapeScore() + r.outcomeScore()
}

func (r round) correctScore() int {
	correctRound := round{
		opponentMove: r.opponentMove,
		myMove:       r.correctMove(),
	}
	return correctRound.score()
}

func (r round) correctMove() string {
	switch r.myMove {
	case needDraw:
		switch r.opponentMove {
		case opponentRock:
			return meRock
		case opponentPaper:
			return mePaper
		default:
			return meScissors
		}
	case needLose:
		switch r.opponentMove {
		case opponentRock:
			return meScissors
		case opponentPaper:
			return meRock
		default:
			return mePaper
		}
	default:
		switch r.opponentMove {
		case opponentRock:
			return mePaper
		case opponentPaper:
			return meScissors
		default:
			return meRock
		}
	}
}

//go:embed input.txt
//go:embed testInput.txt

var f embed.FS

func ReadInput(test bool) ([]round, error) {
	path := "input.txt"
	if test {
		path = "testInput.txt"
	}
	file, err := f.Open(path)
	if err != nil {
		return nil, fmt.Errorf("reading input file %q: %w", path, err)
	}
	defer file.Close()

	fs := bufio.NewScanner(file)
	fs.Split(bufio.ScanLines)
	var rounds []round
	for fs.Scan() {
		l := fs.Text()
		ls := strings.Split(l, " ")
		rounds = append(rounds, round{
			opponentMove: ls[0],
			myMove:       ls[1],
		})
	}
	return rounds, nil
}

func TotalScore(rs []round) int {
	var c int
	for _, r := range rs {
		c += r.score()
	}
	return c
}

func TotalCorrectScore(rs []round) int {
	var c int
	for _, r := range rs {
		c += r.correctScore()
	}
	return c
}

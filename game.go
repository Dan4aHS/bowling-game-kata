package main

type Game struct {
	roll  int
	rolls []int
}

func NewGame() *Game {
	return &Game{
		roll:  0,
		rolls: make([]int, 21),
	}
}

func (g *Game) Roll(pins int) {
	g.rolls[g.roll] = pins
	g.roll++
}

func (g *Game) Rolls(rolls ...int) {
	for _, roll := range rolls {
		g.Roll(roll)
	}
}

func (g *Game) Score() int {
	cursor, score := 0, 0
	for _ = range 10 {
		switch {
		case g.isStrike(cursor):
			score += 10 + g.rolls[cursor+1] + g.rolls[cursor+2]
			cursor++
		case g.isSpare(cursor):
			score += 10 + g.rolls[cursor+2]
			cursor += 2
		default:
			score += g.rolls[cursor] + g.rolls[cursor+1]
			cursor += 2
		}

	}
	return score
}

func (g *Game) isSpare(cursor int) bool {
	return g.rolls[cursor]+g.rolls[cursor+1] == 10
}

func (g *Game) isStrike(cursor int) bool {
	return g.rolls[cursor] == 10
}

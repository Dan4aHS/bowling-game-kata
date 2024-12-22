package main

import "testing"

// TestGame - обертка для тестирования игры
type TestGame struct {
	game *Game
}

// setupTestGame - инициализация тестовой игры
func setupTestGame() *TestGame {
	return &TestGame{
		game: NewGame(),
	}
}

// TestFullGame - запуск всех тестов для игры
func TestFullGame(t *testing.T) {
	t.Run("CanScoreGutterGame", setupTestGame().testGutterGame)
	t.Run("CanScoreGameOfOnes", setupTestGame().testGameOfOnes)
	t.Run("CanScoreSpareFollowedByThree", setupTestGame().testSpareFollowedByThree)
	t.Run("CanStrikeFollowedByFiveAndFive", setupTestGame().testStrikeFollowedByFiveAndFive)
	t.Run("CanPerfectGame", setupTestGame().testPerfectGame)
	t.Run("CanSomeGame1", setupTestGame().testSomeGame1)
}

func (tg *TestGame) testGutterGame(t *testing.T) {
	tg.game.Rolls(
		0, 0,
		0, 0,
		0, 0,
		0, 0,
		0, 0,
		0, 0,
		0, 0,
		0, 0,
		0, 0,
		0, 0,
	)

	assertScore(t, tg.game.Score(), 0)
}

func (tg *TestGame) testGameOfOnes(t *testing.T) {
	tg.game.Rolls(
		1, 1,
		1, 1,
		1, 1,
		1, 1,
		1, 1,
		1, 1,
		1, 1,
		1, 1,
		1, 1,
		1, 1)

	assertScore(t, tg.game.Score(), 20)
}

func (tg *TestGame) testSpareFollowedByThree(t *testing.T) {
	tg.game.Rolls(
		5, 5,
		3, 0,
		0, 0,
		0, 0,
		0, 0,
		0, 0,
		0, 0,
		0, 0,
		0, 0,
		0, 0,
	)
	assertScore(t, tg.game.Score(), 16)
}

func (tg *TestGame) testStrikeFollowedByFiveAndFive(t *testing.T) {
	tg.game.Rolls(
		10,
		5, 5,
		0, 0,
		0, 0,
		0, 0,
		0, 0,
		0, 0,
		0, 0,
		0, 0,
		0, 0,
	)

	assertScore(t, tg.game.Score(), 30)
}

func (tg *TestGame) testPerfectGame(t *testing.T) {
	tg.game.Rolls(
		10,
		10,
		10,
		10,
		10,
		10,
		10,
		10,
		10,
		10,
		10,
		10,
	)

	assertScore(t, tg.game.Score(), 300)
}

func (tg *TestGame) testSomeGame1(t *testing.T) {
	tg.game.Rolls(
		10,
		5, 5,
		3, 0,
		10,
		5,
		0, 0,
		0, 0,
		0, 0,
		0, 0,
		0, 0,
	)
	assertScore(t, tg.game.Score(), 56)
}

func assertScore(t *testing.T, actual, expected int) {
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

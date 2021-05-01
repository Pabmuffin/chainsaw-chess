package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/notnil/chess"
)

func main() {
	game := chess.NewGame()
	rand.Seed(time.Now().UnixNano())
	// generate moves until game is over
	for game.Outcome() == chess.NoOutcome {
		// select a random move
		moves := game.ValidMoves()
		move := moves[rand.Intn(len(moves))]
		game.Move(move)
	}
	// print outcome and game PGN
	fmt.Println(game.Position().Board().Draw())
	fmt.Printf("Game completed. %s by %s.\n", game.Outcome(), game.Method())
	fmt.Println(game.String())
}

func decideBestMove(game chess.Game) *chess.Move{
	moves := game.ValidMoves()
	bestMoveIdx := 0 // initialize to first legal move
	bestMoveScore := 0
	for moveIdx := 0; moveIdx < len(moves); moveIdx++ {
		newMoveScore := scoreMove(moves[moveIdx], game)
		if newMoveScore > bestMoveScore {
			bestMoveScore = newMoveScore
			bestMoveIdx = moveIdx
		}
	}
	return moves[bestMoveIdx]
}

func scoreMove(move *chess.Move, game chess.Game) int {
	// Start with material advantage only, one layer deep



	return 0;
}

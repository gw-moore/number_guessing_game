package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Player struct {
	name  string
	guess int
}

type Game struct {
	player1 Player
	player2 Player
	random  int
}

// create a player
func (p *Player) createPlayer(name string) {
	p.name = name
}

// get random number
func (g *Game) getRandNum() {
	// set a seed for this game so rand is rand
	rand.Seed(time.Now().UnixNano())
	g.random = rand.Intn(100) + 1
}

// create a game
func (g *Game) createGame() {
	g.player1.createPlayer("Player 1")
	g.player2.createPlayer("Player 2")
	g.getRandNum()
}

// read input from the terminal
func (p *Player) readInput() {
	var input int
	fmt.Printf("%s, enter your guess: ", p.name)
	fmt.Scanf("%d", &input)
	p.guess = input
}

// create game loop
func (g *Game) gameLoop() {
	// explain the rules of the game
	fmt.Println("Each player will guess a number between 1 and 100. The player closest to the random number wins.")

	// loop forever
	for {
		// read player input
		g.player1.readInput()
		g.player2.readInput()

		// check which player guessed closer to the random number
		p1Diff := math.Abs(float64(g.player1.guess) - float64(g.random))
		p2Diff := math.Abs(float64(g.player2.guess) - float64(g.random))

		if p1Diff < p2Diff {
			fmt.Printf("%s wins!\n", g.player1.name)
		} else if p1Diff > p2Diff {
			fmt.Printf("%s wins!\n", g.player2.name)
		} else {
			fmt.Println("Its a tie!")
		}
		fmt.Printf("The number was: %d\n", g.random)

		// check if the players want to play again
		var playAgain string
		fmt.Print("Would you play to play again? (y/n): ")
		fmt.Scanf("%s", &playAgain)

		if playAgain == "n" {
			fmt.Println("Thanks for playing!")
			break
		}

		// refresh the random number
		g.getRandNum()
	}
}

func main() {
	game := Game{}
	game.createGame()
	game.gameLoop()
}

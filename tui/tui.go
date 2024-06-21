package tui

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/AtleSkaanes/gosnake/game"
	"github.com/AtleSkaanes/gosnake/game/types"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	width  uint8
	height uint8
}

var newDirection types.Direction = types.NoDirection
var lastGameState types.GameEvent = types.GameContinue
var isPaused bool = false
var timeSinceUpdate time.Time

func Init() {
	timeSinceUpdate = time.Now()

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

}

func initialModel() model {
	return model{
		width:  16,
		height: 16,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// Quit on ctrl+c or q
		case "ctrl+c", "q":
			return m, tea.Quit

			// Pause on p or esc
		case "p", "esc":
			isPaused = !isPaused

		// Move up
		case "up", "k", "w":
			newDirection = types.Up

			// Move down
		case "down", "j", "s":
			newDirection = types.Down

			// Mode left
		case "left", "h", "a":
			newDirection = types.Left

			// Move right
		case "right", "l", "d":
			newDirection = types.Right
		}
	}

	if !isPaused && time.Since(timeSinceUpdate) >= 100*time.Millisecond {
		timeSinceUpdate = time.Now()
		lastGameState = game.Update(newDirection)
		newDirection = types.NoDirection
	}
	if lastGameState == types.GameLost {
		return m, tea.Quit
	}

	return m, func() tea.Msg { return 0 }
}

func (m model) View() string {
	if !game.IsInitialized {
		return "╔═════════╗\n║ LOADING ║\n╚═════════╝"
	}

	pausedSign := "󰐊"
	if isPaused {
		pausedSign = "󰏤"
	}

	score := game.GetScore()
	dimensions := game.GetConf().Dimensions

	s := "╔" + strings.Repeat("═", dimensions.X*2) + "╗\n"
	s += fmt.Sprintf("║ : %d%s %s ║\n", score, strings.Repeat(" ", dimensions.X*2-7-NumLen(score)), pausedSign)
	s += "╠" + strings.Repeat("═", dimensions.X*2) + "╣\n"

	for y := 0; y < dimensions.Y; y++ {
		s += "║"
		for x := 0; x < dimensions.X; x++ {
			pos := types.NewVec2(x, y)
			if game.GetSnake().GetHead().IsEqual(pos) {
				s += "##"
			} else if game.GetSnake().IsOnBody(pos) {
				s += "██"
			} else if game.GetApple().IsEqual(pos) {
				s += ""
			} else {
				s += "  "
			}
		}
		s += "║\n"
	}

	s += "╚" + strings.Repeat("═", dimensions.X*2) + "╝\n"
	return s
}

func NumLen(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

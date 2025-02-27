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
	width     int
	height    int
	isPaused  bool
	speed     int
	isColored bool
}

var newDirection types.Direction = types.NoDirection
var lastGameState types.GameEvent = types.GameContinue
var timeSinceUpdate time.Time

func Init(speed int, withColor bool) {
	timeSinceUpdate = time.Now()

	dimensions := game.GetConf().Dimensions
	m := model{
		width:     dimensions.X,
		height:    dimensions.Y,
		isPaused:  false,
		speed:     speed,
		isColored: withColor,
	}

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
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
			m.isPaused = !m.isPaused

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

	if !m.isPaused && time.Since(timeSinceUpdate) >= time.Duration(m.speed)*time.Millisecond {
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
	if m.isPaused {
		pausedSign = "󰏤"
	}

	score := game.GetScore()

	s := "╔" + strings.Repeat("═", m.width*2) + "╗\n"
	s += fmt.Sprintf("║ : %d%s %s ║\n", score, strings.Repeat(" ", m.width*2-7-NumLen(score)), pausedSign)
	s += "╠" + strings.Repeat("═", m.width*2) + "╣\n"

	for y := 0; y < m.height; y++ {
		s += "║"
		for x := 0; x < m.width; x++ {
			pos := types.NewVec2(x, y)
			if game.GetSnake().GetHead().IsEqual(pos) {
				s += green("##", m.isColored)
			} else if game.GetSnake().IsOnBody(pos) {
				s += green("██", m.isColored)
			} else if game.GetApple().IsEqual(pos) {
				s += red("", m.isColored)
			} else {
				s += "  "
			}
		}
		s += "║\n"
	}

	s += "╚" + strings.Repeat("═", m.width*2) + "╝\n"
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

func red(str string, color bool) string {
	if !color {
		return str
	}
	return "\x1B[31m" + str + "\x1B[0m"
}

func green(str string, color bool) string {
	if !color {
		return str
	}
	return "\x1B[32m" + str + "\x1B[0m"
}

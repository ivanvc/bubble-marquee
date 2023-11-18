package main

// A simple program demonstrating the spinner component from the Bubbles
// component library.

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	marquee "github.com/ivanvc/bubble-marquee"
)

type errMsg error

type model struct {
	marquee  marquee.Model
	quitting bool
	err      error
}

func initialModel() model {
	m := marquee.New()
	m.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	m.Text = "testing 1 2 3"
	return model{marquee: m}
}

func (m model) Init() tea.Cmd {
	return nil //m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}

	case errMsg:
		m.err = msg
		return m, nil

	default:
		var cmd tea.Cmd
		m.marquee, cmd = m.marquee.Update(msg)
		return m, cmd
	}
}

func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	str := fmt.Sprintf("\n\n Marquee:\n%s\n\n", m.marquee.View())
	if m.quitting {
		return str + "\n"
	}
	return str
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package main

// A simple program demonstrating the spinner component from the Bubbles
// component library.

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	marquee "github.com/ivanvc/bubble-marquee"
)

type errMsg error

type model struct {
	m1       marquee.Model
	m2       marquee.Model
	m3       marquee.Model
	m4       marquee.Model
	quitting bool
	err      error
}

func initialModel() model {
	m1 := marquee.New()
	m1.Style = lipgloss.NewStyle().
		Foreground(lipgloss.Color("205")).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("63"))
	m1.SetText("Hello World")
	m1.SetWidth(25)
	m1.ScrollDirection = marquee.Right

	m2 := marquee.New()
	m2.Style = lipgloss.NewStyle().
		Foreground(lipgloss.Color("63")).
		BorderStyle(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("205"))
	m2.SetText("Hello World")
	m2.SetWidth(25)
	m2.ScrollSpeed = 500 * time.Millisecond

	m3 := marquee.New()
	m3.Style = lipgloss.NewStyle().
		Foreground(lipgloss.Color("207"))
	m3.SetText(fmt.Sprintf("The time is: %s", time.Now().Format("15:04:05")))
	m3.SetWidth(54)
	m3.ScrollDirection = marquee.Right
	m3.ScrollSpeed = 100 * time.Millisecond

	m4 := marquee.New()
	m4.Style = lipgloss.NewStyle().
		Foreground(lipgloss.Color("65")).
		BorderStyle(lipgloss.DoubleBorder()).
		BorderForeground(lipgloss.Color("205"))
	m4.SetText("Auto width")

	return model{m1: m1, m2: m2, m3: m3, m4: m4}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.m1.Scroll, m.m2.Scroll, m.m3.Scroll, m.m4.Scroll)
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

	case marquee.ScrollMsg:
		var cmds []tea.Cmd
		var cmd tea.Cmd

		m.m1, cmd = m.m1.Update(msg)
		cmds = append(cmds, cmd)

		m.m2, cmd = m.m2.Update(msg)
		cmds = append(cmds, cmd)

		m.m3.SetText(fmt.Sprintf("The time is: %s", time.Now().Format("15:04:05")))
		m.m3, cmd = m.m3.Update(msg)
		cmds = append(cmds, cmd)

		m.m4, cmd = m.m4.Update(msg)
		cmds = append(cmds, cmd)
		return m, tea.Batch(cmds...)

	case errMsg:
		m.err = msg
		return m, nil

	default:
		return m, nil
	}
}

func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	str := fmt.Sprintf("%s\n%s\n%s\n", lipgloss.JoinHorizontal(lipgloss.Center, m.m1.View(), m.m2.View()), m.m3.View(), m.m4.View())
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

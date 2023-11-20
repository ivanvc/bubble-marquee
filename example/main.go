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
	m5       marquee.Model
	m6       marquee.Model
	m7       marquee.Model
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

	m5 := marquee.New()
	m5.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("205"))
	m5.SetText(
		lipgloss.NewStyle().Background(lipgloss.Color("4")).Bold(true).
			SetString(" Hello ").Render() + " " +
			lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Underline(true).
				SetString("World").Render(),
	)
	m5.SetWidth(50)
	m5.ScrollSpeed = 75 * time.Millisecond

	m6 := marquee.New()
	m6.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.ThickBorder()).
		BorderTop(true).BorderLeft(true).BorderRight(true).
		BorderForeground(lipgloss.Color("254"))
	m6.SetText(
		lipgloss.NewStyle().Background(lipgloss.Color("1")).Bold(true).
			SetString(" BREAKING NEWS ").Render() +
			lipgloss.NewStyle().Foreground(lipgloss.Color("0")).
				Background(lipgloss.Color("15")).
				SetString(" Marquees are cool ").Render())
	m6.SetWidth(50)
	m6.ScrollSpeed = 150 * time.Millisecond
	m6.SetContinuous(true)

	m7 := marquee.New()
	m7.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.ThickBorder()).
		BorderBottom(true).BorderLeft(true).BorderRight(true).
		BorderForeground(lipgloss.Color("248"))
	m7.SetText(
		lipgloss.NewStyle().Background(lipgloss.Color("8")).Bold(true).
			Foreground(lipgloss.Color("15")).
			SetString(" SPX ").Render() +
			lipgloss.NewStyle().Background(lipgloss.Color("10")).Bold(true).
				Foreground(lipgloss.Color("0")).
				SetString(" +0.67% ").Render() +
			lipgloss.NewStyle().Background(lipgloss.Color("8")).Bold(true).
				Foreground(lipgloss.Color("15")).
				SetString(" CCMP ").Render() +
			lipgloss.NewStyle().Background(lipgloss.Color("10")).Bold(true).
				Foreground(lipgloss.Color("0")).
				SetString(" +1.05% ").Render() +
			lipgloss.NewStyle().Background(lipgloss.Color("8")).Bold(true).
				Foreground(lipgloss.Color("15")).
				SetString(" DAX ").Render() +
			lipgloss.NewStyle().Background(lipgloss.Color("9")).Bold(true).
				Foreground(lipgloss.Color("0")).
				SetString(" -0.11% ").Render() +
			lipgloss.NewStyle().Background(lipgloss.Color("8")).Bold(true).
				Foreground(lipgloss.Color("15")).
				SetString(" BTC ").Render() +
			lipgloss.NewStyle().Background(lipgloss.Color("10")).Bold(true).
				Foreground(lipgloss.Color("0")).
				SetString(" +1.73% ").Render() +
			lipgloss.NewStyle().Background(lipgloss.Color("8")).Bold(true).
				Foreground(lipgloss.Color("15")).
				SetString(" ETH ").Render() +
			lipgloss.NewStyle().Background(lipgloss.Color("10")).Bold(true).
				Foreground(lipgloss.Color("0")).
				SetString(" +3.10% ").Render() +
			lipgloss.NewStyle().Background(lipgloss.Color("8")).Bold(true).
				Foreground(lipgloss.Color("15")).
				SetString(" USD/JPY ").Render() +
			lipgloss.NewStyle().Background(lipgloss.Color("9")).Bold(true).
				Foreground(lipgloss.Color("0")).
				SetString(" +0.67% ").Render() +
			lipgloss.NewStyle().Background(lipgloss.Color("8")).Bold(true).
				Foreground(lipgloss.Color("15")).
				SetString(" GC ").Render() +
			lipgloss.NewStyle().Background(lipgloss.Color("9")).Bold(true).
				Foreground(lipgloss.Color("0")).
				SetString(" -0.29% ").Render() +
			lipgloss.NewStyle().Background(lipgloss.Color("8")).Bold(true).
				Foreground(lipgloss.Color("15")).
				SetString(" XAG ").Render() +
			lipgloss.NewStyle().Background(lipgloss.Color("9")).Bold(true).
				Foreground(lipgloss.Color("0")).
				SetString(" -0.94% ").Render(),
	)
	m7.SetWidth(50)
	m7.ScrollDirection = marquee.Right
	m7.ScrollSpeed = 150 * time.Millisecond
	m7.SetContinuous(true)

	return model{m1: m1, m2: m2, m3: m3, m4: m4, m5: m5, m6: m6, m7: m7}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		m.m1.Scroll,
		m.m2.Scroll,
		m.m3.Scroll,
		m.m4.Scroll,
		m.m5.Scroll,
		m.m6.Scroll,
		m.m7.Scroll,
	)
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

		m.m5, cmd = m.m5.Update(msg)
		cmds = append(cmds, cmd)

		m.m6, cmd = m.m6.Update(msg)
		cmds = append(cmds, cmd)

		m.m7, cmd = m.m7.Update(msg)
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
	str := "Two marquees next to each other:\n"
	str += lipgloss.JoinHorizontal(lipgloss.Center, m.m1.View(), m.m2.View())
	str += "\nA marquee with updating text:\n"
	str += m.m3.View()
	str += "\nA marquee without a width:\n"
	str += m.m4.View()
	str += "\nA marquee with lipgloss styled text:\n"
	str += m.m5.View()
	str += "\nMarquees with continuous scrolling:\n"
	str += m.m6.View()
	str += "\n"
	str += m.m7.View()

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

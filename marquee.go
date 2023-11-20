package marquee

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/ansi"
	"github.com/muesli/termenv"
)

const defaultScrollSpeed = time.Millisecond * 250

// Internal ID management. Used during animating to ensure that frame messages
// are received only by marquee components that sent them.
var (
	scrollID int
	idMtx    sync.Mutex
)

// Return the next ID we should use on the Model.
func nextID() int {
	idMtx.Lock()
	defer idMtx.Unlock()
	scrollID++
	return scrollID
}

// initialScrollMsg initializes the marquee scrolling.
type initialScrollMsg struct{}

// ScrollMsg signals that the marquee should scroll.
type ScrollMsg struct {
	id  int
	tag int
}

// scrollCanceled is sent when a scroll operation is canceled.
type scrollCanceled struct{}

// scrollCtx manages marquee scrolling.
type scrollCtx struct {
	ctx    context.Context
	cancel context.CancelFunc
}

// Direction describes the direction for the marquee.
type Direction int

// Available directions.
const (
	Left Direction = iota
	Right
)

// Model is the Bubble Tea model for the marquee.
type Model struct {
	ScrollSpeed time.Duration
	// Style for styling the marquee block.
	Style lipgloss.Style
	// The direction in which the text will scroll.
	ScrollDirection Direction

	// Text holds the text displayed by the marquee.
	text string

	// Width sets the view's width to a fixed value. If this is not specified,
	// the width is assumed to be the length of Text.
	width int

	// The ID of this Model as it relates to other marquees.
	id int
	// Used to manage marquee scrolling
	scrollCtx *scrollCtx
	// The ID of the blink message we're expecting to receive.
	tag int
	// Holds the current text view.
	textView string
	// The current text index being displayed
	textIndex int
	// The non-printable chars prefix.
	prefix string
	// If the text should be continuous.
	isTextContinuous bool
}

// New creates a new model with default settings.
func New() Model {
	return Model{
		ScrollSpeed: defaultScrollSpeed,

		scrollCtx: &scrollCtx{
			ctx: context.Background(),
		},
		id: nextID(),
	}
}

// Sets marquee text content.
func (m *Model) SetText(text string) {
	m.text = text
	if m.width == 0 {
		m.SetWidth(ansi.PrintableRuneWidth(m.text))
	}
	m.resetTextView()
}

func (m *Model) resetTextView() {
	if m.isTextContinuous {
		m.textView = m.text
	} else {
		space := strings.Repeat(" ", m.width)
		m.textView = fmt.Sprintf("%s%s%s", space, m.text, space)
	}
}

// Sets marquee width.
func (m *Model) SetWidth(width int) {
	if width <= 0 {
		return
	}
	m.width = width
	m.resetTextView()
}

// Sets if the text of the marquee should be continuous.
func (m *Model) SetContinuous(continuous bool) {
	m.isTextContinuous = continuous
	m.resetTextView()
}

// Update updates the marquee.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ScrollMsg:
		if msg.id != m.id || msg.tag != m.tag {
			return m, nil
		}

		if m.ScrollDirection == Left {
			m.textIndex++
			if m.textIndex > m.width+len(m.text) || m.isTextContinuous && m.textIndex > len(m.textView) {
				m.textIndex = 0
			}

			var index int
			m.prefix, _, index, _ = m.nonPrintableCharactersBeforeIndex(m.textIndex)
			if index > m.textIndex {
				m.textIndex = index
			}
		} else {
			m.textIndex--
			if m.textIndex < 0 {
				m.textIndex = m.width + len(m.text)
			}

			var index int
			_, m.prefix, _, index = m.nonPrintableCharactersBeforeIndex(m.textIndex)
			if index < m.textIndex {
				m.textIndex = index
			}
		}

		return m, m.tick(m.id, m.tag)
	case scrollCanceled:
		return m, nil
	}
	return m, nil
}

// ScrollCmd is the command to control the marquee scrolling.
func (m *Model) ScrollCmd() tea.Cmd {
	if m.scrollCtx != nil && m.scrollCtx.cancel != nil {
		m.scrollCtx.cancel()
	}

	ctx, cancel := context.WithTimeout(m.scrollCtx.ctx, m.ScrollSpeed)
	m.scrollCtx.cancel = cancel

	m.tag++

	return func() tea.Msg {
		defer cancel()
		<-ctx.Done()
		if ctx.Err() == context.DeadlineExceeded {
			return ScrollMsg{id: m.id, tag: m.tag}
		}
		return scrollCanceled{}
	}
}

// Scroll is a command used to initialize marquee scrolling.
func (m Model) Scroll() tea.Msg {
	return ScrollMsg{
		id:  m.id,
		tag: m.tag,
	}
}

func (m Model) tick(id, tag int) tea.Cmd {
	return tea.Tick(m.ScrollSpeed, func(t time.Time) tea.Msg {
		return ScrollMsg{
			id:  id,
			tag: tag,
		}
	})
}

// View displays the marquee.
func (m Model) View() string {
	var text string
	if m.isTextContinuous {
		var b strings.Builder
		b.WriteString(m.textView[m.textIndex:])
		for i := 0; ansi.PrintableRuneWidth(b.String()) < m.width; i = (i + 1) % len(m.textView) {
			b.WriteByte(m.textView[i])
		}
		text = b.String()
		for ansi.PrintableRuneWidth(text) > m.width {
			text = text[:len(text)-1]
		}
	} else {
		right := m.textIndex + m.width
		for ansi.PrintableRuneWidth(m.textView[m.textIndex:right]) < m.width {
			right++
		}
		text = m.textView[m.textIndex:right]
	}

	return m.Style.Render(fmt.Sprintf("%s%s%sm", m.prefix, text, termenv.CSI+termenv.ResetSeq))
}

func (m Model) nonPrintableCharactersBeforeIndex(index int) (string, string, int, int) {
	var b strings.Builder
	inModifier := false
	i := 0
	prev := len(m.textView)
	prevString := ""
	for i < len(m.textView) {
		c := rune(m.textView[i])
		if c == ansi.Marker {
			b.WriteRune(c)
			inModifier = true
		} else if inModifier {
			b.WriteRune(c)
			if ansi.IsTerminator(c) {
				inModifier = false
			}
		} else {
			if i <= index {
				prev = i
				prevString = b.String()
			}
			if i >= index {
				break
			}
		}
		i++
	}

	return b.String(), prevString, i, prev
}

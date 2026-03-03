package ui

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/fezcode/atlas.deck/internal/model"
)

type Model struct {
	Deck     *model.Deck
	Status   string
	Running  bool
	Width    int
	Height   int
	LastKey  string

	// Logs
	Viewport viewport.Model
	Logs     []string
	Spinner  spinner.Model
}

func NewModel(deck *model.Deck) Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#bd93f9"))

	status := "Ready"
	if deck == nil {
		status = "No deck.piml found."
	}

	vp := viewport.New(0, 0)
	vp.Style = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), true, false, false, false).
		BorderForeground(lipgloss.Color("#44475a")).
		Padding(0, 1)

	return Model{
		Deck:     deck,
		Status:   status,
		Spinner:  s,
		Viewport: vp,
		LastKey:  "None",
	}
}

func (m Model) Init() tea.Cmd {
	return m.Spinner.Tick
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "ctrl+l": // Clear logs
			m.Logs = []string{}
			m.Viewport.SetContent("")
			m.Status = "Logs cleared"
			return m, nil
		default:
			if !m.Running && m.Deck != nil {
				for _, pad := range m.Deck.Pads {
					if msg.String() == pad.Key {
						m.LastKey = pad.Key
						m.Running = true
						m.Status = fmt.Sprintf("Running: %s", pad.Label)
						
						// Add execution log and UPDATE VIEWPORT IMMEDIATELY
						logLine := lipgloss.NewStyle().Foreground(lipgloss.Color("#bd93f9")).Render(fmt.Sprintf(">>> [%s] Executing: %s", pad.Key, pad.Command))
						m.Logs = append(m.Logs, logLine)
						m.Viewport.SetContent(strings.Join(m.Logs, "\n"))
						m.Viewport.GotoBottom()
						
						return m, m.runCommand(pad.Command)
					}
				}
			}
		}

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		m.Viewport.Width = msg.Width - 4
		m.Viewport.Height = 10 

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.Spinner, cmd = m.Spinner.Update(msg)
		return m, cmd

	case commandFinishedMsg:
		m.Running = false
		
		// Add output to logs
		if msg.out != "" {
			m.Logs = append(m.Logs, msg.out)
		}

		if msg.err != nil {
			m.Status = fmt.Sprintf("Error: %v", msg.err)
			m.Logs = append(m.Logs, lipgloss.NewStyle().Foreground(lipgloss.Color("#ff5555")).Render(fmt.Sprintf("!!! Error: %v", msg.err)))
		} else {
			m.Status = "Completed"
			m.Logs = append(m.Logs, lipgloss.NewStyle().Foreground(lipgloss.Color("#50fa7b")).Render("--- Finished Successfully ---"))
		}
		
		m.Viewport.SetContent(strings.Join(m.Logs, "\n"))
		m.Viewport.GotoBottom()
		return m, nil
	}

	var vpCmd tea.Cmd
	m.Viewport, vpCmd = m.Viewport.Update(msg)
	cmds = append(cmds, vpCmd)

	return m, tea.Batch(cmds...)
}

type commandFinishedMsg struct {
	err error
	out string
}

func (m Model) runCommand(command string) tea.Cmd {
	return func() tea.Msg {
		var cmd *exec.Cmd
		if runtime.GOOS == "windows" {
			cmd = exec.Command("powershell", "-Command", command)
		} else {
			cmd = exec.Command("sh", "-c", command)
		}

		out, err := cmd.CombinedOutput()
		return commandFinishedMsg{
			err: err,
			out: string(out),
		}
	}
}

func (m Model) View() string {
	if m.Deck == nil {
		return TitleStyle.Render("Atlas Deck") + "\n\n" + StatusStyle.Render(m.Status)
	}

	var s strings.Builder

	// Header
	header := "🚀 " + m.Deck.Name
	if m.Running {
		header += " " + m.Spinner.View()
	}
	s.WriteString(TitleStyle.Render(header))
	s.WriteString("\n\n")

	// Render Pads in a responsive grid
	var pads []string
	for _, pad := range m.Deck.Pads {
		style := PadStyle
		switch pad.Color {
		case "gold":
			style = style.BorderForeground(GoldColor)
		case "cyan":
			style = style.BorderForeground(BaseColor)
		case "red":
			style = style.BorderForeground(RedColor)
		case "green":
			style = style.BorderForeground(GreenColor)
		}

		content := fmt.Sprintf("%s\n\n%s", KeyStyle.Render("["+pad.Key+"]"), pad.Label)
		pads = append(pads, style.Render(content))
	}

	// Dynamic grid
	padWidth := 24 
	cols := m.Width / padWidth
	if cols < 1 {
		cols = 1
	}

	for i := 0; i < len(pads); i += cols {
		end := i + cols
		if end > len(pads) {
			end = len(pads)
		}
		s.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, pads[i:end]...))
		s.WriteString("\n")
	}

	// Footer / Logs
	s.WriteString("\n")
	s.WriteString(StatusStyle.Render(fmt.Sprintf("Last Key: [%s] • Status: %s", m.LastKey, m.Status)))
	s.WriteString("\n")
	s.WriteString(m.Viewport.View())
	s.WriteString("\n" + StatusStyle.Render("Press ctrl+c to quit • ctrl+l to clear logs"))

	return lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, s.String())
}

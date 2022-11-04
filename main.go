package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// main function
func main() {
	if err := tea.NewProgram(model{}).Start(); err != nil {
		fmt.Printf("There was an error: %v\n", err)
		os.Exit(1)
	}
}

type model struct {
	Tabs       []string
	TabContent []string
	activeTab  int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	return "Hello world!"
}

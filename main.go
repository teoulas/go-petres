package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	menuScreen = iota
	gameScreen
)

type mainModel struct {
	currentScreen int
	menuModel     menuModel
	gameModel     gameModel
}

type menuModel struct {
	menuItems   []string
	currentItem int
}

func (m menuModel) Init() tea.Cmd {
	return tea.ClearScreen
}

func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "up":
			if m.currentItem > 0 {
				m.currentItem--
			}
		case "down":
			if m.currentItem < len(m.menuItems)-1 {
				m.currentItem++
			}
		case "enter":
			switch m.currentItem {
			case 0:
				return m, tea.Quit
			case 1:
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m menuModel) View() string {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("#FFF"))

	v := style.Render(logo)
	v += "\n\n"
	v += "Select an option from the menu below:\n"
	for i, item := range m.menuItems {
		if i == m.currentItem {
			v += "> "
		} else {
			v += "  "
		}
		v += item + "\n"
	}
	return v
}

type gameModel struct {
	playfield [21][10]int
}

func (m gameModel) Init() tea.Cmd {
	return tea.ClearScreen
}

func (m gameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m gameModel) View() string {
	return ""
}

const logo = `
████  █████ █████ ████  █████  ███
█   █ █       █   █   █ █     █
█   █ █       █   █   █ █     █
████  ████    █   ████  ████   ███
█     █       █   █   █ █         █
█     █       █   █   █ █         █
█     █████   █   █   █ █████  ███
`

func (m mainModel) Init() tea.Cmd {
	return tea.ClearScreen
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.currentScreen {
	case menuScreen:
		return m.menuModel.Update(msg)
	case gameScreen:
		return m.gameModel.Update(msg)
	}
	return m, nil
}

func (m mainModel) View() string {
	switch m.currentScreen {
	case menuScreen:
		return m.menuModel.View()
	case gameScreen:
		return m.gameModel.View()
	}
	return ""
}

func main() {
	menuModel := menuModel{
		menuItems:   []string{"Start", "Exit"},
		currentItem: 0,
	}
	gameModel := gameModel{}
	main := mainModel{
		currentScreen: menuScreen,
		menuModel:     menuModel,
		gameModel:     gameModel,
	}
	p := tea.NewProgram(main)
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

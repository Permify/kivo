package users

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var userActionsStyle = lipgloss.NewStyle().Margin(1, 2)

type UserAction struct {
	Name string
	Desc string
}

func (i UserAction) Title() string       { return i.Name }
func (i UserAction) Description() string { return i.Desc }
func (i UserAction) FilterValue() string { return i.Name }

type ActionListModel struct {
	state *State
	list  list.Model
}

func ActionList(state *State) ActionListModel {
	items := []list.Item{
		UserAction{Name: ReachableActions[AttachPolicySlug].Name, Desc: ReachableActions[AttachPolicySlug].Desc},
		UserAction{Name: ReachableActions[DetachPolicySlug].Name, Desc: ReachableActions[DetachPolicySlug].Desc},
		UserAction{Name: ReachableActions[AddToGroupSlug].Name, Desc: ReachableActions[AddToGroupSlug].Desc},
		UserAction{Name: ReachableActions[RemoveFromGroupSlug].Name, Desc: ReachableActions[RemoveFromGroupSlug].Desc},
		UserAction{Name: ReachableActions[AttachCustomPolicySlug].Name, Desc: ReachableActions[AttachCustomPolicySlug].Desc},
	}
	var m ActionListModel
	m.state = state
	m.list.Title = "Actions"
	m.list = list.New(items, list.NewDefaultDelegate(), 0, 0)
	return m
}

func (m ActionListModel) Init() tea.Cmd {
	return nil
}

func (m ActionListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			action := m.list.SelectedItem().(UserAction)
			m.state.SetAction(&action)
			return Switch(m.state.FindFlow(), m.list.Width(), m.list.Height())
		}
	case tea.WindowSizeMsg:
		h, v := userActionsStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m ActionListModel) View() string {
	return userActionsStyle.Render(m.list.View())
}
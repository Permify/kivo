package users

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var policiesStyle = lipgloss.NewStyle().Margin(1, 2)

type Policy struct {
	Arn      string
	Name     string
	Document map[string]interface{}
}

func (i Policy) Title() string       { return i.Name }
func (i Policy) Description() string { return i.Arn }
func (i Policy) FilterValue() string { return i.Arn }

type PolicyListModel struct {
	state *State
	list  list.Model
}

func PolicyList(state *State) PolicyListModel {
	var items []list.Item
	policies := []Policy{
		{
			Name: "AdministratorAccess",
			Arn:  "arn:aws:iam::aws:policy/AdministratorAccess",
		},
		{
			Name: "PowerUserAccess",
			Arn:  "arn:aws:iam::aws:policy/PowerUserAccess",
		},
		{
			Name: "ReadOnlyAccess",
			Arn:  "arn:aws:iam::aws:policy/ReadOnlyAccess",
		},
		{
			Name: "SecurityAudit",
			Arn:  "arn:aws:iam::aws:policy/SecurityAudit",
		},
		{
			Name: "NetworkAdministrator",
			Arn:  "arn:aws:iam::aws:policy/NetworkAdministrator",
		},
	}

	for _, policy := range policies {
		items = append(items, Policy{
			Name: policy.Name,
			Arn:  policy.Arn,
		})
	}

	var m PolicyListModel
	m.state = state
	m.list.Title = "Policies"
	m.list = list.New(items, list.NewDefaultDelegate(), 0, 0)
	return m
}

func (m PolicyListModel) Init() tea.Cmd {
	return nil
}

func (m PolicyListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			policy := m.list.SelectedItem().(Policy)
			m.state.SetPolicy(&policy)
			return Switch(m.state.Next(), m.list.Width(), m.list.Height())
		}
	case tea.WindowSizeMsg:
		h, v := policiesStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m PolicyListModel) View() string {
	return policiesStyle.Render(m.list.View())
}

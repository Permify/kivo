package aws

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var servicesStyle = lipgloss.NewStyle().Margin(1, 2)

type Service struct {
	Name string
	Desc string
}

func (i Service) Title() string       { return i.Name }
func (i Service) Description() string { return i.Desc }
func (i Service) FilterValue() string { return i.Name }

type ServicesModel struct {
	user   User
	action string
	list   list.Model
}

func Services(user User, action string) ServicesModel {
	var items []list.Item
	services := []Service{
		{
			Name: "EC2",
			Desc: "Amazon Elastic Compute Cloud",
		},
		{
			Name: "S3",
			Desc: "Amazon Simple Storage Service",
		},
		{
			Name: "RDS",
			Desc: "Amazon Relational Database Service",
		},
		{
			Name: "IAM",
			Desc: "Amazon Identity and Access Management",
		},
		{
			Name: "Lambda",
			Desc: "AWS Lambda",
		},
		{
			Name: "API Gateway",
			Desc: "Amazon API Gateway",
		},
		{
			Name: "CloudFormation",
			Desc: "AWS CloudFormation",
		},
	}

	for _, service := range services {
		items = append(items, Service{
			Name: service.Name,
			Desc: service.Desc,
		})
	}

	var m ServicesModel
	m.user = user
	m.action = action
	m.list.Title = "Services"
	m.list = list.New(items, list.NewDefaultDelegate(), 0, 0)
	return m
}

func (m ServicesModel) Init() tea.Cmd {
	return nil
}

func (m ServicesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			service := m.list.SelectedItem().(Service)
			resources := Resources(m.user, m.action, service.Name)
			return Switch(&resources, m.list.Width(), m.list.Height())
		}
	case tea.WindowSizeMsg:
		h, v := policiesStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m ServicesModel) View() string {
	return servicesStyle.Render(m.list.View())
}

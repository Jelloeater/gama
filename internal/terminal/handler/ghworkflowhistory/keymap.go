package ghworkflowhistory

import (
	teakey "github.com/charmbracelet/bubbles/key"
)

type keyMap struct {
	NextTab     teakey.Binding
	PreviousTab teakey.Binding
	Quit        teakey.Binding
	LaunchTab   teakey.Binding
}

func (k keyMap) ShortHelp() []teakey.Binding {
	return []teakey.Binding{k.PreviousTab, k.NextTab, k.LaunchTab, k.Quit}
}

func (k keyMap) FullHelp() [][]teakey.Binding {
	return [][]teakey.Binding{
		{k.PreviousTab},
		{k.NextTab},
		{k.LaunchTab},
		{k.Quit},
	}
}

var keys = keyMap{
	LaunchTab: teakey.NewBinding(
		teakey.WithKeys("enter"),
		teakey.WithHelp("enter", "Launch the selected option"),
	),
	PreviousTab: teakey.NewBinding(
		teakey.WithKeys("left"),
		teakey.WithHelp("←", "previous tab"),
	),
	NextTab: teakey.NewBinding(
		teakey.WithKeys("right"),
		teakey.WithHelp("→", "next tab"),
	),
	Quit: teakey.NewBinding(
		teakey.WithKeys("q", "ctrl+c"),
		teakey.WithHelp("q", "quit"),
	),
}

func (m *ModelGithubWorkflowHistory) ViewHelp() string {
	return m.Help.View(m.Keys)
}

package information

import (
	teakey "github.com/charmbracelet/bubbles/key"
)

type keyMap struct {
	Refresh teakey.Binding
	NextTab teakey.Binding
	Quit    teakey.Binding
}

func (k keyMap) ShortHelp() []teakey.Binding {
	return []teakey.Binding{k.NextTab, k.Refresh, k.Quit}
}

func (k keyMap) FullHelp() [][]teakey.Binding {
	return [][]teakey.Binding{
		{k.NextTab},
		{k.Refresh},
		{k.Quit},
	}
}

var keys = keyMap{
	Refresh: teakey.NewBinding(
		teakey.WithKeys("r", "R"),
		teakey.WithHelp("r/R", "Empty cache & reload"),
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

func (m *ModelInfo) ViewHelp() string {
	return m.Help.View(m.Keys)
}
package models

const (
	GroupsCollection  = "groups"
	PlayersCollection = "players"
)

type Group struct {
	Name    string   `firebase:"name" json:"name"`
	Players []string `firebase:"players" json:"players"`
	Admins  []string `firebase:"admins" json:"admins"`
	DM      string   `firebase:"dm" json:"dm"`
}

type Players struct {
	Name   string   `firebase:"name" json:"name"`
	Uid    string   `firebase:"uid" json:"uid"`
	Groups []string `firebase:"groups" json:"groups"`
}

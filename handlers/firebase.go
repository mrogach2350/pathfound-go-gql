package handlers

import (
	"cloud.google.com/go/firestore"
	"fmt"
)

type FirebaseHandler struct {
	C *firestore.Client
}

func (h *FirebaseHandler) Test() {
	fmt.Println("fb handler test")
}

func (h *FirebaseHandler) CreateNewPlayer()  {}
func (h *FirebaseHandler) GetAllPlayers()    {}
func (h *FirebaseHandler) GetPlayerById()    {}
func (h *FirebaseHandler) AddPlayerToGroup() {}
func (h *FirebaseHandler) CreateNewGroup()   {}

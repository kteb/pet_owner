package models

import uuid "github.com/satori/go.uuid"

type PetOwner struct {
	ID        uuid.UUID `json:"id" db:"id"`
	PetName   string    `json:"petname" db:"petname" select:"pets.name as petname"`
	OwnerName string    `json:"ownername" db:"ownername" select:"owners.name as ownername"`
}

type PetOwners []PetOwner

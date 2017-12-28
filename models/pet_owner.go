package models

import uuid "github.com/satori/go.uuid"

type PetOwner struct {
	ID        uuid.UUID `json:"id" db:"id"`
	PetName   string    `json:"name" db:"pet_name"`
	OwnerName string    `json:"owner" db:"owner_name"`
}

type PetOwners []PetOwner

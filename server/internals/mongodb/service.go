package mongodb

import "github.com/thebravebyte/findr/internals"

// LinkerStore is an interface that defines and implements the methods that can be used to interact with the database.
type FindrStore interface {
	CreateAccount(std internals.Students) (bool, bool, error) 
}
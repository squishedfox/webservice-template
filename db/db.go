package db

import "github.com/squishedfox/webservice-prototype/models"

// PersonResourceManager provides an abstract interface for managing Person Resources to a database provider such as
// postgres, ms sql server, couchdb, mongodb, redis, dynamodb, etc.
type PersonResourceManager interface { // alternatively this can be named to PersonEntityDb or PersonEntityManager, DbContext, etc.
	// CreatePerson function puts a new person resource into the database and returns the ID of the newly
	// created Person Resource. if there is an error while attempting to create the Person resource it is
	// returned with a nil for the ID.
	CreatePerson(person models.Person) (interface{}, error)

	// DeletePerson deletes a Person resource from the target database system. If there is an error attempting
	// to delete the resource the error is returned. If the resource does not exist no error is returned.
	DeletePerson(id interface{}) error

	// UpdatePerson modifies and updates a person resource. If there is an error attempting to update the
	// resource or a resource could not be found an error is returned.
	UpdatePerson(id interface{}, person models.Person) error

	// GetById fetches a Person resource by it's identifier. If no resource is found then nil, nil is returned
	// as a successfully "failed" attempt. If there is an issue communicating with the database system the error
	// is returned and nil for the resource.
	GetById(id interface{}) (*models.Person, error)

	// Get fetches all Person resources from target database/resource storage. If none are found an empty slice
	// is returned. If there is an error fetching one or more recrods the error is immediately returned at the
	// opperation is cancelled.
	Get() ([]*models.Person, error)

	// TODO: add query availability as well so we can search for resources based on properties
}

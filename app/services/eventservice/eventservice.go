package eventservice

import (
	"github.com/DLzer/go-gin-api-boilerplate/app/domain/event"
)

// EventRepo is the defined interface
// of functions that need to be implemented
type EventRepo interface {
	CreateEvent(event *event.Event) (*event.Event, error)
	GetAllEvents() ([]event.Event, error)
	GetEventById(id string) (event.Event, error)
	DeleteEventById(id string) (bool, error)
}

// The service
type eventService struct {
	er EventRepo
}

// NewEventService expects an object implementing
// EventRepo and returns a new eventService object
func NewEventService(er EventRepo) *eventService {
	return &eventService{
		er: er,
	}
}

// Creates a new event
func (es *eventService) CreateEvent(event *event.Event) (*event.Event, error) {
	return es.er.CreateEvent(event)
}

// Gets all events
func (es *eventService) GetAllEvents() ([]event.Event, error) {
	return es.er.GetAllEvents()
}

// Gets an event by its Id
func (es *eventService) GetEventById(id string) (event.Event, error) {
	return es.er.GetEventById(id)
}

// Deletes an event by its Id
func (es *eventService) DeleteEventById(id string) (bool, error) {
	return es.er.DeleteEventById(id)
}

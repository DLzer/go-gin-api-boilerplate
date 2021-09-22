package controllers

import (
	"net/http"

	"github.com/DLzer/go-gin-api-boilerplate/app/domain/event"
	"github.com/gin-gonic/gin"
)

// The output object
type EventOutput struct {
	ID       uint64 `json:"id"`
	Type     string `json:"type"`
	Time     string `json:"time"`
	Identity string `json:"identity"`
}

// The input object
type EventInput struct {
	Type     string `json:"type"`
	Time     string `json:"time"`
	Identity string `json:"identity"`
}

// The service interface
type EventService interface {
	CreateEvent(event *event.Event) (*event.Event, error)
	GetAllEvents() ([]event.Event, error)
	GetEventById(id string) (event.Event, error)
	DeleteEventById(id string) (bool, error)
}

// The controller object
type eventController struct {
	es EventService
}

// Return new event controller object
func NewEventController(es EventService) *eventController {
	return &eventController{es: es}
}

// Create new event
func (ctl *eventController) PostEvent(c *gin.Context) {
	// Read input
	var eventInput EventInput
	if err := c.ShouldBindJSON(&eventInput); err != nil {
		HTTPRes(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// Convert
	e := ctl.inputToEvent(eventInput)

	// Create event
	if _, err := ctl.es.CreateEvent(&e); err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// If event created return a structured response
	eventOutput := ctl.mapToEventOutput(&e)
	HTTPRes(c, http.StatusOK, "Event Published", eventOutput)
}

// Get all events
func (ctl *eventController) GetAllEvents(c *gin.Context) {

	events, err := ctl.es.GetAllEvents()
	if err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	HTTPRes(c, http.StatusOK, "Event List", events)
}

// Get event by ID
func (ctl *eventController) GetEventById(c *gin.Context) {

	id := c.Param("id")

	event, err := ctl.es.GetEventById(id)
	if err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	HTTPRes(c, http.StatusOK, "Event", event)

}

// Delete a single Event by ID
func (ctl *eventController) DeleteEventById(c *gin.Context) {

	id := c.Param("id")

	result, err := ctl.es.DeleteEventById(id)
	if err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), false)
		return
	}

	HTTPRes(c, http.StatusOK, "Event Deleted", result)

}

// Private Methods
func (ctl *eventController) inputToEvent(input EventInput) event.Event {
	return event.Event{
		Type:     input.Type,
		Time:     input.Time,
		Identity: input.Identity,
	}
}

func (ctl *eventController) mapToEventOutput(e *event.Event) *EventOutput {
	return &EventOutput{
		ID:       e.ID,
		Type:     e.Type,
		Time:     e.Time,
		Identity: e.Identity,
	}
}

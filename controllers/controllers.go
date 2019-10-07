// Package controllers provides handlers that can be mapped to routes and defines interfaces that define
// what is required by the handlers.
package controllers

import (
	"fmt"

	"github.com/makeict/MESSforMakers/models"
	"github.com/makeict/MESSforMakers/session"
	"github.com/makeict/MESSforMakers/util"
	"github.com/makeict/MESSforMakers/views"
)

// Users interface defines the methods that a Users model must fulfill. Allows mocking with a fake database for testing.
type Users interface {
	Get(int) (*models.User, error)
	GetAll(int, int) ([]models.User, error)
	Create(*models.User) error
	Update(*models.User) error
	Delete(*models.User) error
}

// Controller is a struct Struct to store pointer to cookiestore, database, and logger and any other things common to many controllers
type Controller struct {
	CookieStore *session.CookieStore
	Users       Users
	Logger      *util.Logger
	AppConfig   *util.Config
}

// method to create a new struct and store the information from the app, passed as args
// Requiring the information passed as args avoids imports loop
// the general controller constructor will be called by the specific controller constructors
// the specific controller constructors can then initialize and embed their own templates.
func (c *Controller) setup(cfg *util.Config, cs *session.CookieStore, um Users, l *util.Logger) {
	c.CookieStore = cs
	c.Users = um
	c.Logger = l
	c.AppConfig = cfg
}

// DefaultData ia the method to generate required default template data and return template object
func (c *Controller) DefaultData() (*views.TemplateData, error) {
	td := &views.TemplateData{}
	td.Root = fmt.Sprintf("http://%s:%d/", c.AppConfig.App.Host, c.AppConfig.App.Port)
	return td, nil
}

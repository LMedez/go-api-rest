package album

import (
	"github.com/LMedez/go-api-rest/internal/errors"
	"github.com/LMedez/go-api-rest/pkg/log"
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"net/http"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}

	r.Get("/albums/<id>", res.get)

	r.Use(authHandler)

	// the following endpoints require a valid JWT
	r.Post("/albums", res.create)
	r.Put("/albums/<id>", res.update)
	r.Delete("/albums/<id>", res.delete)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) get(c *routing.Context) error {
	album, err := r.service.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(album)
}

func (r resource) create(c *routing.Context) error {
	var input CreateAlbumRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	album, err := r.service.Create(c.Request.Context(), input)
	if err != nil {
		return err
	}

	return c.WriteWithStatus(album, http.StatusCreated)
}

func (r resource) update(c *routing.Context) error {
	var input UpdateAlbumRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	album, err := r.service.Update(c.Request.Context(), c.Param("id"), input)
	if err != nil {
		return err
	}

	return c.Write(album)
}

func (r resource) delete(c *routing.Context) error {
	album, err := r.service.Delete(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(album)
}

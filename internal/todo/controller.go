package todo

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// Controller defines the interface for admin-related HTTP handlers.
type Controller interface {
	GetAdmins(c *fiber.Ctx) error
	GetAdminByID(c *fiber.Ctx) error
	GetAdminByEmail(c *fiber.Ctx) error
	AddAdmin(c *fiber.Ctx) error
	UpdateAdmin(c *fiber.Ctx) error
	DeleteAdmin(c *fiber.Ctx) error
}

// controller implements the Controller interface.
type controller struct {
	service Service
}

// NewController initializes and returns a new controller instance.
func NewController(s Service) Controller {
	return &controller{
		service: s,
	}
}

// GetAdmins godoc
// @Summary      Get all admins
// @Description  Retrieves all admins from the database.
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Success      200  {object}  []entities.Admin
// @Failure		 500  {object} ErrorResponse
// @App       /api/v1/admin/ [get]
// GetAdmins fetches and returns all admins.
func (ctr *controller) GetAdmins(c *fiber.Ctx) error {
	fetched, err := ctr.service.FetchAdmins()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(CreateErrorResponse(err))
	}
	return c.JSON(fetched)
}

// GetAdminByID godoc
// @Summary      Get an admin by ID
// @Description  Retrieves an admin by its ID.
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Admin ID"
// @Success      200  {object}  entities.Admin
// @Failure		 500  {object} ErrorResponse
// @App       /api/v1/admin/{id} [get]
// GetAdminByID fetches and returns a specific admin by its ID.
func (ctr *controller) GetAdminByID(c *fiber.Ctx) error {
	id := c.Params("id")
	fetched, err := ctr.service.FetchAdminByID(id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(CreateErrorResponse(err))
	}
	return c.JSON(fetched)
}

// GetAdminByEmail godoc
// @Summary      Get an admin by email
// @Description  Retrieves an admin by its email.
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Param        email   path      string  true  "Admin email"
// @Success      200  {object}  entities.Admin
// @Failure		 500  {object} ErrorResponse
// @App       /api/v1/admin/email/{email} [get]
// GetAdminByEmail fetches and returns a specific admin by its email.
func (ctr *controller) GetAdminByEmail(c *fiber.Ctx) error {
	email := c.Params("email")
	fetched, err := ctr.service.FetchAdminByEmail(email)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(CreateErrorResponse(err))
	}
	return c.JSON(fetched)
}

// AddAdmin godoc
// @Summary      Add a new admin
// @Description  Creates a new admin in the database.
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Param        admin  body      entities.Admin  true  "Admin data"
// @Success      200   {object}  entities.Admin
// @Failure		 400  {object} ErrorResponse
// @Failure		 500  {object} ErrorResponse
// @App       /api/v1/admin/ [post]
// AddAdmin creates a new admin.
func (ctr *controller) AddAdmin(c *fiber.Ctx) error {
	var body entities.Admin
	err := c.BodyParser(&body)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(CreateErrorResponse(err))
	}
	if err := ctr.service.InsertAdmin(&body); err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(CreateErrorResponse(err))
	}
	return c.JSON(&body)
}

// UpdateAdmin godoc
// @Summary      Update an existing admin
// @Description  Updates an existing admin in the database.
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Param        admin  body      entities.Admin  true  "Updated admin data"
// @Success      200   {object}  entities.Admin
// @Failure		 400  {object} ErrorResponse
// @Failure		 500  {object} ErrorResponse
// @App       /api/v1/admin/ [put]
// UpdateAdmin updates an existing admin.
func (ctr *controller) UpdateAdmin(c *fiber.Ctx) error {
	var body entities.Admin
	err := c.BodyParser(&body)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(CreateErrorResponse(err))
	}
	if err := ctr.service.UpdateAdmin(&body); err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(CreateErrorResponse(err))
	}
	return c.JSON(&body)
}

// DeleteAdmin godoc
// @Summary      Delete an admin
// @Description  Removes an admin by its ID.
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "Admin ID"
// @Success      200   {object}  string  "Admin ID"
// @Failure		 500  {object} ErrorResponse
// @App       /api/v1/admin/{id} [delete]
// DeleteAdmin removes an admin by its ID.
func (ctr *controller) DeleteAdmin(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := ctr.service.RemoveAdmin(id); err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(CreateErrorResponse(err))
	}
	return c.JSON(id)
}

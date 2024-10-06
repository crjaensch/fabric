package restapi

import (
	"fmt"
	"github.com/danielmiessler/fabric/db"
	"github.com/labstack/echo/v4"
	"net/http"
)

// StorageHandler defines the handler for storage-related operations
type StorageHandler[T any] struct {
	storage db.Storage[T]
}

// NewStorageHandler creates a new StorageHandler
func NewStorageHandler[T any](e *echo.Echo, entityType string, storage db.Storage[T]) (ret *StorageHandler[T]) {
	ret = &StorageHandler[T]{storage: storage}
	e.GET(fmt.Sprintf("/%s/names", entityType), ret.GetNames)
	e.DELETE(fmt.Sprintf("/%s/:name", entityType), ret.Delete)
	e.GET(fmt.Sprintf("/%s/exists/:name", entityType), ret.Exists)
	e.PUT(fmt.Sprintf("/%s/rename/:oldName/:newName", entityType), ret.Rename)
	e.POST(fmt.Sprintf("/%s/save/:name", entityType), ret.Save)
	e.GET(fmt.Sprintf("/%s/load/:name", entityType), ret.Load)
	e.GET(fmt.Sprintf("/%s/list", entityType), ret.ListNames)
	return
}

func (h *ContextsHandler) Get(c echo.Context) error {
	name := c.Param("name")
	context, err := h.contexts.Get(name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, context)
}

// GetNames handles the GET /storage/names route
func (h *StorageHandler[T]) GetNames(c echo.Context) error {
	names, err := h.storage.GetNames()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, names)
}

// Delete handles the DELETE /storage/:name route
func (h *StorageHandler[T]) Delete(c echo.Context) error {
	name := c.Param("name")
	err := h.storage.Delete(name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

// Exists handles the GET /storage/exists/:name route
func (h *StorageHandler[T]) Exists(c echo.Context) error {
	name := c.Param("name")
	exists := h.storage.Exists(name)
	return c.JSON(http.StatusOK, exists)
}

// Rename handles the PUT /storage/rename/:oldName/:newName route
func (h *StorageHandler[T]) Rename(c echo.Context) error {
	oldName := c.Param("oldName")
	newName := c.Param("newName")
	err := h.storage.Rename(oldName, newName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

// Save handles the POST /storage/save/:name route
func (h *StorageHandler[T]) Save(c echo.Context) error {
	name := c.Param("name")
	content := c.Body()
	err := h.storage.Save(name, content)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

// Load handles the GET /storage/load/:name route
func (h *StorageHandler[T]) Load(c echo.Context) error {
	name := c.Param("name")
	content, err := h.storage.Load(name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, content)
}

// ListNames handles the GET /storage/list route
func (h *StorageHandler[T]) ListNames(c echo.Context) error {
	err := h.storage.ListNames()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

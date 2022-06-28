package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ResponseHelper ...
type ResponseHelper struct {
	C        *gin.Context
	Data     interface{}
	Code     int
}

type HTTPResponseHelper struct {}

// setResponse ...
// Set response data.
func (h *HTTPResponseHelper) setResponse(c *gin.Context, data interface{}, code int) ResponseHelper {
	return ResponseHelper{c, data, code}
}

// SendBadRequest ...
// Send bad request response to consumers.
func (h *HTTPResponseHelper) SendBadRequest(c *gin.Context, data interface{}) {
	res := h.setResponse(c,  data, http.StatusBadRequest)

	h.sendResponse(res)
}

// SendSuccess ...
// Send success response to consumers.
func (h *HTTPResponseHelper) SendSuccess(c *gin.Context, data interface{}) {
	res := h.setResponse(c,  data, http.StatusOK)

	h.sendResponse(res)
}

// sendResponse ...
// Send response
func (h *HTTPResponseHelper) sendResponse(res ResponseHelper) {
	res.C.JSON(res.Code, gin.H{"data": res.Data})
	return
}
package helper

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Success will write a default template response when returning a success response
func Success(gc *gin.Context, status int, data interface{}, msg string) {

	w := gc.Writer
	res := map[string]interface{}{
		"message": msg,
		"status":  "success",
		"data":    data,
	}
	js, err := json.Marshal(res)
	if err != nil {
		resp := map[string]interface{}{
			"error": fmt.Sprintf("%s", err),
		}
		js, _ = json.Marshal(resp)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	w.Write(js)
}

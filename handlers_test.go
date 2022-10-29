package simple_http_server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/faisal097/simple_http_server/models"
)

// type tServer struct {}

// func (ts *tServer) Get(w http.ResponseWriter, r *http.Request) {

// }

func Test_get_handler(t *testing.T) {
	var apis Apis
	apis = &Server{}
	t.Run("returns all employees", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/get", nil)
		response := httptest.NewRecorder()
		get_handler(apis, response, request)

		got := response.Body.String()
		want, _ := json.Marshal(models.PLACE_HOLDER_EMPLOYEES)

		if got != string(want) {
			t.Errorf("got %q, want %q", got, string(want))
		}
	})
}

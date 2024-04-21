package httpsrv

import (
	"encoding/json"
	"fmt"
	"goapp/pkg/util"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type hexString struct {
	HexValue string `json:"hex"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: nil,
}

func (s *Server) handlerHexWS(w http.ResponseWriter, r *http.Request) {

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.error(w, http.StatusInternalServerError, fmt.Errorf("failed to upgrade connection: %w", err))
		return
	}
	defer c.Close()

	hexVal := hexString{HexValue: util.RandString(10)}
	data, _ := json.Marshal(hexVal)
	err = c.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			log.Printf("failed to write message: %v\n", err)
		}
		return
	}
}

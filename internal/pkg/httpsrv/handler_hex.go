package httpsrv

import (
	"html/template"
	"net/http"
)

func (s *Server) handlerHex(w http.ResponseWriter, r *http.Request) {
	template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
    var ws;
    var print = function(message) {
        var d = document.createElement("div");
        d.textContent = message;
        output.appendChild(d);
        output.scroll(0, output.scrollHeight);
    };
    document.getElementById("getHex").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        
        ws.onclose = function(evt) {
            ws = null;
        }
        ws.onmessage = function(evt) {
            print(evt.data);
        }
        return false;
    };
});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<button id="getHex">Get Hex</button>
</td><td valign="top" width="50%">
<div id="output" style="max-height: 70vh;overflow-y: scroll;"></div>
</td></tr></table>
</body>
</html>
`)).Execute(w, "ws://"+r.Host+"/hex/ws")
}

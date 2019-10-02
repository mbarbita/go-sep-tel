window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;

    var wsOpen = function(evt) {
        output.value = evt.data;

        if (ws) {
            return false;
        }
        ws = new WebSocket("ws://{{.}}/msg");
        ws.onopen = function(evt) {
            console.log("OPEN");
            // ws.send(inputName.value);
            output.value = "Connected"
        }

        ws.onclose = function(evt) {
            console.log("CLOSE");
            output.value = "NOT Connected"
            ws = null;
        }

        ws.onmessage = function(evt) {
            console.log("MESSAGE:");
            // console.log(evt.data);
            output.value = evt.data;
        }

        ws.onerror = function(evt) {
            console.log("ERROR: " + evt.data);
            output.value = "Error. Refresh browser";

        }
        return false;
    };

    input.onkeyup = function(evt) {
      // console.log("SEND:");
      console.log(input.value);
      ws.send(input.value);
    };

    input.onclick = function(evt) {
      output.value="";
      input.value ="";
    };

    return wsOpen(evt);
});

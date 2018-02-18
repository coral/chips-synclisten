var ws = new WebSocket("ws://" + location.host + '/ws');

ws.onmessage = function(evt){
    var message = JSON.parse(evt.data)

    switch(message.Message) {
        case "Pong":
            console.log("got pong!");
            break;
        case "Compodata":
            var compodata = JSON.parse(message.Data);
            loadCompo(compodata);
            break;
        case "Start":
            var compodata = JSON.parse(message.Data);
            startCompo(compodata);
            break;
        default:
            
    }
}

function sendMessage(msg){
    // Wait until the state of the socket is not ready and send the message when it is...
    waitForSocketConnection(ws, function(){
        
        ws.send(msg);
    });
}

// Make the function wait until the connection is made...
function waitForSocketConnection(socket, callback){
    setTimeout(
        function () {
            if (socket.readyState === 1) {
                console.log("Connection is made")
                if(callback != null){
                    callback();
                }
                return;
                
            } else {
                console.log("wait for connection...")
                waitForSocketConnection(socket, callback);
            }
            
        }, 5); // wait 5 milisecond for the connection...
    }
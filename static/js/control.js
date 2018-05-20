var ws = new WebSocket("ws://" + location.host + '/ws');

ws.onmessage = function(evt){

    console.log(evt);
    var message = JSON.parse(evt.data)

    $('#messagebus').prepend("<p>" + message.Message + "   " + message.Data + "</p>");
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

//DOWNLOAD
$( "#download" ).click(function() {
    sendMessage(JSON.stringify({"function":"download", "message":getCompo()}));
});
$( "#fetch" ).click(function() {
    sendMessage(JSON.stringify({"function":"fetch", "message":getCompo()}));
});
$( "#nointro" ).click(function() {
    sendMessage(JSON.stringify({"function":"nointro", "message":getCompo()}));
});
$( "#start" ).click(function() {
    sendMessage(JSON.stringify({"function":"start", "message":getCompo()}));
});

function getCompo() {
    return $('#compo').val();
}
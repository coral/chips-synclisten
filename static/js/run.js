//HELLO JOHNFN
//YOU WILL HATE ME
//I AM GOING TO USE A GLOBAL STATE
//KILL ME

function setup() {
    
    
}

function draw() {
    
    if (_.invoke(song, "isPlaying") != undefined) {
        if( song.isPlaying())
        {
            var duration = song.currentTime() / song.buffer.duration * 100;
            $(".progress-bar").css("width", duration +"%");
        }
    }
}

sendMessage(JSON.stringify({"function":"fetch", "message":"43"}));


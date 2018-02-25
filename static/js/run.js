//HELLO JOHNFN
//YOU WILL HATE ME
//I AM GOING TO USE A GLOBAL STATE
//KILL ME

var canvas;
var fft;
var numBars = 256;


function setup() {
    canvas = createCanvas(1280,720);
    canvas.class("bgrender")
    fft = new p5.FFT();
    fft.waveform(numBars);
    fft.smooth(0.85);
}

function draw() {
    
    if (_.invoke(song, "isPlaying") != undefined) {
        if( song.isPlaying())
        {
            var duration = song.currentTime() / song.buffer.duration * 100;
            $(".progress-bar").css("width", duration +"%");

        }

        if(typeof fft != "undefined") {
            background(colorBg);
            var spectrum = fft.analyze();
            noStroke();
            fill(colorDarkVibrant);
            for(var i = 0; i < numBars; i++) {
                var x = map(i, 0, numBars, 0, width);
                var h = -height + map(spectrum[i], 0, 255, height, 0);
                rect(x*2.5, height, (width / numBars)*2.5, h);
            }


        }
    }
}



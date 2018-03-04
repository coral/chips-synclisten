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

            var secondsLeft = Math.round(song.buffer.duration);
            var minutes = Math.floor(secondsLeft / 60);
            var seconds = secondsLeft - minutes * 60;
            var totaltime = str_pad_left(minutes,'0',2)+':'+str_pad_left(seconds,'0',2);
            $(".time-total").text(totaltime);

            var secondsLeft = Math.round( song.currentTime());
            var minutes = Math.floor(secondsLeft / 60);
            var seconds = secondsLeft - minutes * 60;
            var countup = str_pad_left(minutes,'0',2)+':'+str_pad_left(seconds,'0',2);
            $(".time-progressed").text(countup);
            

        }

        if(typeof fft != "undefined") {

            clear();
            var spectrum = fft.analyze();
            noStroke();
            
            for(var i = 0; i < numBars; i++) {
                var pos = i / numBars;
                fill(lerpColor(color(colorDarkVibrant), color(colorVibrant), pos*1.2));
                var x = map(i, 0, numBars, 0, width);
                var h = -height + map(spectrum[i], 0, 255, height, 0);
                rect(x*2.5, height, (width / numBars)*2.3, h);
            }


        }
    }
}

function str_pad_left(string,pad,length) {
    return (new Array(length+1).join(pad)+string).slice(-length);
}
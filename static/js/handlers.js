
var playlist = [];
var totalLength;
var songNumber = 1;
var fileBucket = '';
var bgImage;
var song;

var colorBg = "#cbcbcb";
var colorDarkVibrant = "#424040";
var colorVibrant = "#424040";

function loadCompo(c) {
    console.log(c);

    $('#compo').text("CHIPS COMPO: " + c.cdata.compo.name)
    $('.entry-image').css("background-image", "url("+c.cdata.compo.primary_image+")");
    $('.componame').text("CHIPS COMPO: " + c.cdata.compo.name)

    updateColors(c);
}

function startCompo(c) {
    var compo = c;
    $('#compo').text("CHIPS COMPO: " + c.cdata.compo.name)
    $('.entry-image').css("background-image", "url("+c.cdata.compo.primary_image+")"); 

    
    fileBucket = 'tmp/compos/' + compo.cdata.compo.id + '/';
    
    _.forEach(c.cgen.songs, function(entry) {
        if(entry.type == "song" && !entry.is_joke) {
            playlist.push(entry);
        }
    });
    totalLength = playlist.length;



    playSong();
    $('#entry-title').text(playlist[0].title);
    $('#entry-description').text(playlist[0].description);


    var startAnimation = anime.timeline();
    startAnimation
   .add({
        targets: '.pre',
        opacity: 0,
        duration: 1000,
        offset: '+=5000',
        easing: 'linear'
    }).add({
        targets: '.entry-container',
        opacity: 1,
        duration: 2000,
        easing: 'linear'
    })

}

function playSong() {

    //Check if playlist is empty and if end compo
    if(_.isEmpty(playlist)) {
        endCompo();
        return;
    }

    var parts = playlist[0].uploaded_url.split('/');
    var filename = parts.pop() || parts.pop();


    song = loadSound(fileBucket + filename, function(song){

        
        queueNextUp(playlist[0].title, song, function(song){

    
            song.play();
            playlist.shift();
    
            song.onended(function(){
                playSong();
            });

        });

    });

}

function endCompo() {
    var endAnimation = anime.timeline();
    endAnimation
   .add({
        targets: '.entry-container',
        opacity: 0,
        duration: 3000,
        easing: 'linear'
    })
    .add({
        targets: '.end',
        opacity: 1,
        duration: 1000,
        easing: 'linear'
    })
}

function tickSong() {
    
    var songAppend = "";
    if(typeof playlist[1] === 'undefined') {
    } else {
        songAppend = "<b>Next:</b> " + playlist[1].title + " - ";
    }

    $('#entry-counter').html(songAppend + "Song " +  songNumber + "/" + totalLength)
    songNumber++;
}

function updateColors(c) {
    
    var bucket = 'tmp/compos/' + c.cdata.compo.id + '/';
    var parts = c.cdata.compo.primary_image.split('/');
    var primaryImage = parts.pop() || parts.pop();
    var img = document.createElement('img');
    img.setAttribute('src', bucket + primaryImage)

    bgImage = loadImage(bucket + primaryImage.replace(/\.[^/.]+$/, "") + "_blur.jpg");

    img.addEventListener('load', function() {
        var vibrant = new Vibrant(img);
        var swatches = vibrant.swatches()
        var newColors = [];
        console.log(swatches);
        for (var swatch in swatches) {
            if (swatches.hasOwnProperty(swatch) && swatches[swatch]) {
                switch(swatch) {
                    case "LightMuted":
                        newColors.push({Name: 'muted', Color: swatches[swatch].getHex()})
                        colorBg = swatches[swatch].getHex();
                        break;
                    case "LightVibrant":
                        newColors.push({Name: 'muted', Color: swatches[swatch].getHex()})
                        colorBg = swatches[swatch].getHex();
                        break;
                    case "Vibrant":
                        newColors.push({Name: 'vibrant', Color: swatches[swatch].getHex()})
                        colorVibrant = swatches[swatch].getHex();
                        break;
                    case "DarkVibrant":
                        newColors.push({Name: 'darkvibrant', Color: swatches[swatch].getHex()})
                        colorDarkVibrant = swatches[swatch].getHex();
                        break;
                    case "DarkMuted":
                        newColors.push({Name: 'darkmuted', Color: swatches[swatch].getHex()})
                        
                        break;
                    default: 
                        break;
                }
            }
        }

        for (var color in newColors) {
            d = newColors[color];

            anime({
                targets: "." + d.Name + "-fg",
                color: d.Color,
                duration: 1000,
                easing: 'linear'
            })

            anime({
                targets: "." + d.Name + "-bg",
                backgroundColor: d.Color,
                duration: 1000,
                easing: 'linear'
            })
        }
            
    });
}

function queueNextUp(title, song, callback) {
    $('#nextentry-title').html("<b>NEXT UP:</b> " + title);
    var nextUp = anime.timeline();
    nextUp.add({
        targets: '.nextentry',
        translateX: 1280,
        duration: 3000,
        easing: 'easeInOutQuart',
        complete: function(anim) {
            $('#entry-title').text(playlist[0].title);
            $('#entry-description').text(playlist[0].description);
            tickSong();
        }
    })
    .add({
        targets: '.progress',
        opacity: 0,
        duration: 1000,
    })
    .add({
        targets: '.nextentry',
        translateX: 2560,
        duration: 3000,
        easing: 'easeInOutQuart',
        offset: '+=10000',
        complete: function(anim) {
           anime({
                targets: '.nextentry',
                translateX: 0,
                duration: 0
            })
            callback(song);
        }
    })
    .add({
        targets: '.progress',
        opacity: 0.9,
        duration: 2000,
    })
}
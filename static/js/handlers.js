var playlist = [];
var totalLength;
var songNumber = 1;
var fileBucket = '';
var compoName = "";
var bgImage;
var song;
var strings;

var activeCompo;

var intro;
var outro;

var colorBg = "#000000";
var colorDarkVibrant = "#000000";
var colorVibrant = "#000000";

function loadCompo(c) {
    $.getJSON("assets/strings.json", function (data) {
        strings = data;
    });
    compoName = c.cdata.compo.name;
    activeCompo = c;
    $('#compo').text("CHIPS COMPO: " + c.cdata.compo.name);
    $('.entry-image').css("background-image", "url(" + c.cdata.compo.primary_image + ")");
    $('.componame').text("CHIPS COMPO: " + c.cdata.compo.name)

    var nm = JSON.stringify({
        "function": "postsonglist",
        "message": ""
    });
    sendMessage(nm);

    updateColors(c);



    intro = loadSound("assets/intro.mp3", function (song) {
        song.play()

        var startAnimation = anime.timeline();
        startAnimation
            .add({
                targets: '.startingsoon',
                opacity: 0,
                duration: 1000,
                complete: function (anim) {
                    $('.startingsoon').remove();
                },
                easing: 'linear'
            }).add({
                targets: '.cd',
                opacity: 1,
                duration: 2000,
                easing: 'linear'
            })

        song.onended(function () {
            startCompo(activeCompo)
        });
    })
}

function startCompo(c) {
    var compo = c;
    $('#compo').text("CHIPS COMPO: " + c.cdata.compo.name)
    $('.entry-image').css("background-image", "url(" + c.cdata.compo.primary_image + ")");


    fileBucket = 'tmp/compos/' + compo.cdata.compo.id + '/';

    _.forEach(c.cgen.songs, function (entry) {
        if (entry.type == "song" && !entry.is_joke) {
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
    if (_.isEmpty(playlist)) {
        endCompo();
        return;
    }

    var parts = playlist[0].uploaded_url.split('/');
    var filename = parts.pop() || parts.pop();


    song = loadSound(fileBucket + filename, function (song) {


        queueNextUp(playlist[0].title, song, function (song) {

            var nm = JSON.stringify({
                "function": "playsong",
                "message": playlist[0].id.toString()
            });
            sendMessage(nm);

            $(".bgrender").css('opacity', '1.0');
            song.play();
            playlist.shift();

            song.onended(function () {
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
    outro = loadSound("assets/outro.mp3", function (song) {
        $(".bgrender").css('opacity', '0.0');
        song.play()
        song.onended(function () {

        });
    })
}

function tickSong() {

    var songAppend = "";
    if (typeof playlist[1] === 'undefined') {} else {
        songAppend = "<b>Next:</b> " + playlist[1].title + " - ";
    }

    $('#entry-counter').html(songAppend + "Song " + songNumber + "/" + totalLength)
    songNumber++;
}

function updateColors(c) {

    var bucket = 'tmp/compos/' + c.cdata.compo.id + '/';
    var parts = c.cdata.compo.primary_image.split('/');
    var primaryImage = parts.pop() || parts.pop();
    var blurURL = bucket + primaryImage.replace(/\.[^/.]+$/, "") + "_blur.jpg"
    $(".bgimage").css('background-image', 'url(' + blurURL + ')');

    anime({
        targets: ".bgimage",
        opacity: 1,
        duration: 4000,
        easing: 'linear'
    })

    var img = document.createElement('img');
    img.setAttribute('src', bucket + primaryImage)

    bgImage = loadImage(blurURL);

    img.addEventListener('load', function () {
        var vibrant = new Vibrant(img);
        var swatches = vibrant.swatches()
        var newColors = [];
        console.log(swatches);
        for (var swatch in swatches) {
            if (swatches.hasOwnProperty(swatch) && swatches[swatch]) {
                switch (swatch) {
                    case "LightMuted":
                        newColors.push({
                            Name: 'muted',
                            Color: swatches[swatch].getHex()
                        })
                        colorBg = swatches[swatch].getHex();
                        break;
                    case "LightVibrant":
                        newColors.push({
                            Name: 'muted',
                            Color: swatches[swatch].getHex()
                        })
                        colorBg = swatches[swatch].getHex();
                        break;
                    case "Vibrant":
                        newColors.push({
                            Name: 'vibrant',
                            Color: swatches[swatch].getHex()
                        })
                        colorVibrant = swatches[swatch].getHex();
                        break;
                    case "DarkVibrant":
                        newColors.push({
                            Name: 'darkvibrant',
                            Color: swatches[swatch].getHex()
                        })
                        colorDarkVibrant = swatches[swatch].getHex();
                        break;
                    case "DarkMuted":
                        newColors.push({
                            Name: 'darkmuted',
                            Color: swatches[swatch].getHex()
                        })

                        break;
                    default:
                        break;
                }
            }
        }

        if (tinycolor(colorBg).getLuminance() < 0.1 && tinycolor(colorDarkVibrant).getLuminance() < 0.1) {
            colorDarkVibrant = tinycolor(colorVibrant).brighten(20).toString();
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
    $('#nextentry-title').html("<b>NOW PLAYING:</b> " + title);
    var nextUp = anime.timeline();
    nextUp.add({
            targets: '.nextentry',
            translateX: 1280,
            duration: 3000,
            easing: 'easeInOutQuart',
            complete: function (anim) {
                $('#entry-title').text(playlist[0].title);
                $('#entry-description').text(playlist[0].description);
                playTTS(title);
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
            complete: function (anim) {
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

function playTTS(title) {

    st = _.sample(strings.introductions)
    var m = compoName + "," + st + " " + title;
    var request = new XMLHttpRequest();
    request.open("POST", "/tts", true);
    request.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    request.responseType = "blob";
    request.onreadystatechange = function () {
        if (request.readyState == XMLHttpRequest.DONE && request.status == 200) {
            var blob = new Blob([this.response], {
                type: "audio/mpeg"
            });
            var url = URL.createObjectURL(blob);
            var tts = loadSound(url, function (tts) {
                $(".bgrender").css('opacity', '0.0');
                tts.setVolume(2.5);
                tts.play();
            });
        }
    }
    request.send("message=" + m);


}

var playlist = [];
var fileBucket = '';
var song;

function loadCompo(c) {
    console.log(c);

    $('#compo').text("CHIPS COMPO: " + c.compo.name)
    $('.entry-image').css("background-image", "url("+c.compo.primary_image+")"); 
    $('.componame').text("CHIPS COMPO: " + c.compo.name)
}

function startCompo(c) {
    var compo = c;
    $('#compo').text("CHIPS COMPO: " + c.compo.name)
    $('.entry-image').css("background-image", "url("+c.compo.primary_image+")"); 

    
    fileBucket = 'tmp/compos/' + compo.compo.id + '/';
    
    _.forEach(c.entries, function(entry) {
        if(entry.type == "song" && !entry.is_joke) {
            playlist.push(entry);
        }
    });
    playlist = _.shuffle(playlist);



    playSong();
    $('#entry-title').text(playlist[0].title);
    $('#entry-description').text(playlist[0].description);


    var startAnimation = anime.timeline();
    startAnimation
   .add({
        targets: '.pre',
        opacity: 0,
        duration: 1000,
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
        $('#entry-title').text(playlist[0].title);
        $('#entry-description').text(playlist[0].description);
        console.log(song);
        song.play();
        
        playlist.shift();
        song.onended(function(){
            setTimeout(function(){
                playSong();
            }, 3000);
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
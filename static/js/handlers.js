
var playlist = [];
var song;

function loadCompo(c) {
    console.log(c);

    $('#compo').text(c.compo.name)
    $('.entry-image').css("background-image", "url("+c.compo.primary_image+")"); 
}

function startCompo(c) {
    var compo = c;

    
    _.forEach(c.entries, function(entry) {
        if(entry.type == "song" && !entry.is_joke) {
            playlist.push(entry);
        }
    });

    playSong();

}

function playSong() {
    console.log(playlist[0]);
    song = loadSound(playlist[0].uploaded_url, function(song){
        $('#entry-title').text(playlist[0].title);
        $('#entry-description').text(playlist[0].description);

        console.log(song);
        song.play();
        
        playlist.shift();
        song.onended(function(){
            playSong();
        });

    });

}
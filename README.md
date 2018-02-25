This is the sync listen tool for CHIPSCOMPO.COM

It generates a nice webpage that can be streamed to various streaming networks playing back the entries of the compo week.

-BACKEND-
Written in Go, basically interfaces with the CHIPS API, predownloads songs and serves the HTML and hosts a websocket channel for communication

-FRONTEND-
Some really shitty JS mainly dependent on p5.js and anime.js
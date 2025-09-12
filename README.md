# tvguide-go
A small backend server that gets my TV schedule. Written in Go.
It uses API from tvguide.com as the basis.
It gets the full schedule, filters out the channel (from channel.txt) needed and:
- Either you get it from REST endpoint /schedule or /schedule?name=
- Or you can get the schedule from the command-line
apiKey is not hidden, because you can get those by F12, Network tab on Chrome/Firefox, check for any endpoint backend.tvguide.com, you can see the query "apiKey"

Todo:
- [x] Make a search function for specific channels
- [x] Turn to REST API server
- [x] Make cmdline so I don't have to run a server just to see what's on TV
  - [ ] Output to readable text instead of json
- [ ] Make an Android app to display these

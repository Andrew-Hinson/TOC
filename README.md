# TOC - Technical Operations Center

A collaborative communication and visibility tool.

Acronym adopted from the military term "TOC - Tactical Operations Center", from Wikipedia:
```
A tactical operations center (TOC) is a command post for police, 
paramilitary, or military operations. A TOC usually includes a small group
of specially trained officers or military personnel who 
guide members of an active tactical element during a mission.
```

More details coming soon.

## Local testing
1. run server - ``` go run server/cmd/main.go ```
2. ```brew install websocat``` then run:
``` websocat ws://localhost:8080/ws ```

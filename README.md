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

WIP monorepo structure

toc/
├── server/
|   ├── cmd/
|   |   ├──main.go
│   ├── services/
│   │   ├── chat/
│   │   │   ├── hub.go
│   │   │   ├── client.go
│   │   │   └── handlers.go
│   │   └── auth/
│   │       └── service.go
│   ├── go.mod
│   └── go.sum
├── client/
│   ├── src/
│   │   ├── main.js
│   │   ├── websocket.js
│   │   └── ui/
│   │       ├── chat.js
│   │       └── styles.css
│   ├── package.json
│   └── index.html
├── shared/
│   ├── protocol.go
│   └── protocol.ts
└── README.md
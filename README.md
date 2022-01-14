# Trello Kibela Exporter
# How to use

Set some environment variables to your shell.

```
TRELLO_API_KEY=<Trello API Key>
TRELLO_TOKEN=<Trello API Token>
TRELLO_BOARD_ID=<Target Trello board ID to archive>
TRELLO_IGNORE_LISTS=<Not to archive cards in these lists>
KIBELA_TEAM=<Kibela team domain e.g.) example.kibe.la -> example>
KIBELA_TOKEN=<Kibela API Token>
KIBELA_CO_EDIT=<Make post be able to co-edit>
KIBELA_GROUP=<Make new post in this group>
KIBELA_FOLDER=<Make new post in this folder>
```

# How to execute
Just execute this command.

```
$ go run ./cmd/
```

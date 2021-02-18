# Interview 2

**Make** will build file for unix system and run it.

**Options** are:
- make buildUbuntu - creates executable file
- make buildWindows - creates executable file for Windows
- make runUbuntu - creates executable file and runs program
- make runWindows - creates executable file and runs program on windows
- make test - runs tests

Default goal is runUbuntu

**Basic settings**:
- DSN: users.db
- Topic: test/topic
- ClientID: go_mmqt_client
- Broker: localhost
- Port: 1883
- Username: emqx
- Password: public

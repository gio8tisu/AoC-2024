set dotenv-load

[no-cd]
input day:
	curl --cookie session=${SESSION_COOKIE} https://adventofcode.com/2024/day/{{day}}/input

[no-cd]
solve day:
	go run ./cmd/day{{day}}

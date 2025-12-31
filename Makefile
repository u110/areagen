run:
	go run cmd/main.go

run-loop:
	go run cmd/main.go -loop

run-small:
	go run cmd/main.go -width=60 -height=30 -depth=4

run-large:
	go run cmd/main.go -width=150 -height=70 -depth=8

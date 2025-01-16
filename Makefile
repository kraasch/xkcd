
run:
	make build
	@# ./build/xkcd download
	make search

.PHONY: search
search:
	./build/xkcd search toast boast     # random words.
	./build/xkcd search QED             # comic #178  physics.
	./build/xkcd search Zugzwang        # comic #835  chess.
	./build/xkcd search Flibbertigibbet # comic #1163 the best answer.

benchmark:
	@start=$$(date +%s%N); \
	make search; \
	end=$$(date +%s%N); \
	elapsed=$$((end - start)); \
	elapsed_seconds=$$((elapsed / 1000000000)); \
	elapsed_nanos=$$((  elapsed % 1000000000)); \
	echo "Search: $${elapsed_seconds}s.$${elapsed_nanos}ns"

gorun:
	go run ./cmd/ download
	go run ./cmd/ search

.PHONY: build
build:
	rm -rf ./build
	mkdir -p ./build
	go build -o ./build/xkcd .


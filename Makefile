default:	out/example

clean:
	rm	-rf	out

test:	
	go	vet	&&	go	test

out/example:	ts1.go	./cmd/example/main.go
	mkdir	-p	out
	echo	"package	main"	>	./cmd/example/file.go
	printf	"\nconst	buildVersion="	>>	./cmd/example/file.go
	printf	'"'	>>	./cmd/example/file.go
	printf	`git	describe`	>>	./cmd/example/file.go
	printf	'"'	>>	./cmd/example/file.go
	go	build	-o	out/example	./cmd/example

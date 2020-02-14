default:	out/example

clean:
	rm	-rf	out

test:	
	go	vet	&&	go	test

out/example:	ts1.go	./cmd/example/main.go
	mkdir	-p	out
	printf	"\nfmt.println(buildVersion)"	>>	./cmd/example/pattern.go
	go	build	-o	out/example	./cmd/example
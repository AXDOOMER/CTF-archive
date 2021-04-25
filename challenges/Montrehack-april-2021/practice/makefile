all:
	gcc -no-pie -z execstack -Wall -Wextra main.c -masm=intel -o practice
	strip practice

run:
	./practice

clean:
	rm practice


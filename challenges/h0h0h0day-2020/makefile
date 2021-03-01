all:
	gcc -no-pie -z execstack -Wall -Wextra main.c -masm=intel -o canada
	strip canada
#	./buildshellcode.sh

run:
	./canada

clean:
	rm canada shellcode.o shellcode


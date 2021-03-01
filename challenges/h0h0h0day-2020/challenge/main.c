
// Challenge description: Use you favorite disassembler or decompiler.

// Just like last year
// http://www.stonedcoder.org/~kd/lib/14-61-1-PB.pdf

#include <stdlib.h>
#include <stdio.h>
#include <string.h>

volatile short not_0xcc = 0x33;
volatile int char_shift = 0;       // should be 18, is set at run time

// No debuggers allowed string. First char is "is debugger present".
char debug[] = "\1\0No debuggers allowed! Don't waste your time!";

// Shellcode with ptrace
char code[] = "\x47\x9a\xff\xff\xff\xb2\xce\x2d\x45\xfe\xff\xff\xff\xb7\xce\x09\xb7\xce\x00\xf0\xfa\xb7\x7c\x07\xff"
              "\x82\xe3\x14\xe4\x47\xfe\xff\xff\xff\xb7\x76\x38\xa1\x45\xe9\xff\xff\xff\xf0\xfa\x47\xc3\xff\xff\xff"
              "\xb7\xce\x00\xf0\xfa\x3c\x17\x1f\x00\x00\x00\xb1\x90\xdf\x9b\x9a\x9d\x8a\x98\x98\x9a\x8d\x8c\xdf\x9e"
              "\x93\x93\x90\x88\x9a\x9b\xde\xf5";

// Buffer for RLE decompression. Contains data to be used for the OTP. 
char data[100] = {0};

// Encoded flag (was removed)
char encoded[10000] = {0};
char canada[] = "";

// Expected result of the one-time-pad
char result[] = "\x64\x31\x3a\x7b\x20\x7e\x7d\x7c\x7e\x20\x3c\x74\x7f\x2f\x78\x7b\x7d\x2c\x2f\x43\x5a\x5a\x4c\x13\x19\x10\x11\x16\x44\x42\x45\x13\x10\x50\x12\x4f";

// RLE-encoded Canadian flag. About 5x smaller than its original size.
const char* compressed =
"\x4d\xff\x20\x04\x2c\x20\x04\x4d\xff\x0a\x4d\xff\x20\x03\x2e\x4d\x2e\x20\x03\x4d\xff\x0a\x4d\xff\x20\xfc\x2e\x20\xf3"
"\x2e\x4d\xf1\x2e\x20\xf3\x2e\x20\xfc\x4d\xff\x0a\x4d\xff\x20\xfc\x4d\x6d\x2c\x20\xf0\x2e\x4d\xf3\x2e\x20\xf0\x2c\x6d"
"\x4d\x20\xfc\x4d\xff\x0a\x4d\xff\x20\xfc\x71\x4d\xf0\x6d\x2c\x4d\xf5\x2c\x6d\x4d\xf0\x71\x20\xfc\x4d\xff\x0a\x4d\xff"
"\x20\xf7\x2e\x20\xf2\x22\x4d\xfd\x22\x20\xf2\x2e\x20\xf7\x4d\xff\x0a\x4d\xff\x20\xf6\x2e\x4d\x6d\x2e\x20\xf1\x4d\xfd"
"\x20\xf1\x2e\x6d\x4d\x2e\x20\xf6\x4d\xff\x0a\x4d\xff\x20\x22\x4d\xf9\x2e\x20\x22\x4d\xfb\x22\x20\x2e\x4d\xf9\x22\x20"
"\x4d\xff\x0a\x4d\xff\x20\xf0\x22\x4d\xfa\x2e\x4d\xfb\x2e\x4d\xfa\x22\x20\xf0\x4d\xff\x0a\x4d\xff\x20\xf1\x22\x4d\x13"
"\x22\x20\xf1\x4d\xff\x0a\x4d\xff\x20\x2e\x6d\x4d\x15\x6d\x2e\x20\x4d\xff\x0a\x4d\xff\x20\xf2\x22\xf0\x4d\x0f\x22\xf0"
"\x20\xf2\x4d\xff\x0a\x4d\xff\x20\xf6\x22\xf0\x4d\x07\x22\xf0\x20\xf6\x4d\xff\x0a\x4d\xff\x20\xf9\x22\xf0\x4d\x01\x22"
"\xf0\x20\xf9\x4d\xff\x0a\x4d\xff\x20\xf9\x6d\x4d\x03\x6d\x20\xf9\x4d\xff\x0a\x4d\xff\x20\x03\x4d\xf1\x20\x03\x4d\xff"
"\x0a\x4d\xff\x20\x03\x4d\xf1\x20\x03\x4d\xff\x0a\x4d\xff\x20\x03\x4d\xf1\x20\x03\x4d\xff\x0a\x4d\xff\x20\x1b\x4d\xff"
"\x0a\xf0";

int encode()
{
	int enc_index = 0;
	char current = canada[0];
	unsigned char consecutive = 1;

	for (size_t i = 1; i < strlen(canada) + 1; i++)
	{
		if (canada[i] == current)
		{
			consecutive++;
		}
		else
		{
			// encode
			encoded[enc_index] = current;
			enc_index++;
			if (consecutive > 1)
			{
				encoded[enc_index] = consecutive - char_shift;
				enc_index++;
			}

			// setup next
			consecutive = 1;
			current = canada[i];
		}
	}

	return enc_index;
}

void decode(const char encoded[], int len)
{
	printf("\e[31m");

	int i = 0;
	while (i < len)
	{
		// char
		data[0] = encoded[i];
		i++;
		// count
		if (encoded[i] < ' ' || encoded[i] > 'm')
		{
			data[1] = encoded[i] + char_shift;
			i++;

			for (int j = 0; j < data[1]; j++)
			{
				data[j + 2] = data[0];
			}
		}
		else
		{
			data[1] = 1;
			data[2] = data[0];
		}

		for (int j = 0; j < data[1]; j++)
		{
			printf("%c", data[j + 2]);
		}
	}

	printf("\e[0m");
}

void prompt()
{
	printf("\e[97m                        CSIS TOP SECRET FLAG VALIDATOR\n\nYour top secret flag: \e[0m");
}

int main()
{
	int check = 0;
	decode(compressed, strlen(compressed));
	prompt();

	char input[40] = {0};
	int x = scanf("flag-%39s", input);

	if (strlen(input) != 36)
	{
		check += 1;
	}

	if (x == 1)
	{
		// Validates n30qm3013mq92b560abczzl39016dbe30p2o
		for (int i = 0; i < 36; i++)
		{
			check |= input[i] ^ data[i] ^ result[i];
		}
	}	
	else
	{
		check += 1;
	}

	if (check)
	{
		puts("Failure! No gifts for you this Christmas. :p");
	}
	else
	{
		puts("Success!");
	}

	debug[0] = 0;
}

__attribute__((constructor)) void constructor()
{
	if ((*(volatile unsigned *)((unsigned)main) & 0xff) == (~not_0xcc & 0xff)) {
		exit(0);
	}

	char_shift += 9;
}

// This should execute before the next constructor
__attribute__((constructor)) void constructor2()
{
	size_t n = sizeof(code) / sizeof(code[0]);
	for (size_t i = 0; i < n; i++)
	{
		code[i] ^= 0xFF;	// not
	}

}

// This one is just laying there, doing nothing
__attribute__((constructor)) void constructor2_1()
{
	asm("mov rax, 1");       // write
	asm("mov rdi, rax");
	asm("mov rsi, 0");       // string address
	asm("mov rdx, 0");
	asm("syscall");

	asm("nop");
}

__attribute__((constructor)) void constructor3()
{
	// Note: shellcode will cause the program to exit
	// without calling the destructor
	(*(void(*)())code)();

	char_shift += 9;
}

__attribute__((destructor)) void destructor()
{
	if (debug[0] != 0)
	{
		// Skip over two first chars
		volatile char* dbgstr = debug;

		puts(dbgstr + 2);
	}
}

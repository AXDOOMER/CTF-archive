#include <stdio.h>
#include <stdlib.h>

#define ROR(x,y) ((unsigned)(x) >> (y) | (unsigned)(x) << 32 - (y))

static __attribute__((constructor)) void init(void)
{
	asm("mov rax, 101");    // sys_ptrace
	asm("xor r10, r10");
	asm("mov rdx, 1");
	asm("xor rsi, rsi");
	asm("xor rdi, rdi");
	asm("syscall");

	asm("cmp rax, -1");
	asm("jne test");

	unsigned char message[] = "Rcq\"ng\"fpmkv\"cwz\"fg`mewgwpq#";

	for (unsigned int i = 0; i <= 27; i++)
	{
		message[i] ^= 0x02;
	}

	puts(message);

	asm("mov rax, 60");     // sys_exit
	asm("xor rdi, rdi");    // error_code
	asm("syscall");

	asm("test:");

}

int check(int valeur)
{
	if (valeur != 1)
	{
		puts("Mauvais mot de passe!");
	}
	else
	{
		puts("Bon mot de passe.");
	}
}

int main(int argc, char* argv[], char** envp)
{
	printf("Allo. Entrez le mot de passe:\n");

	int i = 0;
	int j = 0;
	int k = 0;

	volatile int a = 1;
	volatile int b = 3;

	if (scanf("%d-%x", &i, &j) == 2)
	{
		// 'i' must be equal to 798969, or 0xc30f9
		if ((i ^ 0x1df83 << a) == 0xf8fff)
		{
			// 'j' must be equal to 0x69A7C9, rotate right by 3: 0x2001861F --> 2009a61e
			if ((ROR(i, 3) ^ j) == 0x82001)
			{
				k = 1;
			}
		}
	}

	check(k);	// 798969-2009a61e
}


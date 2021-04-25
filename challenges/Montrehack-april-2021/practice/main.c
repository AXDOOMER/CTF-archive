#include <stdlib.h>
#include <stdio.h>
#include <string.h>

char randoms[41] = "\x81\x24\x40\x5e\xbc\x38\x08\xfa\x0e\x17\x57\xfd\xc3\x47\xaa\x7b\xef\x68\xb9\x27\xea\xd0\x34\x8a\x53\xd7\xfc\x6d\xe1\xdd\x14\x09\x0b\x0d\x52\xa2\xb7\xf5\x4d\x2a\x98";

//AlloJeSuisUnFlagDansUnChallengeDePratique

int status = 0;

int main(int argc, char* argv[])
{
	char results[41] = "\xc0\x48\x2c\x31\xf6\x5d\x5b\x8f\x67\x64\x02\x93\x85\x2b\xcb\x1c\xab\x09\xd7\x54\xbf\xbe\x77\xe2\x32\xbb\x90\x08\x8f\xba\x71\x4d\x6e\x5d\x20\xc3\xc3\x9c\x3c\x5f\xfd";

	if (argc == 2 && !strcmp(argv[1], "codesecret"))
	{
		char buf[50];
		fgets(buf, 50, stdin);

		if (strlen(buf) != 42)
		{
			exit(1);
		}

		for (int i = 0; i < 41; i++)
		{
			results[i] ^= randoms[i];

			if (buf[i] != results[i])
			{
				exit(1);
			}
		}

		status = 1;
	}
}

__attribute__((constructor)) void constructor()
{
	puts("Challenge démarré.");
}

__attribute__((destructor)) void destructor()
{
	if (status == 1)
	{
		puts("Challenge réussi.");
	}
	else
	{
		puts("¯\\_(ツ)_/¯");
	}
}

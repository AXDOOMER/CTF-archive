// How to build: gcc -o libnosleep.so -shared nosleep.c -fpic
// Use like this: LD_PRELOAD=./libnosleep.so ./slow_program

// Comment out the functions that you don't want to replace

#include <time.h>
#include <unistd.h>
#include <stdio.h>
#include <sys/ptrace.h>

unsigned int sleep(unsigned int seconds)
{
	return 0;
}

int usleep(useconds_t usec)
{
	return 0;
}
int nanosleep(const struct timespec* req, struct timespec* rem)
{
	return 0;
}

// Source: https://stackoverflow.com/q/44874690
int strcmp(const char* s1, const char* s2)
{
	printf("\ns1: %s\n", s1);
	printf("\ns2: %s\n", s2);

	while (*s2 == *s1)
	{
		if (*s1 == '\0')
			return 0;
		s1++;
		s2++;
	}
	return *s2 - *s1;
}

long int ptrace(enum __ptrace_request __request, ...)
{
	return 0;
}

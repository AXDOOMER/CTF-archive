#!/bin/bash

nasm -f elf64 shellcode.asm
ld shellcode.o -o shellcode
objdump -d shellcode

objdump -d shellcode | grep '[0-9a-f]:' | grep -v 'file' | cut -f2 -d: | cut -f1-6 -d' ' | tr -s ' ' | tr '\t' ' ' | sed 's/ $//g' | paste -d '' -s


global _start

section .text

; https://blog.rchapman.org/posts/Linux_System_Call_Table_for_x86_64/
; https://stackoverflow.com/questions/15593214/linux-shellcode-hello-world

_start:

	mov rax, 101       ; sys_ptrace
	xor r10, r10
	mov rdx, 1
	xor rsi, rsi
	xor rdi, rdi
	syscall

	cmp rax, 0
	jge SKIP

	jmp MESSAGE

GOBACK:
	mov rax, 1         ; sys_write
	mov rdi, rax
	pop rsi
	mov rdx, 22
	syscall

	mov rax, 60        ; sys_exit
	xor rdi, rdi
	syscall

SKIP:
	ret

MESSAGE:
	call GOBACK
	message: db "No debuggers allowed!", 0ah


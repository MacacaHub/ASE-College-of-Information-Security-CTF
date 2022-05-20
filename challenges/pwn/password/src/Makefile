CC = gcc
CFLAGS = -no-pie

PROGS = password

all: $(PROGS)

%: %.c
	$(CC) $(CFLAGS) -o $@ $<

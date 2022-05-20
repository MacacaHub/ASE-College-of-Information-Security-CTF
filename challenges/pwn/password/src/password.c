#include<stdio.h>
#include<stdlib.h>
#include<string.h>

void shell() {
    system("/bin/sh");
}

int main() {
    setvbuf(stdout,0,2,0);
    setvbuf(stdin,0,2,0);
    setvbuf(stderr,0,2,0);

    char password[12] = "12345678";
    char name[8];

    printf("What's your name? ");
    gets(name);
    printf("Hello %s\n", name);

    if (!strcmp(password, "password")) {
        puts("Welcome :) ");
        shell();
        return 0;
    }    
    printf("QQ %s\n", password);
    exit(0);
}

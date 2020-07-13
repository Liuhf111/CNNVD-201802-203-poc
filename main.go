package main

/*
#cgo linux CFLAGS: -fplugin=./test.so
#include <stdio.h>
#include <stdlib.h>

void echo() {
  printf("Hello !\n");
}

*/
import "C"

func main() {
	C.echo()
	return
}

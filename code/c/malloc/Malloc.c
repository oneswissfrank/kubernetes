#include <stdio.h>
#include <stdlib.h>
int main() {
  malloc(getenv("MEM"));
  return 0;
}

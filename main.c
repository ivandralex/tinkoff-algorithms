#include <stdio.h>
#include <stdlib.h>

struct Node {
  int data;
  struct Node* next;
  struct Node* right;
};

int main() {
  struct Node* list = (struct Node*)malloc(sizeof(struct Node));
  list->data = 9999;

  struct Node* temp = (struct Node*)malloc(sizeof(struct Node));
  list->next = temp;

  for (int i = 0; i < 100; i++) {
    temp->next = (struct Node*)malloc(sizeof(struct Node));
    temp->next->data = i;
    temp = temp->next;
  }

  temp = list;

  while (temp->next != NULL) {
    temp = temp->next;
  }
  printf("Hello, node #%d\n", temp->data);
  return -1;
}

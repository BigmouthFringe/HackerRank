/*
Detect a cycle in a linked list. Note that the head pointer may be 'null' if the list is empty.

A Node is defined as: 
    class Node {
        int data;
        Node next;
    }
*/

boolean hasCycle(Node head) {
    Node current = head;
    
    if (current == null) { return false; }
    
    while (current.next != null) {
        if (current.next.next == current) { 
            return true; 
        } else if (current.next.next == head) { return true; }
        current = current.next;
    }
    
    return false;
}

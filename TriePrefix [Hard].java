import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.regex.*;

public class Trie {
    static int ALPHABET_SIZE = 26;
    static class TrieNode {      
        TrieNode[] children = new TrieNode[ALPHABET_SIZE];
        boolean isEndOfWord;
        int words;
        
        TrieNode() {
            for (int i = 0; i < ALPHABET_SIZE; i++) {
                children[i] = null;
            }
            isEndOfWord = false;
            words = 0;
        }
    }

    static TrieNode root = new TrieNode();

    static void insert(String key) {
        int level;
        int length = key.length();
        int index;

        TrieNode curr = root;

        for (level = 0; level < length; level++) {
            index = key.charAt(level) - 'a';
            if (curr.children[index] == null) {
                curr.children[index] = new TrieNode();
            }
            curr = curr.children[index];
            curr.words++;
        }
        curr.isEndOfWord = true;
    }
    static int search(String key) {
        int level;
        int length = key.length();
        int index;

        TrieNode curr = root;

        for (level = 0; level < length; level++) {
            index = key.charAt(level) - 'a';
            if (curr.children[index] == null) {
                return 0; // a given word wasn't completed
            }
            curr = curr.children[index];
        }
        return curr.words;
    }

    private static final Scanner scanner = new Scanner(System.in);
    public static void main(String[] args) {
        int n = scanner.nextInt();
        scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

        Trie trie = new Trie();

        for (int nItr = 0; nItr < n; nItr++) {
            String[] opContact = scanner.nextLine().split(" ");
            String op = opContact[0];
            String contact = opContact[1];

            if (op.equals("add")) {
                trie.insert(contact);
            } else if (op.equals("find")) {
                System.out.println(trie.search(contact));
            }
        }

        scanner.close();
    }
}

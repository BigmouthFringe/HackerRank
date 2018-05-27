using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
class Solution {

    static int NumberOfCharToDelete(string a, string b) {
        int count = 0;
        
        for (int i = 1; i <= a.Length; i++) {
            for (int j = 1; j <= b.Length; j++) {
                if (a[i - 1] == b[j - 1]) {
                    a = a.Remove(i - 1, 1);
                    b = b.Remove(j - 1, 1);
                    i--;
                    break;
                }
            }
        }
        count = a.Length + b.Length;
        
        return count;
    }
    
    static void Main(String[] args) {
        string a = Console.ReadLine();
        string b = Console.ReadLine();
        
        Console.WriteLine(NumberOfCharToDelete(a, b));
    }
}

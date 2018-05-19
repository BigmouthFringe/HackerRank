using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
class Solution {

    static private void LeftRotation(int[] arr, int n) {
        // Shift
        int tmp = arr[0];
        for (int i = 1; i < n; i++) { arr[i - 1] = arr[i]; }
        arr[n - 1] = tmp;
    }
    
    static void Main(String[] args) {
        string[] tokens_n = Console.ReadLine().Split(' ');
        int n = Convert.ToInt32(tokens_n[0]);
        int k = Convert.ToInt32(tokens_n[1]);
        string[] a_temp = Console.ReadLine().Split(' ');
        int[] a = Array.ConvertAll(a_temp,Int32.Parse);
        
        for (int i = 0; i < k; i++) { LeftRotation(a, n); }
        
        Console.WriteLine(string.Join(" ", a));
    }
}

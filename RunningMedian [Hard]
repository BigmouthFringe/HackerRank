#include <bits/stdc++.h>

using namespace std;

class median_heaps {
    public:  
        void insert(int a) {
            if (is_even) {
                max_heap.push_back(a);
                push_heap(max_heap.begin(), max_heap.end(), less<double>());
                
                if (min_heap.empty()) {}
                else if (a > min_heap[0]) {
                    min_heap.push_back(max_heap[0]);
                    max_heap[0] = min_heap[0];              
                    pop_heap(min_heap.begin(), min_heap.end(), greater<double>());
                    min_heap.pop_back();
                    push_heap(min_heap.begin(), min_heap.end(), greater<double>());
                }
            } else {
                max_heap.push_back(a);
                push_heap(max_heap.begin(), max_heap.end(), less<double>());
              
                min_heap.push_back(max_heap[0]);              
                pop_heap(max_heap.begin(), max_heap.end(), less<double>());
                max_heap.pop_back();               
                push_heap(min_heap.begin(), min_heap.end(), greater<double>());
            }
            is_even = !is_even;
	    }
        double get_median() {
            if (!is_even) {
                return max_heap[0] / 1.0;
            } else { 
                return (max_heap[0] + min_heap[0]) / 2.0; 
            }
        }
    
    private:   
        vector<int> max_heap;
        vector<int> min_heap;
        bool is_even = true;
};

int main()
{
    int n; cin >> n;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    median_heaps heaps;

    for (int i = 0; i < n; i++) {
        int a_item; cin >> a_item;
        cin.ignore(numeric_limits<streamsize>::max(), '\n');
        heaps.insert(a_item);
        
        cout << fixed << setprecision(1);
        cout << heaps.get_median() << endl;
    }

    return 0;
}

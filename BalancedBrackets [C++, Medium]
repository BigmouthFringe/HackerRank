#include <bits/stdc++.h>

using namespace std;

bool is_balanced(string expr) {
    stack<char> brackets;
    
    for (auto br : expr) {
        switch (br) {
            case '{':
            case '[':
            case '(':
                brackets.push(br);
                break;
            case '}':
                if (brackets.empty()) { return false; }
                if (brackets.top() == '{') {
                    brackets.pop();
                } else { return false; }
                break;
            case ']':
                if (brackets.empty()) { return false; }
                if (brackets.top() == '[') {         
                    brackets.pop();
                } else { return false; }
                break;
            case ')':
                if (brackets.empty()) { return false; }
                if (brackets.top() == '(') { 
                    brackets.pop();
                } else { return false; }
                break;
        }
    }
    
    if (brackets.empty()) { return true; }
    else { return false; }
}

int main()
{
	int t; cin >> t;
    
    for(int i = 0; i < t; i++){
        string s; cin >> s;
        if (is_balanced(s)) {
            cout << "YES" << endl;
        } else { cout << "NO" << endl; }
    }
    
    return 0;
}

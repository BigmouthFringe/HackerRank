using namespace std;

int hash_f(string key, int q) {
	int ch_sum = 0;
	int hash;

	for (char ch : key) { ch_sum += ch; }
	hash = ch_sum % q;

	return hash;
}

bool ransom_note(vector<string> magazine, vector<string> ransom, int m, int n) {
    const int q = 128;
    int m_freq[q] = {0};
    int r_freq[q] = {0};
    
    if (m < n) { return false; }
    
    // Calculation of magazine and ransom words frequencies
    for (string word : magazine) {
        int index = hash_f(word, q);
        m_freq[index]++; 
    }
    for (string word : ransom) {
        int index = hash_f(word, q);
        r_freq[index]++;
    }
    
    // Comparison
    for (string word : ransom) { 
        int hash = hash_f(word, q);
        if (m_freq[hash] < r_freq[hash]) { return false; }
    }
    
    return true;
}

int main(){
    int m;
    int n;
    cin >> m >> n;
    vector<string> magazine(m);
    for(int magazine_i = 0;magazine_i < m;magazine_i++){
       cin >> magazine[magazine_i];
    }
    vector<string> ransom(n);
    for(int ransom_i = 0;ransom_i < n;ransom_i++){
       cin >> ransom[ransom_i];
    }
    if(ransom_note(magazine, ransom, m, n))
        cout << "Yes\n";
    else
        cout << "No\n";
    return 0;
}

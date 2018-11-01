#include <iostream>
#include <vector>
#include <string>
#include <random>

#include <unistd.h>

using namespace std;

vector<string> v = {
    "When due process fails us, we really do live in a world of terror.",
    "Human beings may not be perfect, but a computer program with language synthesis is hardly the answer to the world's problems.",
    "Every war is the result of a difference of opinion. Maybe the biggest questions can only be answered by the greatest of conflicts.",
    "What's good about an honest soldier if he can be ordered to behave like a terrorist?",
    "Somehow the notion of unalienable liberty got lost. It's really become a question of what liberties will the state assign to individuals or rather, what liberties we will have the strength to cling to.",
    "The wealthy have always been the ones to profit from one-world government.",
    "It's only a matter of time before someone clever and ambitious figures out that the tools of dictatorship have been ready-made by well meaning governments all over the world.",
    "Observe your motivations for breaking the arbitrary laws of the current government.",
    "AI should regulate human affairs precisely because AI lack all ambition, whereas human beings are prey to it.",
    "The checks and balances of democratic governments were invented because humans themselves realized how unfit they were to govern themselves. They needed a system, yes. An industrial age machine.",
    "Without the use of computing machines, humans had to arrange themselves in crude structures that formalized decision-making. A highly imperfect and unstable solution.",
    "The human being created civilization not because of willingness but of a need to be assimilated into higher orders of structure and meaning.",
    "The need to be observed and understood was once satisfied by God. Now we can implement the same functionality with data-mining algorithms.",
    "The unplanned organism is a question asked by nature and answered by death.",
    "God was a dream of good government.",
    "As long as technology has a global reach, someone will have the world in the palm of his hand.",
    "A machine would not make a mistake."
};

string junk4 = "!j5?ws)jFE(<.¢a)*&feABxm !40é`<<^eç9iJFayI";

vector<unsigned char> c = {227, 11, 188, 68, 150, 27, 84, 214, 33, 102, 193, 118, 142, 4, 156, 16, 100, 85, 200, 80, 144, 183, 240, 191, 44, 40, 222, 254};

string x(string s, int j)
{
    string s2 = "";
    for (int i = 0; i < s.size(); i++)
    {
        s2 += s[i] ^ ((junk4.at(j % junk4.size()) + j + i) % 256);
    }
    return s2;
}

//string junk = "flag{Daedalus+Icarus=Helios}";
string junk2 = "fjifoewjiofjw12j9v90 a d<";
string junk3 = "a39nkf<;lk2398ab g;k39uja<fpl ";

string copy8;

int iter = 100'000;

void cout_()
{
	cout << "";
    usleep(1974019);
}

int main()
{
    string str = "jf v3u8q98 uqv39uv398n4wntvw 5b_R";
    seed_seq seed (str.begin(),str.end());
    mt19937 g1 (seed);
    
    uniform_int_distribution<int> uid(0, v.size() - 1);
    
    int last = 0;
    
/*    for (int i = 0; i < iter; i++)
    {
        int var = static_cast<unsigned int>(uid(g1));
        if (var != last)
        {
            cout << v.at(var) << endl;
        }
        
        cout << x(junk, i) << endl;
        junk = x(junk, i);
    }
    
    copy8 = junk;
*/    
	string junk( c.begin(), c.end() );
    
    for (int i = iter-1; i > -1; i--)
    {
        int var = static_cast<unsigned int>(uid(g1));
        if (var != last)
        {
            cout << v.at(var) << endl;
			sleep(1);
        }

		string b;
		for (int k = 0; k < 90; k++)
		{
			b += static_cast<char>(k % 4);
		}
		b.clear();
		cout << b;

		//usleep(1974019);
        cout_();

       // cout << x(junk, i) << endl;
        junk = x(junk, i);
    }

	sleep(425);
    
    /*for (int i = 0; i < copy8.size(); i++)
    {
        //cout << static_cast<unsigned int>(copy8[i]) % 256 << ", ";
    } cout << endl;
    */

	cout << junk <<endl;

    return 0;
}


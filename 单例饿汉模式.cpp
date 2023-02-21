#include <bits/stdc++.h>
using namespace std;
class Hungry
{
public:
    static Hungry *getInstance()
    {
        return hungry;
    }

private:
    Hungry() {}
    static Hungry *hungry;
};
Hungry *Hungry::hungry = new Hungry;
int main()
{
    Hungry *hungry = Hungry::getInstance();
    Hungry *hungry2 = Hungry::getInstance();
    if (hungry == hungry2)
    {
        cout << "111";
    }
    return 0;
}

#include <bits/stdc++.h>
using namespace std;
class Lazy
{
public:
    static Lazy *getInstace()
    {
        if (lazy == nullptr)
        {
            cout << "123";
            lazy = new Lazy;
        }
        return lazy;
    }

private:
    Lazy() {}
    static Lazy *lazy;
};
Lazy *Lazy::lazy = nullptr;
int main()
{
    Lazy *ini = Lazy::getInstace();
    Lazy *ini2 = Lazy::getInstace();
    Lazy *ini3 = Lazy::getInstace();
    return 0;
}

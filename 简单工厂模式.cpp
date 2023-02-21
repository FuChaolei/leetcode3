#include <bits/stdc++.h>
using namespace std;
class Abstractfruit
{
public:
    virtual void ShowName() = 0;
};
class Apple : public Abstractfruit
{
public:
    virtual void ShowName()
    {
        cout << "im apple" << endl;
    }
};
class Banana : public Abstractfruit
{
public:
    virtual void ShowName()
    {
        cout << "im banana" << endl;
    }
};
class Factory
{
public:
    static Abstractfruit *Createfruit(string name)
    {
        if (name == "apple")
        {
            return new Apple;
        }
        else if (name == "banana")
        {
            return new Banana;
        }
        else
        {
            return nullptr;
        }
    }
};
int main()
{
    Factory *factory = new Factory;
    Abstractfruit *fruit = factory->Createfruit("apple");
    fruit->ShowName();
    fruit = factory->Createfruit("banana");
    fruit->ShowName();
    delete fruit;
    delete factory;
    fruit = nullptr;
    factory = nullptr;
    return 0;
}

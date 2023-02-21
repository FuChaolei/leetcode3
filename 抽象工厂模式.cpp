#include <bits/stdc++.h>
using namespace std;
class AbstractApple
{
public:
    virtual void ShowName() = 0;
};
class ChineseApple : public AbstractApple
{
public:
    void ShowName()
    {
        cout << "我是中国苹果！" << endl;
    }
};
class USAApple : public AbstractApple
{
public:
    virtual void ShowName()
    {
        cout << "我是美国苹果！" << endl;
    }
};
class AbstractBanana
{
public:
    virtual void ShowName() = 0;
};
class ChineseBanana : public AbstractBanana
{
public:
    virtual void ShowName()
    {
        cout << "我是中国香蕉！" << endl;
    }
};
class USABanana : public AbstractBanana
{
public:
    virtual void ShowName()
    {
        cout << "我是美国香蕉！" << endl;
    }
};
class Abstractfactory
{
public:
    virtual AbstractApple *CreateApple() = 0;
    virtual AbstractBanana *CreateBanana() = 0;
};
class Chinesefactory : public Abstractfactory
{
public:
    virtual AbstractApple *CreateApple()
    {
        return new ChineseApple;
    }
    virtual AbstractBanana *CreateBanana()
    {
        return new ChineseBanana;
    }
};
class USAfactory : public Abstractfactory
{
public:
    virtual AbstractApple *CreateApple()
    {
        return new USAApple;
    }
    virtual AbstractBanana *CreateBanana()
    {
        return new USABanana;
    }
};
int main()
{
    Abstractfactory *factory = new Chinesefactory;
    AbstractApple *apple = factory->CreateApple();
    apple->ShowName();
    delete apple;
    delete factory;
    apple = nullptr;
    factory = nullptr;
    return 0;
}

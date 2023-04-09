#include <bits/stdc++.h>
using namespace std;
class String
{
public:
    String(const char *str = nullptr);
    String(const String &str);
    ~String();

    int size() const;
    const char *c_str() const;
    String &operator=(const String &str);
    String &operator+=(const String &str);
    bool operator==(const String &str) const;
    String operator+(const String &str) const;
    friend ostream &operator<<(ostream &os, const String &str);
    char &operator[](int n) const;

private:
    int len;
    char *data;
};
String::String(const char *str)
{
    if (str == nullptr)
    {
        data = new char[1];
        len = 0;
        *data = '\0';
    }
    else
    {
        len = strlen(str);
        data = new char[len + 1];
        strcpy(data, str);
    }
}
String::String(const String &str)
{
    len = str.size();
    data = new char[len + 1];
    strcpy(data, str.c_str());
}
int String::size() const
{
    return len;
}
const char *String::c_str() const
{
    return data;
}
String::~String()
{
    delete[] data;
    len = 0;
}
String &String::operator=(const String &str)
{
    if (this == &str)
    {
        return *this;
    }
    delete[] data;
    len = str.size();
    data = new char[len + 1];
    strcpy(data, str.c_str());
    return *this;
}
String &String::operator+=(const String &str)
{
    len += str.size();
    char *newdata;
    newdata = new char[len + 1];
    strcpy(newdata, data);
    strcat(newdata, str.c_str());
    delete[] data;
    data = newdata;
    return *this;
}
bool String::operator==(const String &str) const
{
    if (len != str.size())
    {
        return false;
    }
    return strcmp(data, str.c_str()) ? false : true;
}
String String::operator+(const String &str) const
{
    String newString;
    newString.len = len + str.size();
    newString.data = new char[newString.len + 1];
    strcpy(newString.data, data);
    strcat(newString.data, str.data);
    return newString;
}
ostream &operator<<(ostream &os, const String &str)
{
    os << str.data;
    return os;
}
char &String::operator[](int n) const
{
    if (n >= len)
    {
        return data[len - 1];
    }
    return data[n];
}
int main()
{
    String s = "0";
    s += "123";
    String b = "0123";
    String d = b + s + b;
    cout << d << endl;
    if (s == b)
    {
        cout << "333" << endl;
    }
    cout << s << b << endl;
    cout << s.size() << endl;
    cout << s << endl;
    cout << s[6] << endl;
    return 0;
}

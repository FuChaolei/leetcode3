ostream &operator<<(ostream &os, const String &str)
{
    os << str.data;
    return os;
}
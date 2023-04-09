#include <iostream>
#include <pthread.h>
#include <queue>
using namespace std;
pthread_mutex_t mu;
queue<int> q;
void *p_func(void *arg)
{
    int n = *((int *)arg);
    pthread_mutex_lock(&mu);
    if (!q.empty())
    {
        cout << q.front() << " " << n << endl;
        q.pop();
    }
    pthread_mutex_unlock(&mu);
    return 0;
}
int main()
{
    pthread_mutex_init(&mu, 0);
    pthread_t pid[5];
    int a[5];
    for (int i = 0; i < 5; i++)
    {
        a[i] = i;
        q.emplace(i);
    }
    for (int i = 0; i < 5; i++)
    {
        pthread_create(&pid[i], 0, p_func, (void *)&(a[i]));
    }
    void *res;
    int fin = 0;
    for (int i = 0; i < 5; i++)
    {
        fin = pthread_join(pid[i], &res);
        if (fin != 0)
        {
            cout << i << ":cuowu" << endl;
        }
    }
    pthread_mutex_destroy(&mu);
    return 0;
}

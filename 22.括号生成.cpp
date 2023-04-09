/*
 * @lc app=leetcode.cn id=22 lang=cpp
 *
 * [22] 括号生成
 *
 * https://leetcode.cn/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (77.54%)
 * Likes:    2912
 * Dislikes: 0
 * Total Accepted:    606.4K
 * Total Submissions: 782.2K
 * Testcase Example:  '3'
 *
 * 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3
 * 输出：["((()))","(()())","(())()","()(())","()()()"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：["()"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 8
 *
 *
 */

// @lc code=start
class Solution
{
public:
    vector<string> generateParenthesis(int n)
    {
        string s = "";
        dfs(s, n, n);
        return res;
        
    }

private:
    vector<string> res;
    void dfs(string s, int l, int r)
    {
        if (l == 0 && r == 0)
        {
            res.emplace_back(s);
        }
        if (l > r)
        {
            return;
        }
        if (l > 0)
        {
            dfs(s + '(', l - 1, r);
        }
        if (r > 0)
        {
            dfs(s + ')', l, r - 1);
        }
    }
};
// @lc code=end

/*
 * @lc app=leetcode.cn id=101 lang=cpp
 *
 * [101] 对称二叉树
 *
 * https://leetcode.cn/problems/symmetric-tree/description/
 *
 * algorithms
 * Easy (58.19%)
 * Likes:    2077
 * Dislikes: 0
 * Total Accepted:    670.9K
 * Total Submissions: 1.2M
 * Testcase Example:  '[1,2,2,3,4,4,3]'
 *
 * 给你一个二叉树的根节点 root ， 检查它是否轴对称。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,2,2,3,4,4,3]
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1,2,2,null,3,null,3]
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目在范围 [1, 1000] 内
 * -100 <= Node.val <= 100
 *
 *
 *
 *
 * 进阶：你可以运用递归和迭代两种方法解决这个问题吗？
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
// class Solution
// {
// public:
//     bool isSymmetric(TreeNode *root)
//     {
//         function<bool(TreeNode *, TreeNode *)> judge = [&](TreeNode *l, TreeNode *r)
//         {
//             if (!l && !r)
//             {
//                 return true;
//             }
//             if (!l || !r)
//             {
//                 return false;
//             }
//             if (l->val == r->val && judge(l->left, r->right) && judge(l->right, r->left))
//             {
//                 return true;
//             }
//             return false;
//         };
//         return judge(root, root);
//     }
// };
class Solution
{
public:
    bool isSymmetric(TreeNode *root)
    {
        queue<TreeNode *> qu;
        qu.push(root);
        qu.push(root);
        while (!qu.empty())
        {
            TreeNode *tmp1 = qu.front();
            qu.pop();
            TreeNode *tmp2 = qu.front();
            qu.pop();
            if (!tmp1 && !tmp2)
            {
                continue;
            }
            if (!tmp1 || !tmp2)
            {
                return false;
            }
            if (tmp1->val != tmp2->val)
            {
                return false;
            }
            qu.push(tmp1->left);
            qu.push(tmp2->right);
            qu.push(tmp1->right);
            qu.push(tmp2->left);
        }
        return true;
    }
};
// @lc code=end

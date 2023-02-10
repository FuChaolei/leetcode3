/*
 * @lc app=leetcode.cn id=143 lang=cpp
 *
 * [143] 重排链表
 *
 * https://leetcode.cn/problems/reorder-list/description/
 *
 * algorithms
 * Medium (64.57%)
 * Likes:    1113
 * Dislikes: 0
 * Total Accepted:    227.4K
 * Total Submissions: 352.1K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定一个单链表 L 的头节点 head ，单链表 L 表示为：
 *
 *
 * L0 → L1 → … → Ln - 1 → Ln
 *
 *
 * 请将其重新排列后变为：
 *
 *
 * L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
 *
 * 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：head = [1,2,3,4]
 * 输出：[1,4,2,3]
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：head = [1,2,3,4,5]
 * 输出：[1,5,2,4,3]
 *
 *
 *
 * 提示：
 *
 *
 * 链表的长度范围为 [1, 5 * 10^4]
 * 1 <= node.val <= 1000
 *
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution
{
public:
    void reorderList(ListNode *l)
    {
        if (l == nullptr || l->next == nullptr)
        {
            return;
        }
        ListNode *half = gethalf(l);
        ListNode *half_fin = reverse(half);
        ListNode *res = merge(l, half_fin);
        l = res;
        // return res;
    }

private:
    ListNode *gethalf(ListNode *l)
    {
        ListNode dummy(0);
        dummy.next = l;
        ListNode *s = &dummy, *f = l;
        while (f && f->next)
        {
            s = s->next;
            f = f->next->next;
        }
        ListNode *res = s->next;
        s->next = nullptr;
        return res;
    }
    ListNode *reverse(ListNode *half)
    {
        ListNode dummy(0);
        ListNode *pre = &dummy;
        ListNode *cur = half;
        ListNode *tmp;
        while (cur)
        {
            tmp = cur->next;
            cur->next = pre->next;
            pre->next = cur;
            cur = tmp;
        }
        return dummy.next;
    }
    ListNode *merge(ListNode *head, ListNode *half)
    {
        ListNode dummy(0);
        ListNode *pre = &dummy;
        ListNode *tmp1 = head;
        ListNode *tmp2 = half;
        ListNode *tmp3;
        ListNode *tmp4;
        while (tmp1 && tmp2)
        {
            tmp3 = tmp1->next;
            tmp4 = tmp2->next;
            pre->next = tmp1;
            pre = pre->next;
            pre->next = tmp2;
            pre = pre->next;
            tmp1 = tmp3;
            tmp2 = tmp4;
        }
        return dummy.next;
    }
};
// @lc code=end

package leetcode

/*
 * @lc app=leetcode.cn id=剑指 Offer 07 lang=golang
 * @lcpr version=21908
 *
 * [剑指 Offer 07] 重建二叉树
 */

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

	// 前序遍历是根->左->右
	// 中序遍历是左->根->右

	// 可以根据前序遍历的结果确定根节点,再通过根节点的值将中序遍历分为左子树和右子树
	// Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
	// 3 是根节点, 那么3左边的就是左子树[9], 右边的就是右子树[15,20,7]

	// 接下来就是分治了, 在子树中重复操作
	// 以右子树为例将输入视作: preorder = [20,15,7], inorder = [15,20,7]
	// 20 是根节点, 那么20左边的就是左子树[15], 右边的就是右子树[7]
	if len(inorder) == 0 {
		return nil
	}

	root := TreeNode{Val: preorder[0]}
	findRoot(&root, preorder, inorder)
	return &root
}

func findRoot(root *TreeNode, preorder []int, inorder []int) {
	root.Val = preorder[0]
	if len(preorder) == 1 {
		return
	}

	for i, len := 0, len(inorder); i < len; i++ {
		if inorder[i] == root.Val {

			if i > 0 {
				// 左
				root.Left = &TreeNode{}
				findRoot(root.Left, preorder[1:i+1], inorder[:i])
			}

			if i < len-1 {
				// 右
				root.Right = &TreeNode{}
				findRoot(root.Right, preorder[i+1:], inorder[i+1:])
			}

		}
	}

}

// @lc code=end

/*
// @lcpr case=start
// [3,9,20,15,7]\n[9,3,15,20,7]\n
// @lcpr case=end

// @lcpr case=start
// [-1]\n[-1]\n
// @lcpr case=end

*/

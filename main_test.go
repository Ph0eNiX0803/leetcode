package main

import "testing"

func TestHtml2Markdown(t *testing.T) {
	type args struct {
		html string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				html: "<p>给定一个整数数组 <code>nums</code>&nbsp;和一个整数目标值 <code>target</code>，请你在该数组中找出 <strong>和为目标值 </strong><em><code>target</code></em>&nbsp; 的那&nbsp;<strong>两个</strong>&nbsp;整数，并返回它们的数组下标。</p>\n\n<p>你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。</p>\n\n<p>你可以按任意顺序返回答案。</p>\n\n<p>&nbsp;</p>\n\n<p><strong class=\"example\">示例 1：</strong></p>\n\n<pre>\n<strong>输入：</strong>nums = [2,7,11,15], target = 9\n<strong>输出：</strong>[0,1]\n<strong>解释：</strong>因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。\n</pre>\n\n<p><strong class=\"example\">示例 2：</strong></p>\n\n<pre>\n<strong>输入：</strong>nums = [3,2,4], target = 6\n<strong>输出：</strong>[1,2]\n</pre>\n\n<p><strong class=\"example\">示例 3：</strong></p>\n\n<pre>\n<strong>输入：</strong>nums = [3,3], target = 6\n<strong>输出：</strong>[0,1]\n</pre>\n\n<p>&nbsp;</p>\n\n<p><strong>提示：</strong></p>\n\n<ul>\n\t<li><code>2 &lt;= nums.length &lt;= 10<sup>4</sup></code></li>\n\t<li><code>-10<sup>9</sup> &lt;= nums[i] &lt;= 10<sup>9</sup></code></li>\n\t<li><code>-10<sup>9</sup> &lt;= target &lt;= 10<sup>9</sup></code></li>\n\t<li><strong>只会存在一个有效答案</strong></li>\n</ul>\n\n<p>&nbsp;</p>\n\n<p><strong>进阶：</strong>你可以想出一个时间复杂度小于 <code>O(n<sup>2</sup>)</code> 的算法吗？</p>\n",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Html2Markdown(tt.args.html); got != tt.want {
				t.Errorf("Html2Markdown() = %v, want %v", got, tt.want)
			}
		})
	}
}

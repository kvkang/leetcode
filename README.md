# leetcode

## 相关材料
- [线性代数][1]: [教学资料][2]

[1]: https://nbviewer.jupyter.org/github/zlotus/notes-linear-algebra/blob/master/ReadMe.ipynb
[2]: https://web.mit.edu/18.06

## 概念
### Median (中位数)
    作用
        将一个集合划分成两个子集,其中一个子集的元素(左子集)永远小于另外一个子集的元素(有子集) 
        Dividing a set into two equal length subsets, that one subset is always greater than the other.
    值
        当集合是奇数时, 不在两个子集中的元素就是中位数
        当集合是偶数时, float(左子集的最大值 + 右子集的最大值) / 2 的值就是中位数
    例: 
        [1, 2, 5]  中位数是 2
        [1, 2, 4, 5] 中位数是 (2+4)/2 = 3
## longest-palindromic-substring (最长回文子串)
    概念
        一个回文字符串无论从那一边开始读(从左往右、从右往左),得到的字符顺序是一样的
        A palindrome is a string which reads the same in both directions
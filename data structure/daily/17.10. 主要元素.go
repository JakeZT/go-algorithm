package main

func main() {

}

/*
这还是道「摩尔投票」模板题。

摩尔投票 ：在集合中寻找可能存在的多数元素，这一元素在输入的序列重复出现并占到了序列元素的一半以上；在第一遍遍历之后应该再进行一个遍历以统计第一次算法遍历的结果出现次数，确定其是否为众数；如果一个序列中没有占到多数的元素，那么第一次的结果就可能是无效的随机元素。

换句话说，每次将两个不同的元素进行「抵消」，如果最后有元素剩余，则「可能」为元素个数大于总数一半的那个。

具体的，我们定义一个变量 xx 来保存那个可能为主要元素的值，cntcnt 用来记录该值的出现次数。然后在遍历数组 numsnums 过程中执行如下逻辑：

如果 cntcnt 为 00：说明之前出现过的 xx 已经被抵消完了，更新一下 xx 为当前值，出现次数为 11：x = nums[i], cnt = 1；
如果 cntcnt 不为 00：说明之前统计的 xx 还没被抵消完，这是根据 nums[i]nums[i] 与 xx 是否相等进行计算即可：cnt += nums[i] == x ? 1 : -1。
当处理完 numsnums 之后，我们得到了一个「可能」的主要元素。注意只是可能，因为我们在处理过程中只使用了 xx 和 cntcnt 来记录，我们是无法确定最后剩下的 xx 是经过多次抵消后剩余的主要元素，还是只是不存在主要元素的数组中的无效随机元素。

因此我们需要再进行一次遍历，检查这个「可能」的主要元素 xx 的出现次数是否超过总数一半。

*/
func majorityElement(nums []int) int {
	c := -1 //主要元素
	count := 0
	for _, num := range nums {
		if count == 0 {
			c = num
		}
		if num == c {
			count++
		} else {
			count--
		}
	}
	count = 0
	for _, num := range nums {
		if num == c {
			count++
		}
	}
	if count*2 > len(nums) {
		return c
	}
	return -1
}

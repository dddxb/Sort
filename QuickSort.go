//快速排序
func QuickSort(a []int) []int {
	/*递归，但此时却没有终结条件，因为a是切片，属于引用传递，并且当最后
	* 至少有两个元素时，在QuerySort会直接排序，如果选区的key不合适，分割后
	* 两部分中的一部分只有一个元素或没有，都不影响下一次的递归
	 */
	if len(a)>1 { // 最少有两个元素，进入QuerySort后会直接进行排序
		key := QuerySort(a) //将a一分为二，算出枢轴值key，是相对中间值的下标
		QuickSort(a[0:key]) //对低子表递归排序
		QuickSort(a[key+1:]) //对高子表递归排序
	}
	return a
}
/*交换切片中子表的记录，使枢轴记录到位，并返回所在位置
* 此时在它之前(后)的记录均不大(小)于它
* 注意，此时并不是平分切片，而是找一个差不多大小的元素，使之左右两侧尽量平衡
* 因为只有涉及到平衡二叉树时排序才会高效
*/
func QuerySort(a []int) int {
	key := a[0] //用切片中的第一个记录作枢轴记录，后期可以进行优化，尽量选取中间数
	i,j := 0,len(a)-1
	for ;i<j; {
		for ;i<j&&a[j]>key; {
			j--
		}
		key,a[j] = a[j],key //将比枢轴值小的记录交换到低端
		for ;i<j&&a[i]<key; {
			i++
		}
		key,a[i] = a[i],key //将比枢轴值大的记录交换到高端
	}
	return i //返回枢轴所在位置
}
/*********************************************************************************/
	//改进
func QuickSort(a []int) []int {
	if len(a)>7 { //第二步优化：当记录个数大于常数时用快速排序
		pivotkey := Partition(a)
		QuickSort(a[0:pivotkey])
		QuickSort(a[pivotkey+1:])
	}else { //第二步优化：当记录个数小于常数时用直接插入排序
		if len(a)==1 {
			return a
		}else {
			for i:=1;i<len(a);i++ { //从第二个开始选取，下标此时为1，切记!!!
				if a[i-1]>a[i] {
					temp := a[i]
					var j int
					for j=i-1;j>=0&&a[j]>temp;j-- {
						a[j+1] = a[j]
					}
					a[j+1] = temp
				}
			}
		}
	}
	return a
}
func Partition(a []int) int {
	//第一步优化：优化选取枢轴值
	middle := len(a)/2     //计算切片中间的元素的下标
	low,high := 0,len(a)-1
	if a[middle]>a[high] { //交换中间与右端数据，保证中间较小
		a[middle],a[high] = a[high],a[middle]
	}
	if a[low]>a[high] { //交换左端与右端数据，保证左端较小
		a[low],a[high] = a[high],a[low]
	}
	if a[middle]>a[low] { //交换中间与左端数据，保证中间较小
		a[middle],a[low] = a[low],a[middle]
	}
	//此时a[low]已经为整个序列表左中右三个关键字的中间值
	temp := a[0]
	i,j := 0,len(a)-1
	for ;i<j; {
		for temp<a[j] {
			j--
		}
		a[i],a[j] = a[j],a[i]
		for temp>a[i] {
			i++
		}
		a[j],a[i] = a[i],a[j]
	}
	return i
}
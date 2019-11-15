//归并排序
func MSort(a []int) []int {
	if len(a)<2 {
		return a
	}
	m := len(a)/2 //将切片平分为左右两部分
	left := MSort(a[0:m]) //递归将左部分归并
	right := MSort(a[m:]) //递归将右部分归并
	result := Merge(left,right) //将左右两部分合并
	return result
}
func Merge(left,right []int) []int {
	result := make([]int,0)
	i,j,l,r := 0,0,len(left),len(right)
	for ;i<l && j<r; {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++ //注意此时的位置，在append之后，即先添加进result，后下标志加一
		}else {
			result = append(result,right[j])
			j++ ////注意此时的位置，在append之后，即先添加进result，后下标志加一
		}
	}
	if i<l { //左部分还有剩余，将剩余元素复制到result中
		result = append(result,left[i:]...)
	}else if j<r { //右部分还有剩余，将剩余元素复制到result中
		result = append(result,right[j:]...)
	}
	return result
}

/********************************************************************************/
//改进
func MSort(a []int) []int {
	if len(a)<2 {
		return a
	}
	var result []int
	m := len(a)/2
	left := MSort(a[0:m])
	right := MSort(a[m:])
	//第一步优化，左右两部分已排好序，只有当左边的最大值大于右边的最小值，才需要对这两部分进行merge操作
	if left[len(left)-1] > right[0] {
		result = Merge(left,right)
	}else {
		result = append(result,left[:]...)
		result = append(result,right[:]...)
	}
	return result
}
func Merge(left,right []int) []int {
	result := make([]int,0)
	//第二步优化，当数据规模足够小的时候，可以使用直接插入排序
	if len(right)+len(left) <= 15 {
		result = append(result,left[:]...)
		result = append(result,right[:]...)
		// 对 left,right合并后的数据执行插入排序
		for i:=1;i<len(result);i++ {
			if result[i-1] > result[i] {
				var j int
				temp := result[i]
				for j=i-1;j>=0&&result[j]>temp;j-- {
					result[j+1] = result[j]
				}
				result[j+1] = temp
			}
		}
	}else {
		i,j,l,r := 0,0,len(left),len(right)
		for ;i<l && j<r; {
			if left[i]<right[j] {
				result = append(result,left[i])
				i++
			}else {
				result = append(result,right[j])
				j++
			}
		}
		if i<l {
			result = append(result,left[i:]...)
		}else {
			result = append(result,right[j:]...)
		}
	}
	return result
}
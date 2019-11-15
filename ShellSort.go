//希尔排序
//对比直接插入排序适用于元素个数少的情况，对于元素个数多的我们可以分组进行
func ShellSort(a [10]int) [10]int {
	increment := len(a)
	for { //第一层for循环控制increment的变化
		increment = increment/3+1
		for i:=increment;i<len(a);i++ { //第二层for循环控制分组的不断改变
			if a[i-increment] > a[i] { //比较每一分组中相邻两数是否前大后小，满足则进行插入
				temp := a[i]
				var j int
				for j=i-increment;j>=0&&a[j]>temp;j-=increment { //第三层for循环控制插入
					a[j+increment] = a[j]
				}
				a[j+increment] = temp
			}
		}
		if increment == 1 {
			break
		}
	}
	return a
}
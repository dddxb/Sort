//直接插入排序
//类似于打牌，对手里的几张牌进行插入排序，适用于元素个数较少的情况
//将一个记录插入到已经排好序的有序表中，从而得到一个新的记录数增1的有序表
func InsertSort(a []int) []int {
	var temp,j int
	for i:=1;i<len(a);i++ {
		if a[i-1]>a[i] { // 需将a[i]插入有序表
			temp = a[i]  //类似于设置哨兵
			for j=i-1;j>=0&&a[j]>temp;j-- {
				a[j+1] = a[j] //依次后移
			}
			a[j+1] = temp //插入到正确位置
		}
	}
	return a
}
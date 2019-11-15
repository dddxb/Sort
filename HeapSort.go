//堆排序
func HeapSort(a [10]int) [10]int {
	b := [11]int{}
	for k := range a {
		b[k+1] = a[k]
	}
	//第一步先构造大顶堆,大顶堆用层序遍历的方式保存为数组或链表
	for i:=(len(b)-1)/2;i>=1;i-- {
		b = HeapAdjust(b,i,len(b)-1)
	}
	//第二步调整大顶堆
	for j:=len(b)-1;j>1;j-- {
		b[j],b[1] = b[1],b[j] //将堆顶记录和当前未经排序子序列的最后一个记录交换
		b = HeapAdjust(b,1,j-1) //将b[1……j-1]重新调整为大顶堆
	}
	for k := range b {
		if k >=1 {
			a[k-1] = b[k]
		}
	}
	return a
}
//需要深刻理解调整大顶堆的过程
func HeapAdjust(b [11]int,s int,m int) [11]int {
	var temp,j int
	temp = b[s]
	for j=s*2;j<=m;j=j*2 { //沿关键字较大的孩子结点向下筛选，j一定存在，j+1不一定存在
		if j<m && b[j] < b[j+1] { //左右孩子都存在的情况下
			j++        //j为左右子孩子中较大的记录的下标
		}
		if temp >= b[j] { //堆顶值大于左右孩子值，不需要调整，退出
			break
		}
		b[s] = b[j] //三者中的最大值调整到堆顶
		s = j //继续沿着左右孩子中较大的孩子向下筛选，此时j变成了新的s
	}
	b[s] = temp
	return b
}
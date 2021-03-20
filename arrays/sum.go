package arrays

func Sum(numbers []int) int {
	sum := 0
	//range 迭代数组返回数组元素的索引和值 _空白标识符忽略索引
	for _, number := range numbers {
		sum += number
	}
	return sum
}

package main

import (
	"fmt"
)

func main() {

	/*
		3.针对问题2的切片, 将最大的数字移动到切片的最后一位 原来的数字都在移动后的切片中都存在, 只是最大的数字再最后一位

	*/
	var nu_list = []int{108, 107, 105, 109, 103, 102, 110, 106}
	for i := 0; i < len(nu_list)-1; i++ {
		fmt.Println(i)
		fmt.Printf("s1: %d\ts2:%d\n", nu_list[i], nu_list[i+1])
		/*
			0
			s1: 108 s2:107	==》	{107, 108, 105, 109, 103, 102, 110, 106}
			1
			s1: 108 s2:105	==》    {107, 105, 108, 109, 103, 102, 110, 106}
			2
			s1: 108 s2:109  ==》    {107, 105, 108, 109, 103, 102, 110, 106}
			3
			s1: 109 s2:103	==》    {107, 105, 108, 103, 109, 102, 110, 106}
			4
			s1: 109 s2:102	==》    {107, 105, 108, 103, 102, 109, 110, 106}
			5
			s1: 109 s2:110	==》    {107, 105, 108, 103, 102, 109, 110, 106}
			6
			s1: 110 s2:106	==》    {107, 105, 108, 103, 102, 109, 106, 110}

			最终输出：[107 105 108 103 102 109 106 110]
		*/

		if nu_list[i] > nu_list[i+1] {
			nu_list[i], nu_list[i+1] = nu_list[i+1], nu_list[i]
		}

	}
	fmt.Println(nu_list)

}

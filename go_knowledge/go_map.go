package go_knowledge

import "fmt"

func MapForNoSort() {
	/*
		由于遍历的时候，遍历v使用的同一块地址，同时这块地址是临时分配的。虽然v的地址没有变化，但v的内容在一直变化，
		当遍历完成后，v的内容是map遍历时最后遍历的元素的值(map遍历无序，每次不确定哪个元素是最后一个元素)。
		程序将v的地址放入到slice中的时候，slice在不断地v的地址插入，由于v一直是那块地址，因此slice中的每个元素记录的都是v的地址。
		因此当打印slice中的内容的时候，都是同一个值
	*/
	m := map[string]int{
		"a": 1,
		"b": 2,
	}
	var bs []*int
	for k, v := range m {
		fmt.Printf("k:[%p].v:[%p]\n", &k, &v) // 这里的输出可以看到，k一直使用同一块内存，v也是这个状况
		bs = append(bs, &v)                   // 对v取了地址
	}

	// 输出
	for _, b := range bs {
		fmt.Println(*b) // 输出都是1或者都是2
	}
}

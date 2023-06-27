package go_knowledge

import (
	"fmt"
	"unsafe"
)

type Course struct {
	price int
	name  string
	url   string
}

func changeCourse(c Course) {
	c.price = 200
}

func GoStructInit() {
	//
	var c1 Course = Course{
		price: 100,
		name:  "scrapy分布式爬虫",
		url:   "", // 注意这里的逗号不能少
	}
	fmt.Printf("%+v\n", c1)

	var c2 Course = Course{100, "scrapy分布式爬虫", ""}
	fmt.Printf("%+v\n", c2)

	var c3 *Course = &Course{100, "scrapy分布式爬虫", ""}
	fmt.Printf("%+v\n", c3)

	var c4 Course = Course{}
	fmt.Printf("%+v\n", c4)
	var c5 Course
	fmt.Printf("%+v\n", c5)
	var c6 *Course = new(Course)
	fmt.Printf("%+v\n", c6)

	// 结构题拷贝
	var c7 Course = Course{50, "scrapy分布式爬虫", ""}
	var c8 Course = c7
	fmt.Printf("%+v\n", c7)
	fmt.Printf("%+v\n", c8)
	c7.price = 100
	fmt.Printf("%+v\n", c7)
	fmt.Printf("%+v\n", c8)

	var c9 *Course = &Course{50, "scrapy分布式爬虫", ""}
	var c10 *Course = c9
	fmt.Printf("%+v\n", c9)
	fmt.Printf("%+v\n", c10)
	c9.price = 100
	fmt.Printf("%+v\n", c9)
	fmt.Printf("%+v\n", c10)

	// slice 结构体
	type slice struct {
		array unsafe.Pointer // 底层数组的地址
		len   int            // 长度
		cap   int            // 容量
	}

	// string 结构体
	type string struct {
		array unsafe.Pointer // 底层数组的指针
		len   int            // 长度
	}

	// map 结构体
	//type hmap struct {
	//	count int
	//	...
	//	buckets unsafe.Pointer  // hash桶地址
	//	...
	//}

	// 在数组与切片章节，我们自习分析了数组与切片在内存形式上的区别。
	//数组只有「体」，切片除了「体」之外，还有「头」部。
	//切片的头部和内容体是分离的，使用指针关联起来。请读者尝试解释一下下面代码的输出结果
	type ArrayStruct struct {
		value [10]int
	}

	type SliceStruct struct {
		value []int
	}

	var as = ArrayStruct{[...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}
	var ss = SliceStruct{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}
	fmt.Println("=========================================")
	fmt.Println(unsafe.Sizeof(as), unsafe.Sizeof(ss))
	fmt.Println("=========================================")
	// 结构体是值传递
	var c Course = Course{
		price: 100,
		name:  "scrapy分布式爬虫",
		url:   "", // 注意这里的逗号不能少
	}
	changeCourse(c)
	fmt.Println(c.price)
}

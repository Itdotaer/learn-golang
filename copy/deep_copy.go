package copy

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) DeepCopy() *Person {
	newPerson := &Person{
		Name: p.Name,
		Age:  p.Age,
	}
	return newPerson
}

func main() {
	// 创建源对象
	person1 := &Person{
		Name: "Alice",
		Age:  30,
	}

	// 进行深拷贝
	person2 := person1.DeepCopy()

	// 修改 person2 的字段
	person2.Name = "Bob"
	person2.Age = 25

	// 打印源对象和目标对象的字段值
	fmt.Println("person1:", person1) // 输出: person1: &{Alice 30}
	fmt.Println("person2:", person2) // 输出: person2: &{Bob 25}
}

package main

import (
	"errors"
	"fmt"
	"google.golang.org/protobuf/runtime/protoimpl"
	"hash/crc32"
	"math/rand"
	"time"
)

type Sayble interface {
	Say()
}

type People struct {
	name string
	age  int32
}

func (s *People) Say(msg string) {
	fmt.Printf("name: %s, age: %d; say: %s", s.name, s.age, msg)
}

type PortalType int

const (
	Pms        PortalType = 1
	NormalSrm  PortalType = 2
	FashionSrm PortalType = 3
	Fashion    PortalType = 4
	Obs        PortalType = 5
	ObsSrm     PortalType = 6
)

type CommonReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Requestor  string `protobuf:"bytes,1,opt,name=requestor,proto3" json:"requestor,omitempty"`
	PortalType *int32 `protobuf:"varint,2,opt,name=portalType,proto3,oneof" json:"portalType,omitempty"` // 1.Normal Ops 用户; 2.Normal SRM 用户; 3.Fashion SRM 用户；4. Fashion Ops用户， 5. SCS Ops用户；6. SCS SRM用户
}

var BranchLayout = "00000102"

func test() (int, error) {
	rand.Seed(time.Now().UnixNano())
	rs := rand.Intn(100)
	if rs%2 == 0 {
		return rs, nil
	}
	return 0, errors.New("test error")
}

type Person struct {
	Name string
	Age  int
}

func modifyPersonList(personList []Person) {
	for i := range personList {
		personList[i].Age += 1
	}
}

func main() {
	var testMap map[int]struct{}
	if value, ok := testMap[0]; ok {
		fmt.Println(value)
	}

	var testArr []struct{}
	fmt.Println(testArr[0])

	personList := []Person{{
		Name: "1",
		Age:  22,
	}, {
		Name: "2",
		Age:  23,
	}}

	fmt.Println(personList)
	modifyPersonList(personList)
	fmt.Println(personList)

	tableIdx := crc32.ChecksumIEEE([]byte("706518")) % 30
	println(fmt.Sprintf("supplier_sales_history_tab_%08d", tableIdx))

	println(fmt.Sprintf("%.2f = Max(%.2f - %.2f, 0) * %d",
		float64(33), float64(44), float64(55), 66))

	//date := time.Now().AddDate(0, 0, -2).Format("2006-01-02")
	//dateTime, _ := time.Parse("2006-01-02", date)
	//table := fmt.Sprintf("inventory_health_history_tab_%s", dateTime.Format(BranchLayout))
	//println(table)

	var (
		rs, counter int
		err         error
	)
	for {
		if counter >= 3 {
			break
		}

		counter++
		rs, err = test()

		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("success")
			break
		}
	}

	fmt.Printf("rs:%d, counter:%d\n", rs, counter)

	//println(fmt.Sprintf("cancel:%v; confirm:%v", nil, nil))
	//
	//rs := crc32.ChecksumIEEE([]byte("df"))
	//println(rs)
	//md := md5.New()
	//md.Write([]byte("123"))
	//fmt.Println(hex.EncodeToString(md.Sum(nil)))
	//
	//params := []string{"a", "b", "5"}
	//portalType := Pms
	//if len(params) >= 3 {
	//	portalTypeInt, _ := strconv.ParseUint(params[2], 10, 64)
	//	portalType = PortalType(portalTypeInt)
	//}
	//
	//fmt.Println(portalType)

	//portalType := Obs
	//
	//println(fmt.Sprintf("test %d", int32(portalType)))
	//println(fmt.Sprintf("(PortalType=%d) failed to query user not in emails", portalType))
	//println(fmt.Errorf("(PortalType=%d) failed to query user not in emails", portalType).Error())

	//sli1 := make([]int, 2, 5)
	//println(len(sli1), cap(sli1))
	//sli1 = append(sli1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17)
	//println(len(sli1), cap(sli1))
	//
	//sli2 := make([]int, 2, 5)
	//sli2[0] = 1
	//println(len(sli2), cap(sli2))
	//
	//fmt.Println(sli1)
	//fmt.Println(sli2)
	//
	//copy(sli2, sli1) // 这一步就证明了，copy操作也是对于len，并且它会将前者清空。
	//fmt.Println(sli1)
	//fmt.Println(sli2)

	//counter := int64(0)
	//incr := func(counter *int64) int64 {
	//	old := atomic.LoadInt64(counter)
	//
	//	flag := false
	//	for !flag {
	//		flag = atomic.CompareAndSwapInt64(counter, old, old+1)
	//	}
	//
	//	return old + 1
	//}
	//
	//routineNum := 10
	//sleepDuration := time.Second
	//wg := sync.WaitGroup{}
	//wg.Add(routineNum)
	//var duplicateMap sync.Map
	//
	//for i := 0; i < routineNum; i++ {
	//	go func() {
	//		for {
	//			newCounter := incr(&counter)
	//			println(newCounter)
	//
	//			if value, ok := duplicateMap.Load(newCounter); ok {
	//				println(fmt.Sprintf("duplicated number: %d", value))
	//			} else {
	//				duplicateMap.Store(newCounter, newCounter)
	//			}
	//
	//			time.Sleep(sleepDuration)
	//		}
	//
	//		wg.Done()
	//	}()
	//}
	//
	//wg.Wait()
	//println("finished")

	//var num int
	//var testArr []People
	//var people People
	//fmt.Println(num)
	//fmt.Println(people)
	//fmt.Println(testArr)
	//
	//people.age = 12
	//people.name = "jiodjf"
	//
	//fmt.Println(people)
	//
	//myMap := make(map[int]int)
	//myArr := make([]int, 0, 100)
	//
	//modifyMap := func(m map[int]int) {
	//	for i := 0; i < 10; i++ {
	//		m[i] = i
	//	}
	//}
	//
	//modifyArr := func(arr []int) {
	//	for i := 0; i < 10; i++ {
	//		arr = append(arr, i)
	//	}
	//
	//	println(arr)
	//}
	//
	//modifyMap(myMap)
	//modifyArr(myArr)
	//
	//println(myArr)
	//
	//fmt.Println(myMap)
	//fmt.Println(myArr)
}

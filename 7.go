package main

import (
	"encoding/gob"
	"sync"
	"sync/atomic"
)

func deepCopy(dst, src interfac{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

// go 语言中的传递都是值传递，传递一个对象，就会把对象拷贝一份传入函数中，传递一个指针，就会把指针拷贝一份传入进去。

// 赋值的时候也是这样，res:=*e 就会把传递的 Example 对象拷贝一份，如果是 res:=e 的话，那么拷贝的就是对象的指针了。

// 而深拷贝和浅拷贝也可以这样理解，深拷贝就是拷贝整个对象，浅拷贝就是拷贝对象指针。

// 对于深度拷贝，go 和其他语言还经常使用序列化后反序列化的形式进行拷贝:

// 实际上 gob 包序列化的时候也是用到了 reflect 包来实现拷贝的。

// 注意： golang 完全是按值传递，所以如果深度拷贝的对象中包含有指针的话，那么深度拷贝后，这些指针也会相同，会导致部分数据共享，要注意这一点

// 设计一个类的时候，我们通常会使用到构造函数，这里类和对象的关系好比模具和构件的关系，
// 对象总是从类中创建的。但是某些场景下是不允许类的调用者直接调用构造函数，也就说对象未
// 必需要从类中衍生出来，现实生活中存在太多案例是通过直接“克隆” 来产生新的对象，而且
// 克隆出来的本体和克隆体看不出任何区别。
// 原型模式不单是一种设计模式，也是一种编程范型。 简单理解原型模式Prototype:不根据类
// 来生成实例，而是根据实例生成新的实例。也就说，如果需要一个和某对象一 模一样的对象，那
// 么就可以使用原型模式。


type Role interface {
	SetHeadColor(string)
	SetEyesColor(string)
	SetTall(int64)
	Show()
	Clone() Role
}

type RoleChinese struct {
	HeadColor string
	EyesColor string
	Tall int64
}

func (pR *RoleChinese) Clone() Role  {
	var pChinese = &RoleChinese{HeadColor: pR.HeadColor, EyesColor: pR.EyesColor, Tall: pR.Tall}
	return pChinese
}

func (pR *RoleChinese) SetHeadColor(color string)  {
	pR.HeadColor = color
}


func (pR *RoleChinese) SetEyesColor(color string) {
	pR.EyesColor = color
}

func (pR *RoleChinese) SetTall(tall int64) {
	pR.Tall = tall
}

func (pR *RoleChinese) Show() {
	fmt.Println("Role's headcolor is:", pR.HeadColor, " EyesColor is:", pR.EyesColor, " tall is:", pR.Tall)
}


func main21(){
	role := &RoleChinese{
			HeadColor: "black",
			EyesColor: "black",
			Tall:      170,
	}
	role.Show()
	// 根据一个实例，生成新的实例
	copyer := role.Clone()
	copyer.Show()
	copyer.SetEyesColor("bule")
	role.Show()
	copyer.Show()
}


// 懒汉模式 非线程安全
// type singleton struct{}

// var instance *singleton

// func GetInstance() *singleton {
// 	if singleton == nil {
// 		instance = &singleton{}
// 	}
// 	return instance
// }


type singleton struct {}

var instance *singleton
var mu sync.Mutex

func GetInstance() *singleton {
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		instance = &singleton{}  // unnecessary locking if instance already created
	}
	return instance
}


// if instance == nil {     // <-- Not yet perfect. since it's not fully atomic
// mu.Lock()
// defer mu.Unlock()

// if instance == nil {
// 		instance = &singleton{}
// }
// }
// return instance
// }

// 这是一个不错的方法，但是还并不是很完美。因为编译器优化没有检查实例存储状态。如果使用 sync/atomic 包的话 就可以自动帮我们加载和设置标记。

import "sync/atomic"

var initialized uint32

func GetInstance1() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	mu.Lock()
	defer mu.Unlock()
	if initialized == 0 {
		instance = &singleton{}
		atomic.StoreUint32(&initialized, 1)
	}
	return instance
}

// 这种方式也是 Go 官方建议使用的方法，使用 Once 来执行函数保证了就算有多个程序在并发调用这个函数，这个函数也只会执行一次。

// import (
// 	"sync"
// )
// type singleton struct {
// }

// var instance *singleton
// var once sync.Once

// func GetInstance() *singleton {
// 	// 使用Once保证创建实例的方法永远只能运行一次
// 	once.Do(func() {
// 			instance = &singleton{}
// 	})
// 	return instance
// }
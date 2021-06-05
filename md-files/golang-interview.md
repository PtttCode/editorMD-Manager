# 知识点： [![](https://segmentfault.com/img/remote/1460000038922264)](https://segmentfault.com/img/remote/1460000038922264)

# 问题目录


[Q1:nil切片和空切片指向的地址一样吗](#Q1)
[Q2:字符串转成byte数组，会发生内存拷贝吗？](#Q2)



------------

# 数据类型

## [map](#map)

------------

# map
- 类似其它语言中的哈希表或字典，以key-value形式存储数据
- key必须是支持==或!=比较运算的类型，不可以是函数、map或slice
- Map通过key查找value比线性搜索快很多
- Map使用make()创建，支持:=这种简写方式
- make([keyType]valueType,cap),cap表示容量，可省略，超出容量时会自动扩容，但尽量提供一个合理的初始值
- 使用len()获取元素个数
- 键值对不存在时自动添加，使用delete()删除某键值对
- map在golang里属于**[引用类型](#引用类型)**，在作为形参或者返回类型时，传递的是引用的地址
- 声明且未初始化的map是**nil map**, 对其任意取值都会返回**（nil, err）**,赋值则会panic
```go
var example map[string]int	# 声明map，此时是nil map
example = make(map[string]int)	# 初始化
```
- map在golang是非线程安全的，所以在多线程情况下根据需要加上读写锁（**[sync.RWMutex](#RWMutex)**）；但是sync.map是线程安全的，似乎性能也强于map


# 值类型
> 值类型包括基本数据类型，int,float,bool,string,以及数组和结构体(struct)。注意：sync.WaitGroup 对象是值类型，不是一个引用类型

>值类型变量声明后，不管是否已经赋值，编译器为其分配内存，此时该值存储于栈上。

>当使用等号=将一个变量的值赋给另一个变量时，如 j = i ,实际上是在内存中将 i 的值进行了拷贝，可以通过 &i 获取变量 i 的内存地址。此时如果修改某个变量的值，不会影响另一个。


# 引用类型(reference type)
>引用类型包括指针，slice切片，map ，chan，interface。
变量直接存放的就是一个内存地址值，这个地址值指向的空间存的才是值。所以修改其中一个，另外一个也会修改（同一个内存地址）。

>引用类型必须申请内存才可以使用，make()是给引用类型申请内存空间。


# sync.RWMutex



------------


# Q1
## nil切片和空切片指向的地址一样吗

> **nil切片和空切片指向的地址不一样。nil空切片引用数组指针地址为0（无指向任何实际地址）
空切片的引用数组指针地址是有的，且固定为一个值**

#### [Q1详解][1]
[1]: https://mp.weixin.qq.com/s/myGJ4TrEoVGqLAN3tbZHMw

------------

# Q2 
## 字符串转成byte数组，会发生内存拷贝吗？

> **字符串转成切片，会产生拷贝。严格来说，只要是发生类型强转都会发生内存拷贝。那么问题来了。
频繁的内存拷贝操作对性能不大友好.**

**StringHeader** 是**字符串**在go的底层结构。

```
    type StringHeader struct {
		 Data uintptr
		 Len  int
    }
```

**SliceHeader** 是**切片**在go的底层结构。
```
    type SliceHeader struct {
		 Data uintptr
		 Len  int
		 Cap  int
    }
```
转换以上两种结构体可以使用unsafe.Pointer转换成通用指针类型后经reflect强制转换
####[Q2详解][2]
[2]: https://mp.weixin.qq.com/s/d80m0hgoKcHfKp4ZXH1M4A


------------


# 知识点： [![](https://segmentfault.com/img/remote/1460000038922264)](https://segmentfault.com/img/remote/1460000038922264)

# 问题目录


[Q1:nil切片和空切片指向的地址一样吗](#Q1)
[Q2:字符串转成byte数组，会发生内存拷贝吗？](#Q2)



------------

# 数据类型

## [map](#map)

------------

# map
- 类似其它语言中的哈希表或字典，以key-value形式存储数据
- key必须是支持==或!=比较运算的类型，不可以是函数、map或slice
- Map通过key查找value比线性搜索快很多
- Map使用make()创建，支持:=这种简写方式
- make([keyType]valueType,cap),cap表示容量，可省略，超出容量时会自动扩容，但尽量提供一个合理的初始值
- 使用len()获取元素个数
- 键值对不存在时自动添加，使用delete()删除某键值对
- map在golang里属于**[引用类型](#引用类型)**，在作为形参或者返回类型时，传递的是引用的地址
- 声明且未初始化的map是**nil map**, 对其任意取值都会返回**（nil, err）**,赋值则会panic
```go
var example map[string]int	# 声明map，此时是nil map
example = make(map[string]int)	# 初始化
```
- map在golang是非线程安全的，所以在多线程情况下根据需要加上读写锁（**[sync.RWMutex](#RWMutex)**）；但是sync.map是线程安全的，似乎性能也强于map


# 值类型
> 值类型包括基本数据类型，int,float,bool,string,以及数组和结构体(struct)。注意：sync.WaitGroup 对象是值类型，不是一个引用类型

>值类型变量声明后，不管是否已经赋值，编译器为其分配内存，此时该值存储于栈上。

>当使用等号=将一个变量的值赋给另一个变量时，如 j = i ,实际上是在内存中将 i 的值进行了拷贝，可以通过 &i 获取变量 i 的内存地址。此时如果修改某个变量的值，不会影响另一个。


# 引用类型(reference type)
>引用类型包括指针，slice切片，map ，chan，interface。
变量直接存放的就是一个内存地址值，这个地址值指向的空间存的才是值。所以修改其中一个，另外一个也会修改（同一个内存地址）。

>引用类型必须申请内存才可以使用，make()是给引用类型申请内存空间。


# sync.RWMutex



------------


# Q1
## nil切片和空切片指向的地址一样吗

> **nil切片和空切片指向的地址不一样。nil空切片引用数组指针地址为0（无指向任何实际地址）
空切片的引用数组指针地址是有的，且固定为一个值**

#### [Q1详解][1]
[1]: https://mp.weixin.qq.com/s/myGJ4TrEoVGqLAN3tbZHMw

------------

# Q2 
## 字符串转成byte数组，会发生内存拷贝吗？

> **字符串转成切片，会产生拷贝。严格来说，只要是发生类型强转都会发生内存拷贝。那么问题来了。
频繁的内存拷贝操作对性能不大友好.**

**StringHeader** 是**字符串**在go的底层结构。

```
    type StringHeader struct {
		 Data uintptr
		 Len  int
    }
```

**SliceHeader** 是**切片**在go的底层结构。
```
    type SliceHeader struct {
		 Data uintptr
		 Len  int
		 Cap  int
    }
```
转换以上两种结构体可以使用unsafe.Pointer转换成通用指针类型后经reflect强制转换
####[Q2详解][2]
[2]: https://mp.weixin.qq.com/s/d80m0hgoKcHfKp4ZXH1M4A


------------



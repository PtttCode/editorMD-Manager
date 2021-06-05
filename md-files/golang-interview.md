# 知识点： [![](https://segmentfault.com/img/remote/1460000038922264)](https://segmentfault.com/img/remote/1460000038922264)

# 问题目录


[Q1:nil切片和空切片指向的地址一样吗](#Q1)
[Q2:字符串转成byte数组，会发生内存拷贝吗？](#Q2)



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

####[Q2详解][2]
[2]: https://mp.weixin.qq.com/s/d80m0hgoKcHfKp4ZXH1M4A


------------



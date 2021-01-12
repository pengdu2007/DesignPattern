package Singleton

import "sync"

// 懒汉方式
// 全局的单例实例在第一次被使用时创建
// 缺点: 非线程安全. 当正在创建时, 有现成来访问, 此时ins = nil 就会再创建, 单例类就会有多个实例了..
type singleton struct {}
var ins *singleton
func GetIns() *singleton {
    if ins == nil {
        ins = &singleton{}
    }
    return ins
}
// 饿汉方式
// 全局的单例实例在类装载时构建
// 缺点: 如果singleton创建初始化比较复杂耗时时, 加载时间会延长
type singleton2 struct {}
var ins2 *singleton2 = &singleton2{}
func GetIns2() *singleton2 {
    return ins2
}

// 懒汉加锁
// 缺点: 虽然解决并发的问题, 但每次加锁是要付出代价的
type singleton3 struct{}
var ins3 *singleton3
var mu sync.Mutex
func GetIns3() *singleton3 {
    mu.Lock()
    defer mu.Unlock()

    if ins3 == nil {
        ins3 = &singleton3{}
    }
    return ins3
}

// 双重锁
// 避免了每次加锁, 提高代码效率
type singleton4 struct {}
var ins4 *singleton4
var mu4 sync.Mutex
func GetIns4() *singleton4 {
    if ins4 == nil {
        mu4.Lock()
        defer mu4.Unlock()
        if ins == nil {
            ins4 = &singleton4{}
        }
    }

    return ins4
}
// sync.Once 实现
type singleton5 struct{}
var ins5 *singleton5
var once sync.Once
func GetIns5() *singleton5 {
    once.Do(func() {
        ins5 = &singleton5{}
    })
    return ins5
}


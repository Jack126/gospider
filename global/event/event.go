package event

import (
	MyError "gospider/global/error"
	"log"
	"strings"
	"sync"
)

// 定义一个全局事件存储变量，本模块只负责存储 键 => 函数 ， 相对容器来说功能稍弱，但是调用更加简单、方便、快捷
var smap sync.Map

// CreateEventFactory 创建一个事件管理工厂
func CreateEventFactory() *Event {
	return &Event{}
}

// Event 定义一个事件管理结构体
type Event struct {
}

// Set 注册事件
func (e *Event) Set(key string, key_func func(args ...interface{})) bool {
	//判断key下是否已有事件
	if _, exists := e.Get(key); exists == false {
		smap.Store(key, key_func)
		return true
	}
	return false
}

// Get 获取事件
func (e *Event) Get(key string) (interface{}, bool) {
	if value, exists := smap.Load(key); exists {
		return value, exists
	}
	return nil, false
}

// Call 执行事件
func (e *Event) Call(key string, args ...interface{}) {
	if value_interface, exists := e.Get(key); exists {
		if fn, ok := value_interface.(func(args ...interface{})); ok {
			fn(args...)
		} else {
			log.Panic(MyError.ErrorsFuncEventNotCall + ", 键名：" + key + ", 相关函数无法调用")
		}

	} else {
		log.Panic(MyError.ErrorsFuncEventNotRegister + ", 键名：" + key)
	}
}

// Delete 删除事件
func (e *Event) Delete(key string) {
	smap.Delete(key)
}

// FuzzyCall 根据键的前缀，模糊调用. 使用请谨慎.
func (e *Event) FuzzyCall(key_pre string) {

	smap.Range(func(key, value interface{}) bool {
		if keyname, ok := key.(string); ok {
			if strings.HasPrefix(keyname, key_pre) {
				e.Call(keyname)
			}
		}
		return true
	})
}

package tmallgenie

import (
	"sync"
)

// IotCommonDeviceProperty 结构体
type IotCommonDeviceProperty struct {
	// 异常检测项名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 异常检测项值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}

var poolIotCommonDeviceProperty = sync.Pool{
	New: func() any {
		return new(IotCommonDeviceProperty)
	},
}

// GetIotCommonDeviceProperty() 从对象池中获取IotCommonDeviceProperty
func GetIotCommonDeviceProperty() *IotCommonDeviceProperty {
	return poolIotCommonDeviceProperty.Get().(*IotCommonDeviceProperty)
}

// ReleaseIotCommonDeviceProperty 释放IotCommonDeviceProperty
func ReleaseIotCommonDeviceProperty(v *IotCommonDeviceProperty) {
	v.Name = ""
	v.Value = ""
	poolIotCommonDeviceProperty.Put(v)
}

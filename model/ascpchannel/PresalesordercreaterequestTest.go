package ascpchannel

import (
	"sync"
)

// PresalesordercreaterequestTest 结构体
type PresalesordercreaterequestTest struct {
	// 预售单
	PresalesOrder *PresalesorderTest `json:"presales_order,omitempty" xml:"presales_order,omitempty"`
}

var poolPresalesordercreaterequestTest = sync.Pool{
	New: func() any {
		return new(PresalesordercreaterequestTest)
	},
}

// GetPresalesordercreaterequestTest() 从对象池中获取PresalesordercreaterequestTest
func GetPresalesordercreaterequestTest() *PresalesordercreaterequestTest {
	return poolPresalesordercreaterequestTest.Get().(*PresalesordercreaterequestTest)
}

// ReleasePresalesordercreaterequestTest 释放PresalesordercreaterequestTest
func ReleasePresalesordercreaterequestTest(v *PresalesordercreaterequestTest) {
	v.PresalesOrder = nil
	poolPresalesordercreaterequestTest.Put(v)
}

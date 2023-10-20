package promotion

import (
	"sync"
)

// LogicGroup 结构体
type LogicGroup struct {
	// 五道口参与者名称
	WdkGroupName string `json:"wdk_group_name,omitempty" xml:"wdk_group_name,omitempty"`
	// 参与者分组序号
	Number int64 `json:"number,omitempty" xml:"number,omitempty"`
	// 逻辑分组类型  COMMON(1, &#34;普通分组&#34;), EXCHANGE(2, &#34;换购分组&#34;), BUY_GIFT(3, &#34;买赠分组&#34;), EXCHANGE_TJ_OVERLAY(4, &#34;特价换购叠加分组&#34;),
	LogicGroupType int64 `json:"logic_group_type,omitempty" xml:"logic_group_type,omitempty"`
}

var poolLogicGroup = sync.Pool{
	New: func() any {
		return new(LogicGroup)
	},
}

// GetLogicGroup() 从对象池中获取LogicGroup
func GetLogicGroup() *LogicGroup {
	return poolLogicGroup.Get().(*LogicGroup)
}

// ReleaseLogicGroup 释放LogicGroup
func ReleaseLogicGroup(v *LogicGroup) {
	v.WdkGroupName = ""
	v.Number = 0
	v.LogicGroupType = 0
	poolLogicGroup.Put(v)
}

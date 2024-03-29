package traveltrade

import (
	"sync"
)

// OrderTipFormInfo 结构体
type OrderTipFormInfo struct {
	// 字段可接收数据类型
	AcceptTypes []string `json:"accept_types,omitempty" xml:"accept_types>string,omitempty"`
	// 选择项, 当type = select / selects 的时候存在
	Options []OrderTipOptionInfo `json:"options,omitempty" xml:"options>order_tip_option_info,omitempty"`
	// 字段对应的值，多选值
	Values []string `json:"values,omitempty" xml:"values>string,omitempty"`
	// 字段具体描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 字段中文名称
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 字段属性名，在做服务信息回传时需要用到的KEY
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 字段对应的值，单选值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 最大数量 (如最大文件上传数量, 最大图片上传数量等)
	MaxNum int64 `json:"max_num,omitempty" xml:"max_num,omitempty"`
	// 字段匹配规则
	Rule *OrderTipRuleInfo `json:"rule,omitempty" xml:"rule,omitempty"`
}

var poolOrderTipFormInfo = sync.Pool{
	New: func() any {
		return new(OrderTipFormInfo)
	},
}

// GetOrderTipFormInfo() 从对象池中获取OrderTipFormInfo
func GetOrderTipFormInfo() *OrderTipFormInfo {
	return poolOrderTipFormInfo.Get().(*OrderTipFormInfo)
}

// ReleaseOrderTipFormInfo 释放OrderTipFormInfo
func ReleaseOrderTipFormInfo(v *OrderTipFormInfo) {
	v.AcceptTypes = v.AcceptTypes[:0]
	v.Options = v.Options[:0]
	v.Values = v.Values[:0]
	v.Desc = ""
	v.Title = ""
	v.Name = ""
	v.Value = ""
	v.MaxNum = 0
	v.Rule = nil
	poolOrderTipFormInfo.Put(v)
}

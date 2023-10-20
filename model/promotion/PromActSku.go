package promotion

import (
	"sync"
)

// PromActSku 结构体
type PromActSku struct {
	// 参与者id
	ParticipateId string `json:"participate_id,omitempty" xml:"participate_id,omitempty"`
	// 参与者名称
	ParticipateName string `json:"participate_name,omitempty" xml:"participate_name,omitempty"`
	// 商家逻辑分组序号
	LogicGroupNumber int64 `json:"logic_group_number,omitempty" xml:"logic_group_number,omitempty"`
	// 参与者类型 SKU_CODE(1, &#34;商品skuCode&#34;), SHOP(2, &#34;店铺&#34;), CATEGORY(3, &#34;品类&#34;)
	ParticipateType int64 `json:"participate_type,omitempty" xml:"participate_type,omitempty"`
}

var poolPromActSku = sync.Pool{
	New: func() any {
		return new(PromActSku)
	},
}

// GetPromActSku() 从对象池中获取PromActSku
func GetPromActSku() *PromActSku {
	return poolPromActSku.Get().(*PromActSku)
}

// ReleasePromActSku 释放PromActSku
func ReleasePromActSku(v *PromActSku) {
	v.ParticipateId = ""
	v.ParticipateName = ""
	v.LogicGroupNumber = 0
	v.ParticipateType = 0
	poolPromActSku.Put(v)
}

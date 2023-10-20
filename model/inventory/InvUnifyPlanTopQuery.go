package inventory

import (
	"sync"
)

// InvUnifyPlanTopQuery 结构体
type InvUnifyPlanTopQuery struct {
	// 生成计划库存的外部单据号
	PlanOrderId string `json:"plan_order_id,omitempty" xml:"plan_order_id,omitempty"`
	// 商品或者货品的id，计划建在哪，就用哪个id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// skuid。如果是货品，则skuid是0
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// item_id的类型，1是前端宝贝，2是后端货品
	ItemType int64 `json:"item_type,omitempty" xml:"item_type,omitempty"`
}

var poolInvUnifyPlanTopQuery = sync.Pool{
	New: func() any {
		return new(InvUnifyPlanTopQuery)
	},
}

// GetInvUnifyPlanTopQuery() 从对象池中获取InvUnifyPlanTopQuery
func GetInvUnifyPlanTopQuery() *InvUnifyPlanTopQuery {
	return poolInvUnifyPlanTopQuery.Get().(*InvUnifyPlanTopQuery)
}

// ReleaseInvUnifyPlanTopQuery 释放InvUnifyPlanTopQuery
func ReleaseInvUnifyPlanTopQuery(v *InvUnifyPlanTopQuery) {
	v.PlanOrderId = ""
	v.ItemId = 0
	v.SkuId = 0
	v.ItemType = 0
	poolInvUnifyPlanTopQuery.Put(v)
}

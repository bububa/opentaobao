package lstlogistics

import (
	"sync"
)

// LstThirdPartMainShipOrderDto 结构体
type LstThirdPartMainShipOrderDto struct {
	// 货品列表
	Details []LstThirdPartDetailShipOrderDto `json:"details,omitempty" xml:"details>lst_third_part_detail_ship_order_dto,omitempty"`
	// 签收时间
	SignTime string `json:"sign_time,omitempty" xml:"sign_time,omitempty"`
	// 发货单状态：NEW ---&gt;新建，LOAD_WAIT---&gt;待装车，LOAD_SUCCESS---&gt;已装车，SIGN_SUCCESS---&gt;签收，SIGN_PART_SUCCESS---&gt;部分签收，SIGN_FAILED---&gt;拒签，CANCEL---&gt;取消
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 外部订单ID
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 发货单更新时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 发货单生成时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 装车时间
	LoadTime string `json:"load_time,omitempty" xml:"load_time,omitempty"`
	// 订单
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 发货单ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolLstThirdPartMainShipOrderDto = sync.Pool{
	New: func() any {
		return new(LstThirdPartMainShipOrderDto)
	},
}

// GetLstThirdPartMainShipOrderDto() 从对象池中获取LstThirdPartMainShipOrderDto
func GetLstThirdPartMainShipOrderDto() *LstThirdPartMainShipOrderDto {
	return poolLstThirdPartMainShipOrderDto.Get().(*LstThirdPartMainShipOrderDto)
}

// ReleaseLstThirdPartMainShipOrderDto 释放LstThirdPartMainShipOrderDto
func ReleaseLstThirdPartMainShipOrderDto(v *LstThirdPartMainShipOrderDto) {
	v.Details = v.Details[:0]
	v.SignTime = ""
	v.Status = ""
	v.OutOrderId = ""
	v.GmtModified = ""
	v.GmtCreate = ""
	v.LoadTime = ""
	v.BizOrderId = 0
	v.Id = 0
	poolLstThirdPartMainShipOrderDto.Put(v)
}

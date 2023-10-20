package tmallnr

import (
	"sync"
)

// NrInventoryUpdateReqDto 结构体
type NrInventoryUpdateReqDto struct {
	// 更新库存的列表值
	DetailList []NrInventoryCheckDetailDto `json:"detail_list,omitempty" xml:"detail_list>nr_inventory_check_detail_dto,omitempty"`
	// 定时送dss，jsd极速达
	BizIdentity string `json:"biz_identity,omitempty" xml:"biz_identity,omitempty"`
	// 幂等值
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 门店编号
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 1表示全量，2表示增量
	CheckMode int64 `json:"check_mode,omitempty" xml:"check_mode,omitempty"`
	// 商家的sellerID，如果是零售商需要给品牌的sellerId
	OwnerId int64 `json:"owner_id,omitempty" xml:"owner_id,omitempty"`
	// 默认为6：门店库存，2：电商库存
	StoreType int64 `json:"store_type,omitempty" xml:"store_type,omitempty"`
}

var poolNrInventoryUpdateReqDto = sync.Pool{
	New: func() any {
		return new(NrInventoryUpdateReqDto)
	},
}

// GetNrInventoryUpdateReqDto() 从对象池中获取NrInventoryUpdateReqDto
func GetNrInventoryUpdateReqDto() *NrInventoryUpdateReqDto {
	return poolNrInventoryUpdateReqDto.Get().(*NrInventoryUpdateReqDto)
}

// ReleaseNrInventoryUpdateReqDto 释放NrInventoryUpdateReqDto
func ReleaseNrInventoryUpdateReqDto(v *NrInventoryUpdateReqDto) {
	v.DetailList = v.DetailList[:0]
	v.BizIdentity = ""
	v.OrderId = ""
	v.StoreCode = ""
	v.CheckMode = 0
	v.OwnerId = 0
	v.StoreType = 0
	poolNrInventoryUpdateReqDto.Put(v)
}

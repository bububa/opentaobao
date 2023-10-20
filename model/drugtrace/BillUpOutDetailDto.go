package drugtrace

import (
	"sync"
)

// BillUpOutDetailDto 结构体
type BillUpOutDetailDto struct {
	// 药品信息数据
	DrugInfosDtoList []DrugInfosDto `json:"drug_infos_dto_list,omitempty" xml:"drug_infos_dto_list>drug_infos_dto,omitempty"`
	// 单据编码
	BillCode string `json:"bill_code,omitempty" xml:"bill_code,omitempty"`
	// 单据类型描述
	BillTypeName string `json:"bill_type_name,omitempty" xml:"bill_type_name,omitempty"`
	// 单据类型
	BillType string `json:"bill_type,omitempty" xml:"bill_type,omitempty"`
	// 发货企业名称
	EntSendName string `json:"ent_send_name,omitempty" xml:"ent_send_name,omitempty"`
	// 发货企业的ref_ent_id
	EntSendId string `json:"ent_send_id,omitempty" xml:"ent_send_id,omitempty"`
	// 收货企业名称
	EntRecvName string `json:"ent_recv_name,omitempty" xml:"ent_recv_name,omitempty"`
	// 收货企业ref_ent_id
	EntRecvId string `json:"ent_recv_id,omitempty" xml:"ent_recv_id,omitempty"`
	// 单据日期
	StoreOutDate string `json:"store_out_date,omitempty" xml:"store_out_date,omitempty"`
	// 最后更新时间
	UpdateDate string `json:"update_date,omitempty" xml:"update_date,omitempty"`
}

var poolBillUpOutDetailDto = sync.Pool{
	New: func() any {
		return new(BillUpOutDetailDto)
	},
}

// GetBillUpOutDetailDto() 从对象池中获取BillUpOutDetailDto
func GetBillUpOutDetailDto() *BillUpOutDetailDto {
	return poolBillUpOutDetailDto.Get().(*BillUpOutDetailDto)
}

// ReleaseBillUpOutDetailDto 释放BillUpOutDetailDto
func ReleaseBillUpOutDetailDto(v *BillUpOutDetailDto) {
	v.DrugInfosDtoList = v.DrugInfosDtoList[:0]
	v.BillCode = ""
	v.BillTypeName = ""
	v.BillType = ""
	v.EntSendName = ""
	v.EntSendId = ""
	v.EntRecvName = ""
	v.EntRecvId = ""
	v.StoreOutDate = ""
	v.UpdateDate = ""
	poolBillUpOutDetailDto.Put(v)
}

package simba

import (
	"sync"
)

// RptEffectEntityDto 结构体
type RptEffectEntityDto struct {
	// 日期
	Thedate string `json:"thedate,omitempty" xml:"thedate,omitempty"`
	// 计划id
	Campaignid string `json:"campaignid,omitempty" xml:"campaignid,omitempty"`
	// 推广组id
	Adgroupid string `json:"adgroupid,omitempty" xml:"adgroupid,omitempty"`
	// 直接成交金额
	Directtransaction string `json:"directtransaction,omitempty" xml:"directtransaction,omitempty"`
	// 间接成交金额
	Indirecttransaction string `json:"indirecttransaction,omitempty" xml:"indirecttransaction,omitempty"`
	// 直接成交笔数
	Directtransactionshipping string `json:"directtransactionshipping,omitempty" xml:"directtransactionshipping,omitempty"`
	// 间接成交笔数
	Indirecttransactionshipping string `json:"indirecttransactionshipping,omitempty" xml:"indirecttransactionshipping,omitempty"`
	// 收藏宝贝数
	Favitemtotal string `json:"favitemtotal,omitempty" xml:"favitemtotal,omitempty"`
	// 收藏店铺数
	Favshoptotal string `json:"favshoptotal,omitempty" xml:"favshoptotal,omitempty"`
	// 投入产出比
	Roi string `json:"roi,omitempty" xml:"roi,omitempty"`
	// 点击转化率
	Coverage string `json:"coverage,omitempty" xml:"coverage,omitempty"`
	// 直接购物车数
	Directcarttotal string `json:"directcarttotal,omitempty" xml:"directcarttotal,omitempty"`
	// 间接购物车数
	Indirectcarttotal string `json:"indirectcarttotal,omitempty" xml:"indirectcarttotal,omitempty"`
	// 人群名称
	Crowdname string `json:"crowdname,omitempty" xml:"crowdname,omitempty"`
	// 流量来源，设备和网络，包含值：PC站内，PC站外,移动站内，移动站外
	Source string `json:"source,omitempty" xml:"source,omitempty"`
	// 人群id
	Crowdid string `json:"crowdid,omitempty" xml:"crowdid,omitempty"`
}

var poolRptEffectEntityDto = sync.Pool{
	New: func() any {
		return new(RptEffectEntityDto)
	},
}

// GetRptEffectEntityDto() 从对象池中获取RptEffectEntityDto
func GetRptEffectEntityDto() *RptEffectEntityDto {
	return poolRptEffectEntityDto.Get().(*RptEffectEntityDto)
}

// ReleaseRptEffectEntityDto 释放RptEffectEntityDto
func ReleaseRptEffectEntityDto(v *RptEffectEntityDto) {
	v.Thedate = ""
	v.Campaignid = ""
	v.Adgroupid = ""
	v.Directtransaction = ""
	v.Indirecttransaction = ""
	v.Directtransactionshipping = ""
	v.Indirecttransactionshipping = ""
	v.Favitemtotal = ""
	v.Favshoptotal = ""
	v.Roi = ""
	v.Coverage = ""
	v.Directcarttotal = ""
	v.Indirectcarttotal = ""
	v.Crowdname = ""
	v.Source = ""
	v.Crowdid = ""
	poolRptEffectEntityDto.Put(v)
}

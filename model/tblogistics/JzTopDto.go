package tblogistics

import (
	"sync"
)

// JzTopDto 结构体
type JzTopDto struct {
	// 快递公司列表
	Expresses []TPDto `json:"expresses,omitempty" xml:"expresses>tp_dto,omitempty"`
	// 安装公司列表
	InsTps []TPDto `json:"ins_tps,omitempty" xml:"ins_tps>tp_dto,omitempty"`
	// 物流公司列表
	LgCps []TPDto `json:"lg_cps,omitempty" xml:"lg_cps>tp_dto,omitempty"`
	// 商品对应的服务信息
	GoodsRelations string `json:"goods_relations,omitempty" xml:"goods_relations,omitempty"`
	// 是否支持快递
	SupportDelivery bool `json:"support_delivery,omitempty" xml:"support_delivery,omitempty"`
	// 是否支持安装
	SupportInstall bool `json:"support_install,omitempty" xml:"support_install,omitempty"`
	// 是否支持修改安装地址
	SuppModifyInsAdd bool `json:"supp_modify_ins_add,omitempty" xml:"supp_modify_ins_add,omitempty"`
}

var poolJzTopDto = sync.Pool{
	New: func() any {
		return new(JzTopDto)
	},
}

// GetJzTopDto() 从对象池中获取JzTopDto
func GetJzTopDto() *JzTopDto {
	return poolJzTopDto.Get().(*JzTopDto)
}

// ReleaseJzTopDto 释放JzTopDto
func ReleaseJzTopDto(v *JzTopDto) {
	v.Expresses = v.Expresses[:0]
	v.InsTps = v.InsTps[:0]
	v.LgCps = v.LgCps[:0]
	v.GoodsRelations = ""
	v.SupportDelivery = false
	v.SupportInstall = false
	v.SuppModifyInsAdd = false
	poolJzTopDto.Put(v)
}

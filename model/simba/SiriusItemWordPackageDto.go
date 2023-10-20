package simba

import (
	"sync"
)

// SiriusItemWordPackageDto 结构体
type SiriusItemWordPackageDto struct {
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 词包名称
	WordPackageName string `json:"word_package_name,omitempty" xml:"word_package_name,omitempty"`
	// 词包id
	WordPackageId int64 `json:"word_package_id,omitempty" xml:"word_package_id,omitempty"`
	// 词包在线状态（0:关闭，1:开启）
	OnlineStatus int64 `json:"online_status,omitempty" xml:"online_status,omitempty"`
	// pc端出价
	PcBidPrice int64 `json:"pc_bid_price,omitempty" xml:"pc_bid_price,omitempty"`
	// 词包类型（1:流量智选，19:捡漏词包）
	PackageType int64 `json:"package_type,omitempty" xml:"package_type,omitempty"`
	// 无线端出价
	WlBidPrice int64 `json:"wl_bid_price,omitempty" xml:"wl_bid_price,omitempty"`
}

var poolSiriusItemWordPackageDto = sync.Pool{
	New: func() any {
		return new(SiriusItemWordPackageDto)
	},
}

// GetSiriusItemWordPackageDto() 从对象池中获取SiriusItemWordPackageDto
func GetSiriusItemWordPackageDto() *SiriusItemWordPackageDto {
	return poolSiriusItemWordPackageDto.Get().(*SiriusItemWordPackageDto)
}

// ReleaseSiriusItemWordPackageDto 释放SiriusItemWordPackageDto
func ReleaseSiriusItemWordPackageDto(v *SiriusItemWordPackageDto) {
	v.GmtModified = ""
	v.WordPackageName = ""
	v.WordPackageId = 0
	v.OnlineStatus = 0
	v.PcBidPrice = 0
	v.PackageType = 0
	v.WlBidPrice = 0
	poolSiriusItemWordPackageDto.Put(v)
}

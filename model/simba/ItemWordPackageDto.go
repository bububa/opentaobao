package simba

import (
	"sync"
)

// ItemWordPackageDto 结构体
type ItemWordPackageDto struct {
	// 词包名称（可选：流量智选词包、捡漏词包）
	WordPackageName string `json:"word_package_name,omitempty" xml:"word_package_name,omitempty"`
	// 词包id（1-流量智选，2-捡漏词包）
	WordPackageId int64 `json:"word_package_id,omitempty" xml:"word_package_id,omitempty"`
	// 开/关词包（0-关闭，1-开启）
	OnlineStatus int64 `json:"online_status,omitempty" xml:"online_status,omitempty"`
	// pc端出价（单位为分）
	PcBidPrice int64 `json:"pc_bid_price,omitempty" xml:"pc_bid_price,omitempty"`
	// 词包类型（1-流量智选， 19-捡漏词包）
	PackageType int64 `json:"package_type,omitempty" xml:"package_type,omitempty"`
	// 无线端出价（单位为分）
	WlBidPrice int64 `json:"wl_bid_price,omitempty" xml:"wl_bid_price,omitempty"`
}

var poolItemWordPackageDto = sync.Pool{
	New: func() any {
		return new(ItemWordPackageDto)
	},
}

// GetItemWordPackageDto() 从对象池中获取ItemWordPackageDto
func GetItemWordPackageDto() *ItemWordPackageDto {
	return poolItemWordPackageDto.Get().(*ItemWordPackageDto)
}

// ReleaseItemWordPackageDto 释放ItemWordPackageDto
func ReleaseItemWordPackageDto(v *ItemWordPackageDto) {
	v.WordPackageName = ""
	v.WordPackageId = 0
	v.OnlineStatus = 0
	v.PcBidPrice = 0
	v.PackageType = 0
	v.WlBidPrice = 0
	poolItemWordPackageDto.Put(v)
}

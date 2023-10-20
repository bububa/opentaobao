package axintrade

import (
	"sync"
)

// PackageTicketInfoDto 结构体
type PackageTicketInfoDto struct {
	// 门票产品信息
	TicketInfo *TicketInfoDto `json:"ticket_info,omitempty" xml:"ticket_info,omitempty"`
	// 门票张数
	TicketNum int64 `json:"ticket_num,omitempty" xml:"ticket_num,omitempty"`
}

var poolPackageTicketInfoDto = sync.Pool{
	New: func() any {
		return new(PackageTicketInfoDto)
	},
}

// GetPackageTicketInfoDto() 从对象池中获取PackageTicketInfoDto
func GetPackageTicketInfoDto() *PackageTicketInfoDto {
	return poolPackageTicketInfoDto.Get().(*PackageTicketInfoDto)
}

// ReleasePackageTicketInfoDto 释放PackageTicketInfoDto
func ReleasePackageTicketInfoDto(v *PackageTicketInfoDto) {
	v.TicketInfo = nil
	v.TicketNum = 0
	poolPackageTicketInfoDto.Put(v)
}

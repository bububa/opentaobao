package tblogistics

import (
	"sync"
)

// PackageDto 结构体
type PackageDto struct {
	// 正向出库回传二段快递公司 逆向出库回传逆向退货商家仓的快递公司
	TmsCpCode string `json:"tms_cp_code,omitempty" xml:"tms_cp_code,omitempty"`
	// 正向出库回传二段运单号 逆向出库回传逆向退回商家仓的运单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 包裹重量(单位：克)，小数点后2位
	Weight string `json:"weight,omitempty" xml:"weight,omitempty"`
	// 包裹长度 (单位：cm)，小数点后2位
	Length string `json:"length,omitempty" xml:"length,omitempty"`
	// 包裹宽度 (单位：cm)，小数点后2位
	Width string `json:"width,omitempty" xml:"width,omitempty"`
	// 包裹高度（单位：cm)，小数点后2位
	Height string `json:"height,omitempty" xml:"height,omitempty"`
	// 包裹仓储单号
	PackageLine string `json:"package_line,omitempty" xml:"package_line,omitempty"`
	// wms出库包裹单号
	OutPackageCode string `json:"out_package_code,omitempty" xml:"out_package_code,omitempty"`
}

var poolPackageDto = sync.Pool{
	New: func() any {
		return new(PackageDto)
	},
}

// GetPackageDto() 从对象池中获取PackageDto
func GetPackageDto() *PackageDto {
	return poolPackageDto.Get().(*PackageDto)
}

// ReleasePackageDto 释放PackageDto
func ReleasePackageDto(v *PackageDto) {
	v.TmsCpCode = ""
	v.MailNo = ""
	v.Weight = ""
	v.Length = ""
	v.Width = ""
	v.Height = ""
	v.PackageLine = ""
	v.OutPackageCode = ""
	poolPackageDto.Put(v)
}

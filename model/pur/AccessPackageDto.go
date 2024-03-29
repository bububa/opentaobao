package pur

import (
	"sync"
)

// AccessPackageDto 结构体
type AccessPackageDto struct {
	// 图片地址列表
	ImgUrlList []string `json:"img_url_list,omitempty" xml:"img_url_list>string,omitempty"`
	// 套餐关联的产品source列表
	ProductSourceValues []string `json:"product_source_values,omitempty" xml:"product_source_values>string,omitempty"`
	// 套餐标识
	SecurityId string `json:"security_id,omitempty" xml:"security_id,omitempty"`
	// 第三方商城标识
	DataSource string `json:"data_source,omitempty" xml:"data_source,omitempty"`
	// 套餐名称
	PackageName string `json:"package_name,omitempty" xml:"package_name,omitempty"`
	// 套餐描述
	PackageDesc string `json:"package_desc,omitempty" xml:"package_desc,omitempty"`
	// 套餐详情
	DetailInfo string `json:"detail_info,omitempty" xml:"detail_info,omitempty"`
}

var poolAccessPackageDto = sync.Pool{
	New: func() any {
		return new(AccessPackageDto)
	},
}

// GetAccessPackageDto() 从对象池中获取AccessPackageDto
func GetAccessPackageDto() *AccessPackageDto {
	return poolAccessPackageDto.Get().(*AccessPackageDto)
}

// ReleaseAccessPackageDto 释放AccessPackageDto
func ReleaseAccessPackageDto(v *AccessPackageDto) {
	v.ImgUrlList = v.ImgUrlList[:0]
	v.ProductSourceValues = v.ProductSourceValues[:0]
	v.SecurityId = ""
	v.DataSource = ""
	v.PackageName = ""
	v.PackageDesc = ""
	v.DetailInfo = ""
	poolAccessPackageDto.Put(v)
}

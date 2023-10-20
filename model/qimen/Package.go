package qimen

import (
	"sync"
)

// Package 结构体
type Package struct {
	// 包材信息
	PackageMaterialList []PackageMaterial `json:"packageMaterialList,omitempty" xml:"packageMaterialList>package_material,omitempty"`
	// 商品列表
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
	// 物流公司编码(SF=顺丰、EMS=标准快递、EYB=经济快件、ZJS=宅急送、YTO=圆通、ZTO=中通 (ZTO)、HTKY=百世汇通、 UC=优速、STO=申通、TTKDEX=天天快递、QFKD=全峰、FAST=快捷、POSTB=邮政小 包、GTO=国通、YUNDA=韵达、JD=京东配送、DD=当当宅配、 AMAZON=亚马逊物流、OTHER=其他;只传英文编码)
	LogisticsCode string `json:"logisticsCode,omitempty" xml:"logisticsCode,omitempty"`
	// 物流公司名称
	LogisticsName string `json:"logisticsName,omitempty" xml:"logisticsName,omitempty"`
	// 运单号
	ExpressCode string `json:"expressCode,omitempty" xml:"expressCode,omitempty"`
	// 包裹编号
	PackageCode string `json:"packageCode,omitempty" xml:"packageCode,omitempty"`
	// 包裹长度(单位：厘米)
	Length string `json:"length,omitempty" xml:"length,omitempty"`
	// 包裹宽度(单位：厘米)
	Width string `json:"width,omitempty" xml:"width,omitempty"`
	// 包裹高度(单位：厘米)
	Height string `json:"height,omitempty" xml:"height,omitempty"`
	// 包裹理论重量(单位：千克)
	TheoreticalWeight string `json:"theoreticalWeight,omitempty" xml:"theoreticalWeight,omitempty"`
	// 包裹重量(单位：千克)
	Weight string `json:"weight,omitempty" xml:"weight,omitempty"`
	// 包裹体积(单位：升)
	Volume string `json:"volume,omitempty" xml:"volume,omitempty"`
	// 发票号
	InvoiceNo string `json:"invoiceNo,omitempty" xml:"invoiceNo,omitempty"`
	// 签收人姓名, string (50) ，必填
	SignUserName string `json:"signUserName,omitempty" xml:"signUserName,omitempty"`
	// 当前状态操作时间, string (19) , YYYY-MM-DD HH:MM:SS
	SignTime string `json:"signTime,omitempty" xml:"signTime,omitempty"`
	// 状态, sign-已签收string (50)
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 备注, string (500) ,
	Remarks string `json:"remarks,omitempty" xml:"remarks,omitempty"`
}

var poolPackage = sync.Pool{
	New: func() any {
		return new(Package)
	},
}

// GetPackage() 从对象池中获取Package
func GetPackage() *Package {
	return poolPackage.Get().(*Package)
}

// ReleasePackage 释放Package
func ReleasePackage(v *Package) {
	v.PackageMaterialList = v.PackageMaterialList[:0]
	v.Items = v.Items[:0]
	v.LogisticsCode = ""
	v.LogisticsName = ""
	v.ExpressCode = ""
	v.PackageCode = ""
	v.Length = ""
	v.Width = ""
	v.Height = ""
	v.TheoreticalWeight = ""
	v.Weight = ""
	v.Volume = ""
	v.InvoiceNo = ""
	v.SignUserName = ""
	v.SignTime = ""
	v.Status = ""
	v.Remarks = ""
	poolPackage.Put(v)
}

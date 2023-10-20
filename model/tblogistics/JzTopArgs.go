package tblogistics

import (
	"sync"
)

// JzTopArgs 结构体
type JzTopArgs struct {
	// 运单号,用快递或商家自有发货时,必填
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 包裹数量
	PackageNumber string `json:"package_number,omitempty" xml:"package_number,omitempty"`
	// 包裹备注
	PackageRemark string `json:"package_remark,omitempty" xml:"package_remark,omitempty"`
	// 包裹体积
	PackageVolume string `json:"package_volume,omitempty" xml:"package_volume,omitempty"`
	// 包裹重量
	PackageWeight string `json:"package_weight,omitempty" xml:"package_weight,omitempty"`
	// 自有物流公司名称
	ZyCompany string `json:"zy_company,omitempty" xml:"zy_company,omitempty"`
	// 自有物流发货时间,时间不能早于当前时间
	ZyConsignTime string `json:"zy_consign_time,omitempty" xml:"zy_consign_time,omitempty"`
	// 自有物流公司电话
	ZyPhoneNumber string `json:"zy_phone_number,omitempty" xml:"zy_phone_number,omitempty"`
}

var poolJzTopArgs = sync.Pool{
	New: func() any {
		return new(JzTopArgs)
	},
}

// GetJzTopArgs() 从对象池中获取JzTopArgs
func GetJzTopArgs() *JzTopArgs {
	return poolJzTopArgs.Get().(*JzTopArgs)
}

// ReleaseJzTopArgs 释放JzTopArgs
func ReleaseJzTopArgs(v *JzTopArgs) {
	v.MailNo = ""
	v.PackageNumber = ""
	v.PackageRemark = ""
	v.PackageVolume = ""
	v.PackageWeight = ""
	v.ZyCompany = ""
	v.ZyConsignTime = ""
	v.ZyPhoneNumber = ""
	poolJzTopArgs.Put(v)
}

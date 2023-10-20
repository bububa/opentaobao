package alsc

import (
	"sync"
)

// VoucherTemplateSettingOpenInfo 结构体
type VoucherTemplateSettingOpenInfo struct {
	// 适配的券模板id列表，为空代表适配全部
	VoucherTemplateIds []string `json:"voucher_template_ids,omitempty" xml:"voucher_template_ids>string,omitempty"`
	// 渠道POS,移动门店
	OrderChannel string `json:"order_channel,omitempty" xml:"order_channel,omitempty"`
	// 每天限制
	PerDay int64 `json:"per_day,omitempty" xml:"per_day,omitempty"`
	// 每单限制
	PerOrder int64 `json:"per_order,omitempty" xml:"per_order,omitempty"`
}

var poolVoucherTemplateSettingOpenInfo = sync.Pool{
	New: func() any {
		return new(VoucherTemplateSettingOpenInfo)
	},
}

// GetVoucherTemplateSettingOpenInfo() 从对象池中获取VoucherTemplateSettingOpenInfo
func GetVoucherTemplateSettingOpenInfo() *VoucherTemplateSettingOpenInfo {
	return poolVoucherTemplateSettingOpenInfo.Get().(*VoucherTemplateSettingOpenInfo)
}

// ReleaseVoucherTemplateSettingOpenInfo 释放VoucherTemplateSettingOpenInfo
func ReleaseVoucherTemplateSettingOpenInfo(v *VoucherTemplateSettingOpenInfo) {
	v.VoucherTemplateIds = v.VoucherTemplateIds[:0]
	v.OrderChannel = ""
	v.PerDay = 0
	v.PerOrder = 0
	poolVoucherTemplateSettingOpenInfo.Put(v)
}

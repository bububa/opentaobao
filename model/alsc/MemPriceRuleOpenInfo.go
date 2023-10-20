package alsc

import (
	"sync"
)

// MemPriceRuleOpenInfo 结构体
type MemPriceRuleOpenInfo struct {
	// 会员等级和特价菜单的关系
	LevelMenuList []LevelMenuOpenInfo `json:"level_menu_list,omitempty" xml:"level_menu_list>level_menu_open_info,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 支付方式限制，UNLIMITED_PAY：无限制，RECHARGE_PAY：储值整单支付
	PayType string `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// 扩展信息
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 是否已删除
	Deleted bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
}

var poolMemPriceRuleOpenInfo = sync.Pool{
	New: func() any {
		return new(MemPriceRuleOpenInfo)
	},
}

// GetMemPriceRuleOpenInfo() 从对象池中获取MemPriceRuleOpenInfo
func GetMemPriceRuleOpenInfo() *MemPriceRuleOpenInfo {
	return poolMemPriceRuleOpenInfo.Get().(*MemPriceRuleOpenInfo)
}

// ReleaseMemPriceRuleOpenInfo 释放MemPriceRuleOpenInfo
func ReleaseMemPriceRuleOpenInfo(v *MemPriceRuleOpenInfo) {
	v.LevelMenuList = v.LevelMenuList[:0]
	v.GmtCreate = ""
	v.GmtModified = ""
	v.PayType = ""
	v.ExtInfo = ""
	v.Deleted = false
	poolMemPriceRuleOpenInfo.Put(v)
}

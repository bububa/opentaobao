package alsc

import (
	"sync"
)

// MenuOpenInfo 结构体
type MenuOpenInfo struct {
	// 菜品集合
	MenuDetailOpenInfoList []MenuDetailOpenInfo `json:"menu_detail_open_info_list,omitempty" xml:"menu_detail_open_info_list>menu_detail_open_info,omitempty"`
	// 生效时段结束
	EffectEnd string `json:"effect_end,omitempty" xml:"effect_end,omitempty"`
	// 生效时段起始
	EffectStart string `json:"effect_start,omitempty" xml:"effect_start,omitempty"`
	// 特价菜单id
	MenuId string `json:"menu_id,omitempty" xml:"menu_id,omitempty"`
	// 特价菜单名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 统一折扣价
	ProDiscount string `json:"pro_discount,omitempty" xml:"pro_discount,omitempty"`
	// 特价模式：统一折扣、不同菜不同折扣
	ProMode string `json:"pro_mode,omitempty" xml:"pro_mode,omitempty"`
	// 更新时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 是否逻辑删除
	Deleted bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
}

var poolMenuOpenInfo = sync.Pool{
	New: func() any {
		return new(MenuOpenInfo)
	},
}

// GetMenuOpenInfo() 从对象池中获取MenuOpenInfo
func GetMenuOpenInfo() *MenuOpenInfo {
	return poolMenuOpenInfo.Get().(*MenuOpenInfo)
}

// ReleaseMenuOpenInfo 释放MenuOpenInfo
func ReleaseMenuOpenInfo(v *MenuOpenInfo) {
	v.MenuDetailOpenInfoList = v.MenuDetailOpenInfoList[:0]
	v.EffectEnd = ""
	v.EffectStart = ""
	v.MenuId = ""
	v.Name = ""
	v.ProDiscount = ""
	v.ProMode = ""
	v.GmtModified = ""
	v.GmtCreate = ""
	v.Deleted = false
	poolMenuOpenInfo.Put(v)
}

package simba

import (
	"sync"
)

// DmpModuleConfigVo 结构体
type DmpModuleConfigVo struct {
	// dmp数据来源选项
	DmpBaseOptionalSelectAppList []DmpBaseOptionalSelectVo `json:"dmp_base_optional_select_app_list,omitempty" xml:"dmp_base_optional_select_app_list>dmp_base_optional_select_vo,omitempty"`
	// 投放渠道id，筛在达摩盘平台上创建且同步到业务线的人群
	DeliverAppId int64 `json:"deliver_app_id,omitempty" xml:"deliver_app_id,omitempty"`
}

var poolDmpModuleConfigVo = sync.Pool{
	New: func() any {
		return new(DmpModuleConfigVo)
	},
}

// GetDmpModuleConfigVo() 从对象池中获取DmpModuleConfigVo
func GetDmpModuleConfigVo() *DmpModuleConfigVo {
	return poolDmpModuleConfigVo.Get().(*DmpModuleConfigVo)
}

// ReleaseDmpModuleConfigVo 释放DmpModuleConfigVo
func ReleaseDmpModuleConfigVo(v *DmpModuleConfigVo) {
	v.DmpBaseOptionalSelectAppList = v.DmpBaseOptionalSelectAppList[:0]
	v.DeliverAppId = 0
	poolDmpModuleConfigVo.Put(v)
}

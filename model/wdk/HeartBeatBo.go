package wdk

import (
	"sync"
)

// HeartBeatBo 结构体
type HeartBeatBo struct {
	// MARKET-营销，ITEM-商品
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 当前版本信息
	VersionId int64 `json:"version_id,omitempty" xml:"version_id,omitempty"`
}

var poolHeartBeatBo = sync.Pool{
	New: func() any {
		return new(HeartBeatBo)
	},
}

// GetHeartBeatBo() 从对象池中获取HeartBeatBo
func GetHeartBeatBo() *HeartBeatBo {
	return poolHeartBeatBo.Get().(*HeartBeatBo)
}

// ReleaseHeartBeatBo 释放HeartBeatBo
func ReleaseHeartBeatBo(v *HeartBeatBo) {
	v.BizCode = ""
	v.VersionId = 0
	poolHeartBeatBo.Put(v)
}

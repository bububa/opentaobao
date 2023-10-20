package travel

import (
	"sync"
)

// ItemResourceInfo 结构体
type ItemResourceInfo struct {
	// 对应的说明
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 1-保险 2-餐饮 3-租车 4-签证 5-购物点，预留  6-赠品，预留 7-券，预留  8-其他费用
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 关联的套餐id
	RelatedPackageId int64 `json:"related_package_id,omitempty" xml:"related_package_id,omitempty"`
}

var poolItemResourceInfo = sync.Pool{
	New: func() any {
		return new(ItemResourceInfo)
	},
}

// GetItemResourceInfo() 从对象池中获取ItemResourceInfo
func GetItemResourceInfo() *ItemResourceInfo {
	return poolItemResourceInfo.Get().(*ItemResourceInfo)
}

// ReleaseItemResourceInfo 释放ItemResourceInfo
func ReleaseItemResourceInfo(v *ItemResourceInfo) {
	v.Desc = ""
	v.Type = 0
	v.RelatedPackageId = 0
	poolItemResourceInfo.Put(v)
}

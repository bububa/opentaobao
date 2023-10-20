package travel

import (
	"sync"
)

// ItemScenicInfo 结构体
type ItemScenicInfo struct {
	// 结构化景点信息 景点结构化信息和文本描述二选一
	ScenicList []ItemScenic `json:"scenic_list,omitempty" xml:"scenic_list>item_scenic,omitempty"`
}

var poolItemScenicInfo = sync.Pool{
	New: func() any {
		return new(ItemScenicInfo)
	},
}

// GetItemScenicInfo() 从对象池中获取ItemScenicInfo
func GetItemScenicInfo() *ItemScenicInfo {
	return poolItemScenicInfo.Get().(*ItemScenicInfo)
}

// ReleaseItemScenicInfo 释放ItemScenicInfo
func ReleaseItemScenicInfo(v *ItemScenicInfo) {
	v.ScenicList = v.ScenicList[:0]
	poolItemScenicInfo.Put(v)
}

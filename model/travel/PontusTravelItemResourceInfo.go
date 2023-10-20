package travel

import (
	"sync"
)

// PontusTravelItemResourceInfo 结构体
type PontusTravelItemResourceInfo struct {
	// 小于1500字符
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 1-保险2-餐饮3-租车4-签证5-购物点6-赠品7-券99-其他
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolPontusTravelItemResourceInfo = sync.Pool{
	New: func() any {
		return new(PontusTravelItemResourceInfo)
	},
}

// GetPontusTravelItemResourceInfo() 从对象池中获取PontusTravelItemResourceInfo
func GetPontusTravelItemResourceInfo() *PontusTravelItemResourceInfo {
	return poolPontusTravelItemResourceInfo.Get().(*PontusTravelItemResourceInfo)
}

// ReleasePontusTravelItemResourceInfo 释放PontusTravelItemResourceInfo
func ReleasePontusTravelItemResourceInfo(v *PontusTravelItemResourceInfo) {
	v.Desc = ""
	v.Type = 0
	poolPontusTravelItemResourceInfo.Put(v)
}

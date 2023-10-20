package idle

import (
	"sync"
)

// TenderItemDeleteCmd 结构体
type TenderItemDeleteCmd struct {
	// 商品id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolTenderItemDeleteCmd = sync.Pool{
	New: func() any {
		return new(TenderItemDeleteCmd)
	},
}

// GetTenderItemDeleteCmd() 从对象池中获取TenderItemDeleteCmd
func GetTenderItemDeleteCmd() *TenderItemDeleteCmd {
	return poolTenderItemDeleteCmd.Get().(*TenderItemDeleteCmd)
}

// ReleaseTenderItemDeleteCmd 释放TenderItemDeleteCmd
func ReleaseTenderItemDeleteCmd(v *TenderItemDeleteCmd) {
	v.ItemId = ""
	poolTenderItemDeleteCmd.Put(v)
}

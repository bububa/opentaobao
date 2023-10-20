package perfect

import (
	"sync"
)

// TcStoreList 结构体
type TcStoreList struct {
	// 门店名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 门店id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}

var poolTcStoreList = sync.Pool{
	New: func() any {
		return new(TcStoreList)
	},
}

// GetTcStoreList() 从对象池中获取TcStoreList
func GetTcStoreList() *TcStoreList {
	return poolTcStoreList.Get().(*TcStoreList)
}

// ReleaseTcStoreList 释放TcStoreList
func ReleaseTcStoreList(v *TcStoreList) {
	v.StoreName = ""
	v.StoreId = 0
	poolTcStoreList.Put(v)
}

package drug

import (
	"sync"
)

// Cat 结构体
type Cat struct {
	// catId
	CatId string `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
	// catName
	CatName string `json:"cat_name,omitempty" xml:"cat_name,omitempty"`
	// itemCount
	ItemCount int64 `json:"item_count,omitempty" xml:"item_count,omitempty"`
}

var poolCat = sync.Pool{
	New: func() any {
		return new(Cat)
	},
}

// GetCat() 从对象池中获取Cat
func GetCat() *Cat {
	return poolCat.Get().(*Cat)
}

// ReleaseCat 释放Cat
func ReleaseCat(v *Cat) {
	v.CatId = ""
	v.CatName = ""
	v.ItemCount = 0
	poolCat.Put(v)
}

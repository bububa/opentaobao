package drug

import (
	"sync"
)

// StoreDetailDto 结构体
type StoreDetailDto struct {
	// tags
	Tags []Tags `json:"tags,omitempty" xml:"tags>tags,omitempty"`
	// cat
	Cats []Cat `json:"cats,omitempty" xml:"cats>cat,omitempty"`
	// storeDetail
	StoreDetail *StoreDto `json:"store_detail,omitempty" xml:"store_detail,omitempty"`
}

var poolStoreDetailDto = sync.Pool{
	New: func() any {
		return new(StoreDetailDto)
	},
}

// GetStoreDetailDto() 从对象池中获取StoreDetailDto
func GetStoreDetailDto() *StoreDetailDto {
	return poolStoreDetailDto.Get().(*StoreDetailDto)
}

// ReleaseStoreDetailDto 释放StoreDetailDto
func ReleaseStoreDetailDto(v *StoreDetailDto) {
	v.Tags = v.Tags[:0]
	v.Cats = v.Cats[:0]
	v.StoreDetail = nil
	poolStoreDetailDto.Put(v)
}

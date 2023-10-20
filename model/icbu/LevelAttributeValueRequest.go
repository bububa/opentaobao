package icbu

import (
	"sync"
)

// LevelAttributeValueRequest 结构体
type LevelAttributeValueRequest struct {
	// 类目属性id，放到数组第一个位置
	AttrId []string `json:"attr_id,omitempty" xml:"attr_id>string,omitempty"`
	// 属性值id, 不同取值范围时的查询策略如下:  &lt;=0：列出当前类目属性的所有属性值  &gt;0：指定当前类目属性的某一个属性值，列出该属性值下的子属性和该子属性的所有属性值
	ValueId int64 `json:"value_id,omitempty" xml:"value_id,omitempty"`
	// 必填；要查询的属性值所属发布类目
	CatId int64 `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
}

var poolLevelAttributeValueRequest = sync.Pool{
	New: func() any {
		return new(LevelAttributeValueRequest)
	},
}

// GetLevelAttributeValueRequest() 从对象池中获取LevelAttributeValueRequest
func GetLevelAttributeValueRequest() *LevelAttributeValueRequest {
	return poolLevelAttributeValueRequest.Get().(*LevelAttributeValueRequest)
}

// ReleaseLevelAttributeValueRequest 释放LevelAttributeValueRequest
func ReleaseLevelAttributeValueRequest(v *LevelAttributeValueRequest) {
	v.AttrId = v.AttrId[:0]
	v.ValueId = 0
	v.CatId = 0
	poolLevelAttributeValueRequest.Put(v)
}

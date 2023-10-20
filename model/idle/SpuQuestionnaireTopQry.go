package idle

import (
	"sync"
)

// SpuQuestionnaireTopQry 结构体
type SpuQuestionnaireTopQry struct {
	// 投放业务，RECYCLE_3C（回收），RECYCLE_TENDER（寄拍）
	BizType string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// SPU ID
	SpuId int64 `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
}

var poolSpuQuestionnaireTopQry = sync.Pool{
	New: func() any {
		return new(SpuQuestionnaireTopQry)
	},
}

// GetSpuQuestionnaireTopQry() 从对象池中获取SpuQuestionnaireTopQry
func GetSpuQuestionnaireTopQry() *SpuQuestionnaireTopQry {
	return poolSpuQuestionnaireTopQry.Get().(*SpuQuestionnaireTopQry)
}

// ReleaseSpuQuestionnaireTopQry 释放SpuQuestionnaireTopQry
func ReleaseSpuQuestionnaireTopQry(v *SpuQuestionnaireTopQry) {
	v.BizType = ""
	v.SpuId = 0
	poolSpuQuestionnaireTopQry.Put(v)
}

package idle

import (
	"sync"
)

// QuestionnaireSupportResult 结构体
type QuestionnaireSupportResult struct {
	// 投放业务
	BizType string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 预上线版本号
	PreVersion string `json:"pre_version,omitempty" xml:"pre_version,omitempty"`
	// spu id
	SpuId int64 `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
}

var poolQuestionnaireSupportResult = sync.Pool{
	New: func() any {
		return new(QuestionnaireSupportResult)
	},
}

// GetQuestionnaireSupportResult() 从对象池中获取QuestionnaireSupportResult
func GetQuestionnaireSupportResult() *QuestionnaireSupportResult {
	return poolQuestionnaireSupportResult.Get().(*QuestionnaireSupportResult)
}

// ReleaseQuestionnaireSupportResult 释放QuestionnaireSupportResult
func ReleaseQuestionnaireSupportResult(v *QuestionnaireSupportResult) {
	v.BizType = ""
	v.PreVersion = ""
	v.SpuId = 0
	poolQuestionnaireSupportResult.Put(v)
}

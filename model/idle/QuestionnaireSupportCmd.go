package idle

import (
	"sync"
)

// QuestionnaireSupportCmd 结构体
type QuestionnaireSupportCmd struct {
	// 投放业务，RECYCLE_3C（回收），RECYCLE_TENDER（寄拍）
	BizType string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 预上线版本号
	PreVersion string `json:"pre_version,omitempty" xml:"pre_version,omitempty"`
	// spu id
	SpuId int64 `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
}

var poolQuestionnaireSupportCmd = sync.Pool{
	New: func() any {
		return new(QuestionnaireSupportCmd)
	},
}

// GetQuestionnaireSupportCmd() 从对象池中获取QuestionnaireSupportCmd
func GetQuestionnaireSupportCmd() *QuestionnaireSupportCmd {
	return poolQuestionnaireSupportCmd.Get().(*QuestionnaireSupportCmd)
}

// ReleaseQuestionnaireSupportCmd 释放QuestionnaireSupportCmd
func ReleaseQuestionnaireSupportCmd(v *QuestionnaireSupportCmd) {
	v.BizType = ""
	v.PreVersion = ""
	v.SpuId = 0
	poolQuestionnaireSupportCmd.Put(v)
}

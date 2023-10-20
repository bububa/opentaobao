package usergrowth

import (
	"sync"
)

// BatchAskResultItem 结构体
type BatchAskResultItem struct {
	// 在巨浪平台可投放的任务ID列表
	TaskIdList []string `json:"task_id_list,omitempty" xml:"task_id_list>string,omitempty"`
	// oaid的md5值， 32位小写
	OaidMd5 string `json:"oaid_md5,omitempty" xml:"oaid_md5,omitempty"`
	// idfa的md5值， 32位小写
	IdfaMd5 string `json:"idfa_md5,omitempty" xml:"idfa_md5,omitempty"`
	// imei的md5值， 32位小写
	ImeiMd5 string `json:"imei_md5,omitempty" xml:"imei_md5,omitempty"`
	// 该设备要做的大航海的任务id
	TaskId string `json:"task_id,omitempty" xml:"task_id,omitempty"`
	// caid的md5值， 32位小写，前面拼接上caid版本号，当前支持20220111、20211207版本
	CaidMd5 string `json:"caid_md5,omitempty" xml:"caid_md5,omitempty"`
}

var poolBatchAskResultItem = sync.Pool{
	New: func() any {
		return new(BatchAskResultItem)
	},
}

// GetBatchAskResultItem() 从对象池中获取BatchAskResultItem
func GetBatchAskResultItem() *BatchAskResultItem {
	return poolBatchAskResultItem.Get().(*BatchAskResultItem)
}

// ReleaseBatchAskResultItem 释放BatchAskResultItem
func ReleaseBatchAskResultItem(v *BatchAskResultItem) {
	v.TaskIdList = v.TaskIdList[:0]
	v.OaidMd5 = ""
	v.IdfaMd5 = ""
	v.ImeiMd5 = ""
	v.TaskId = ""
	v.CaidMd5 = ""
	poolBatchAskResultItem.Put(v)
}

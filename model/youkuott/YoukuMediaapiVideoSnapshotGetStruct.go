package youkuott

import (
	"sync"
)

// YoukuMediaapiVideoSnapshotGetStruct 结构体
type YoukuMediaapiVideoSnapshotGetStruct struct {
	// 图片url列表
	ThumbIdList []string `json:"thumb_id_list,omitempty" xml:"thumb_id_list>string,omitempty"`
	// 图片域名
	DomainName string `json:"domain_name,omitempty" xml:"domain_name,omitempty"`
	// 毫秒
	Sectiontime int64 `json:"sectiontime,omitempty" xml:"sectiontime,omitempty"`
}

var poolYoukuMediaapiVideoSnapshotGetStruct = sync.Pool{
	New: func() any {
		return new(YoukuMediaapiVideoSnapshotGetStruct)
	},
}

// GetYoukuMediaapiVideoSnapshotGetStruct() 从对象池中获取YoukuMediaapiVideoSnapshotGetStruct
func GetYoukuMediaapiVideoSnapshotGetStruct() *YoukuMediaapiVideoSnapshotGetStruct {
	return poolYoukuMediaapiVideoSnapshotGetStruct.Get().(*YoukuMediaapiVideoSnapshotGetStruct)
}

// ReleaseYoukuMediaapiVideoSnapshotGetStruct 释放YoukuMediaapiVideoSnapshotGetStruct
func ReleaseYoukuMediaapiVideoSnapshotGetStruct(v *YoukuMediaapiVideoSnapshotGetStruct) {
	v.ThumbIdList = v.ThumbIdList[:0]
	v.DomainName = ""
	v.Sectiontime = 0
	poolYoukuMediaapiVideoSnapshotGetStruct.Put(v)
}

package wenyuvideo

import (
	"sync"
)

// PersonRbo 结构体
type PersonRbo struct {
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 人物头像
	ThumbUrl string `json:"thumb_url,omitempty" xml:"thumb_url,omitempty"`
	// 300*300人物头像
	ThumbUrlLg string `json:"thumb_url_lg,omitempty" xml:"thumb_url_lg,omitempty"`
	// 人物海报
	PosterUrl string `json:"poster_url,omitempty" xml:"poster_url,omitempty"`
	// 演职人员工作
	Job string `json:"job,omitempty" xml:"job,omitempty"`
	// taotv媒资的演职人员id
	PersonId int64 `json:"person_id,omitempty" xml:"person_id,omitempty"`
}

var poolPersonRbo = sync.Pool{
	New: func() any {
		return new(PersonRbo)
	},
}

// GetPersonRbo() 从对象池中获取PersonRbo
func GetPersonRbo() *PersonRbo {
	return poolPersonRbo.Get().(*PersonRbo)
}

// ReleasePersonRbo 释放PersonRbo
func ReleasePersonRbo(v *PersonRbo) {
	v.Name = ""
	v.ThumbUrl = ""
	v.ThumbUrlLg = ""
	v.PosterUrl = ""
	v.Job = ""
	v.PersonId = 0
	poolPersonRbo.Put(v)
}

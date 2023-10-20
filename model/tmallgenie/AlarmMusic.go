package tmallgenie

import (
	"sync"
)

// AlarmMusic 结构体
type AlarmMusic struct {
	// 铃声类型
	Category string `json:"category,omitempty" xml:"category,omitempty"`
	// 铃声id
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 铃声名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 铃声url
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 铃声来源，默认为虾米
	Source string `json:"source,omitempty" xml:"source,omitempty"`
}

var poolAlarmMusic = sync.Pool{
	New: func() any {
		return new(AlarmMusic)
	},
}

// GetAlarmMusic() 从对象池中获取AlarmMusic
func GetAlarmMusic() *AlarmMusic {
	return poolAlarmMusic.Get().(*AlarmMusic)
}

// ReleaseAlarmMusic 释放AlarmMusic
func ReleaseAlarmMusic(v *AlarmMusic) {
	v.Category = ""
	v.Id = ""
	v.Name = ""
	v.Url = ""
	v.Source = ""
	poolAlarmMusic.Put(v)
}

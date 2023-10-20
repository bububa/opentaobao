package viapi

import (
	"sync"
)

// ProgramCodeData 结构体
type ProgramCodeData struct {
	// 以图片左上角为坐标原点，小程序码区域左上角到y轴距离
	X int64 `json:"x,omitempty" xml:"x,omitempty"`
	// 以图片左上角为坐标原点，小程序码区域左上角到x轴距离
	Y int64 `json:"y,omitempty" xml:"y,omitempty"`
	// 小程序码区域宽度
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
	// 小程序码区域高度
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
}

var poolProgramCodeData = sync.Pool{
	New: func() any {
		return new(ProgramCodeData)
	},
}

// GetProgramCodeData() 从对象池中获取ProgramCodeData
func GetProgramCodeData() *ProgramCodeData {
	return poolProgramCodeData.Get().(*ProgramCodeData)
}

// ReleaseProgramCodeData 释放ProgramCodeData
func ReleaseProgramCodeData(v *ProgramCodeData) {
	v.X = 0
	v.Y = 0
	v.Width = 0
	v.Height = 0
	poolProgramCodeData.Put(v)
}

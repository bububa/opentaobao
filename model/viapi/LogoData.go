package viapi

import (
	"sync"
)

// LogoData 结构体
type LogoData struct {
	// 识别出的logo类型，取值为TV （台标）
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 识别出的logo名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 以图片左上角为坐标原点，logo区域左上角到y轴距离。
	X int64 `json:"x,omitempty" xml:"x,omitempty"`
	// 以图片左上角为坐标原点，logo区域左上角到x轴距离。
	Y int64 `json:"y,omitempty" xml:"y,omitempty"`
	// logo区域宽度
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
	// logo区域高度
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
}

var poolLogoData = sync.Pool{
	New: func() any {
		return new(LogoData)
	},
}

// GetLogoData() 从对象池中获取LogoData
func GetLogoData() *LogoData {
	return poolLogoData.Get().(*LogoData)
}

// ReleaseLogoData 释放LogoData
func ReleaseLogoData(v *LogoData) {
	v.Type = ""
	v.Name = ""
	v.X = 0
	v.Y = 0
	v.Width = 0
	v.Height = 0
	poolLogoData.Put(v)
}

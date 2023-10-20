package xiamicontent

import (
	"sync"
)

// ArtistDto 结构体
type ArtistDto struct {
	// 地区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 性别
	Gender string `json:"gender,omitempty" xml:"gender,omitempty"`
	// 别名
	Alias string `json:"alias,omitempty" xml:"alias,omitempty"`
	// 艺人名
	ArtistName string `json:"artist_name,omitempty" xml:"artist_name,omitempty"`
	// 艺人封面
	ArtistLogo string `json:"artist_logo,omitempty" xml:"artist_logo,omitempty"`
	// 艺人id
	ArtistId int64 `json:"artist_id,omitempty" xml:"artist_id,omitempty"`
}

var poolArtistDto = sync.Pool{
	New: func() any {
		return new(ArtistDto)
	},
}

// GetArtistDto() 从对象池中获取ArtistDto
func GetArtistDto() *ArtistDto {
	return poolArtistDto.Get().(*ArtistDto)
}

// ReleaseArtistDto 释放ArtistDto
func ReleaseArtistDto(v *ArtistDto) {
	v.Area = ""
	v.Gender = ""
	v.Alias = ""
	v.ArtistName = ""
	v.ArtistLogo = ""
	v.ArtistId = 0
	poolArtistDto.Put(v)
}

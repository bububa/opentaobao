package xiamicontent

import (
	"sync"
)

// SongCatTagDto 结构体
type SongCatTagDto struct {
	// tag code列表
	TagCodes []string `json:"tag_codes,omitempty" xml:"tag_codes>string,omitempty"`
	// tag间的关系查询 1 and 2：or
	Relation int64 `json:"relation,omitempty" xml:"relation,omitempty"`
}

var poolSongCatTagDto = sync.Pool{
	New: func() any {
		return new(SongCatTagDto)
	},
}

// GetSongCatTagDto() 从对象池中获取SongCatTagDto
func GetSongCatTagDto() *SongCatTagDto {
	return poolSongCatTagDto.Get().(*SongCatTagDto)
}

// ReleaseSongCatTagDto 释放SongCatTagDto
func ReleaseSongCatTagDto(v *SongCatTagDto) {
	v.TagCodes = v.TagCodes[:0]
	v.Relation = 0
	poolSongCatTagDto.Put(v)
}

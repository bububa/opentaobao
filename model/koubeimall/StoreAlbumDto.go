package koubeimall

import (
	"sync"
)

// StoreAlbumDto 结构体
type StoreAlbumDto struct {
	// 照片列表
	PictureList []Picture `json:"picture_list,omitempty" xml:"picture_list>picture,omitempty"`
	// 相册名称
	AlbumName string `json:"album_name,omitempty" xml:"album_name,omitempty"`
	// 描述信息
	AlbumDesc string `json:"album_desc,omitempty" xml:"album_desc,omitempty"`
}

var poolStoreAlbumDto = sync.Pool{
	New: func() any {
		return new(StoreAlbumDto)
	},
}

// GetStoreAlbumDto() 从对象池中获取StoreAlbumDto
func GetStoreAlbumDto() *StoreAlbumDto {
	return poolStoreAlbumDto.Get().(*StoreAlbumDto)
}

// ReleaseStoreAlbumDto 释放StoreAlbumDto
func ReleaseStoreAlbumDto(v *StoreAlbumDto) {
	v.PictureList = v.PictureList[:0]
	v.AlbumName = ""
	v.AlbumDesc = ""
	poolStoreAlbumDto.Put(v)
}

package koubeimall

// StoreAlbumDto 结构体
type StoreAlbumDto struct {
	// 相册名称
	AlbumName string `json:"album_name,omitempty" xml:"album_name,omitempty"`
	// 描述信息
	AlbumDesc string `json:"album_desc,omitempty" xml:"album_desc,omitempty"`
	// 照片列表
	PictureList []Picture `json:"picture_list,omitempty" xml:"picture_list>picture,omitempty"`
}

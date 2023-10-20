package wdkitem

import (
	"sync"
)

// PictureDo 结构体
type PictureDo struct {
	// 返回的是相对路劲
	PicturePath string `json:"picture_path,omitempty" xml:"picture_path,omitempty"`
	// 图片标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 图片状态,unfroze代表没有被冻结，froze代表被冻结,pass代表排查通过
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 图片是否删除的标记
	Deleted string `json:"deleted,omitempty" xml:"deleted,omitempty"`
	// 图片的创建时间
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 图片的修改时间
	Modified string `json:"modified,omitempty" xml:"modified,omitempty"`
	// 图片相素,格式:长x宽，如450x150
	Pixel string `json:"pixel,omitempty" xml:"pixel,omitempty"`
	// 完整的路劲
	FullUrl string `json:"full_url,omitempty" xml:"full_url,omitempty"`
	// 图片ID
	PictureId int64 `json:"picture_id,omitempty" xml:"picture_id,omitempty"`
	// 图片分类ID
	PictureCategoryId int64 `json:"picture_category_id,omitempty" xml:"picture_category_id,omitempty"`
	// 图片大小,bite单位
	Sizes int64 `json:"sizes,omitempty" xml:"sizes,omitempty"`
}

var poolPictureDo = sync.Pool{
	New: func() any {
		return new(PictureDo)
	},
}

// GetPictureDo() 从对象池中获取PictureDo
func GetPictureDo() *PictureDo {
	return poolPictureDo.Get().(*PictureDo)
}

// ReleasePictureDo 释放PictureDo
func ReleasePictureDo(v *PictureDo) {
	v.PicturePath = ""
	v.Title = ""
	v.Status = ""
	v.Deleted = ""
	v.Created = ""
	v.Modified = ""
	v.Pixel = ""
	v.FullUrl = ""
	v.PictureId = 0
	v.PictureCategoryId = 0
	v.Sizes = 0
	poolPictureDo.Put(v)
}

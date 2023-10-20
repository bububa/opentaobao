package media

import (
	"sync"
)

// Picture 结构体
type Picture struct {
	// 返回的是绝对路径如：http://img07.taobaocdn.com/imgextra/i7/22670458/T2dD0kXb4cXXXXXXXX_!!22670458.jpg
	PicturePath string `json:"picture_path,omitempty" xml:"picture_path,omitempty"`
	// 图片标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 图片相素,格式:长x宽，如450x150
	Pixel string `json:"pixel,omitempty" xml:"pixel,omitempty"`
	// 图片状态,0 未审核没冻结 1  冻结 2 审核通过
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 图片是否删除的标记
	Deleted string `json:"deleted,omitempty" xml:"deleted,omitempty"`
	// 图片上传的来源，有电脑版本宝贝发布，手机版本宝贝发布
	ClientType string `json:"client_type,omitempty" xml:"client_type,omitempty"`
	// 图片的创建时间
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 图片的修改时间
	Modified string `json:"modified,omitempty" xml:"modified,omitempty"`
	// 图片在后台处理之后的md5值当md5值为32位长度的字符串时为图片搬家后的文件md5验证码md5值为长整数时为图片替换后的时间戳
	Md5 string `json:"md5,omitempty" xml:"md5,omitempty"`
	// 图片ID
	PictureId int64 `json:"picture_id,omitempty" xml:"picture_id,omitempty"`
	// 图片分类ID
	PictureCategoryId int64 `json:"picture_category_id,omitempty" xml:"picture_category_id,omitempty"`
	// 图片大小,bite单位
	Sizes int64 `json:"sizes,omitempty" xml:"sizes,omitempty"`
	// 图片是否被引用
	Referenced bool `json:"referenced,omitempty" xml:"referenced,omitempty"`
}

var poolPicture = sync.Pool{
	New: func() any {
		return new(Picture)
	},
}

// GetPicture() 从对象池中获取Picture
func GetPicture() *Picture {
	return poolPicture.Get().(*Picture)
}

// ReleasePicture 释放Picture
func ReleasePicture(v *Picture) {
	v.PicturePath = ""
	v.Title = ""
	v.Pixel = ""
	v.Status = ""
	v.Deleted = ""
	v.ClientType = ""
	v.Created = ""
	v.Modified = ""
	v.Md5 = ""
	v.PictureId = 0
	v.PictureCategoryId = 0
	v.Sizes = 0
	v.Referenced = false
	poolPicture.Put(v)
}

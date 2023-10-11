package util

import (
	"github.com/bububa/opentaobao/model"
)

// UploadRequest 结构体
type UploadRequest struct {
	// 图片url，image_data和image_url二选一
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	// 123.jpg
	ImageTitle string `json:"image_title,omitempty" xml:"image_title,omitempty"`
	// 用户id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// tu业务线图片上传文件夹id，填0上传到根目录
	DirIdForTu int64 `json:"dir_id_for_tu,omitempty" xml:"dir_id_for_tu,omitempty"`
	// qnaigc业务线图片上传文件夹id，填0上传到根目录
	DirIdForQnaigc int64 `json:"dir_id_for_qnaigc,omitempty" xml:"dir_id_for_qnaigc,omitempty"`
	// 图片字节列表，image_data和image_url二选一
	ImageData *model.File `json:"image_data,omitempty" xml:"image_data,omitempty"`
}

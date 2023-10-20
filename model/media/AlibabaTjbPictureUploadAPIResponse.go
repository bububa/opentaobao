package media

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabatjbpictureuploadAPIResponse 淘特图片空间上传单张图片 API返回值
// alibaba.tjb.picture.upload
//
// 淘特图片空间上传单张图片
type AlibabatjbpictureuploadAPIResponse struct {
	model.CommonResponse
	AlibabatjbpictureuploadAPIResponseModel
}

// AlibabatjbpictureuploadAPIResponseModel is 淘特图片空间上传单张图片 成功返回结果
type AlibabatjbpictureuploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tjb_picture_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 图片信息
	Picture *TopPictureDto `json:"picture,omitempty" xml:"picture,omitempty"`
}

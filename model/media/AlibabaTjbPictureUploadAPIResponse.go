package media

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTjbPictureUploadAPIResponse 淘特图片空间上传单张图片 API返回值
// alibaba.tjb.picture.upload
//
// 淘特图片空间上传单张图片
type AlibabaTjbPictureUploadAPIResponse struct {
	model.CommonResponse
	AlibabaTjbPictureUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTjbPictureUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTjbPictureUploadAPIResponseModel).Reset()
}

// AlibabaTjbPictureUploadAPIResponseModel is 淘特图片空间上传单张图片 成功返回结果
type AlibabaTjbPictureUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tjb_picture_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 图片信息
	Picture *TopPictureDto `json:"picture,omitempty" xml:"picture,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTjbPictureUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Picture = nil
}

var poolAlibabaTjbPictureUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTjbPictureUploadAPIResponse)
	},
}

// GetAlibabaTjbPictureUploadAPIResponse 从 sync.Pool 获取 AlibabaTjbPictureUploadAPIResponse
func GetAlibabaTjbPictureUploadAPIResponse() *AlibabaTjbPictureUploadAPIResponse {
	return poolAlibabaTjbPictureUploadAPIResponse.Get().(*AlibabaTjbPictureUploadAPIResponse)
}

// ReleaseAlibabaTjbPictureUploadAPIResponse 将 AlibabaTjbPictureUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTjbPictureUploadAPIResponse(v *AlibabaTjbPictureUploadAPIResponse) {
	v.Reset()
	poolAlibabaTjbPictureUploadAPIResponse.Put(v)
}

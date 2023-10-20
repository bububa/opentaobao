package media

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPictureUploadAPIResponse 上传单张图片 API返回值
// taobao.picture.upload
//
// 图片空间上传接口
type TaobaoPictureUploadAPIResponse struct {
	model.CommonResponse
	TaobaoPictureUploadAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPictureUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPictureUploadAPIResponseModel).Reset()
}

// TaobaoPictureUploadAPIResponseModel is 上传单张图片 成功返回结果
type TaobaoPictureUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"picture_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 当前上传的一张图片信息
	Picture *Picture `json:"picture,omitempty" xml:"picture,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPictureUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Picture = nil
}

var poolTaobaoPictureUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPictureUploadAPIResponse)
	},
}

// GetTaobaoPictureUploadAPIResponse 从 sync.Pool 获取 TaobaoPictureUploadAPIResponse
func GetTaobaoPictureUploadAPIResponse() *TaobaoPictureUploadAPIResponse {
	return poolTaobaoPictureUploadAPIResponse.Get().(*TaobaoPictureUploadAPIResponse)
}

// ReleaseTaobaoPictureUploadAPIResponse 将 TaobaoPictureUploadAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPictureUploadAPIResponse(v *TaobaoPictureUploadAPIResponse) {
	v.Reset()
	poolTaobaoPictureUploadAPIResponse.Put(v)
}

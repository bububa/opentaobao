package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRentMediaUploadAPIResponse 闲鱼多媒体上传接口 API返回值
// alibaba.idle.rent.media.upload
//
// 上传多媒体信息，包括图片、视频（暂不支持）
type AlibabaIdleRentMediaUploadAPIResponse struct {
	model.CommonResponse
	AlibabaIdleRentMediaUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleRentMediaUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleRentMediaUploadAPIResponseModel).Reset()
}

// AlibabaIdleRentMediaUploadAPIResponseModel is 闲鱼多媒体上传接口 成功返回结果
type AlibabaIdleRentMediaUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_rent_media_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应数据
	Result *AlibabaIdleRentMediaUploadTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleRentMediaUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleRentMediaUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleRentMediaUploadAPIResponse)
	},
}

// GetAlibabaIdleRentMediaUploadAPIResponse 从 sync.Pool 获取 AlibabaIdleRentMediaUploadAPIResponse
func GetAlibabaIdleRentMediaUploadAPIResponse() *AlibabaIdleRentMediaUploadAPIResponse {
	return poolAlibabaIdleRentMediaUploadAPIResponse.Get().(*AlibabaIdleRentMediaUploadAPIResponse)
}

// ReleaseAlibabaIdleRentMediaUploadAPIResponse 将 AlibabaIdleRentMediaUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleRentMediaUploadAPIResponse(v *AlibabaIdleRentMediaUploadAPIResponse) {
	v.Reset()
	poolAlibabaIdleRentMediaUploadAPIResponse.Put(v)
}

package idleisv

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvMediaUploadAPIResponse 闲鱼服务商-图片上传 API返回值
// alibaba.idle.isv.media.upload
//
// 供外部服务商ISV进行闲鱼商品发布时上传商品所需图片
type AlibabaIdleIsvMediaUploadAPIResponse struct {
	model.CommonResponse
	AlibabaIdleIsvMediaUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleIsvMediaUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleIsvMediaUploadAPIResponseModel).Reset()
}

// AlibabaIdleIsvMediaUploadAPIResponseModel is 闲鱼服务商-图片上传 成功返回结果
type AlibabaIdleIsvMediaUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_isv_media_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应数据
	Result *AlibabaIdleIsvMediaUploadTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleIsvMediaUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleIsvMediaUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleIsvMediaUploadAPIResponse)
	},
}

// GetAlibabaIdleIsvMediaUploadAPIResponse 从 sync.Pool 获取 AlibabaIdleIsvMediaUploadAPIResponse
func GetAlibabaIdleIsvMediaUploadAPIResponse() *AlibabaIdleIsvMediaUploadAPIResponse {
	return poolAlibabaIdleIsvMediaUploadAPIResponse.Get().(*AlibabaIdleIsvMediaUploadAPIResponse)
}

// ReleaseAlibabaIdleIsvMediaUploadAPIResponse 将 AlibabaIdleIsvMediaUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleIsvMediaUploadAPIResponse(v *AlibabaIdleIsvMediaUploadAPIResponse) {
	v.Reset()
	poolAlibabaIdleIsvMediaUploadAPIResponse.Put(v)
}

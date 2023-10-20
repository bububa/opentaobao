package idleitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleItemMediaAddAPIResponse 图片上传 API返回值
// alibaba.idle.item.media.add
//
// 上传图片
type AlibabaIdleItemMediaAddAPIResponse struct {
	model.CommonResponse
	AlibabaIdleItemMediaAddAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleItemMediaAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleItemMediaAddAPIResponseModel).Reset()
}

// AlibabaIdleItemMediaAddAPIResponseModel is 图片上传 成功返回结果
type AlibabaIdleItemMediaAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_item_media_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *EasyResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleItemMediaAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleItemMediaAddAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleItemMediaAddAPIResponse)
	},
}

// GetAlibabaIdleItemMediaAddAPIResponse 从 sync.Pool 获取 AlibabaIdleItemMediaAddAPIResponse
func GetAlibabaIdleItemMediaAddAPIResponse() *AlibabaIdleItemMediaAddAPIResponse {
	return poolAlibabaIdleItemMediaAddAPIResponse.Get().(*AlibabaIdleItemMediaAddAPIResponse)
}

// ReleaseAlibabaIdleItemMediaAddAPIResponse 将 AlibabaIdleItemMediaAddAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleItemMediaAddAPIResponse(v *AlibabaIdleItemMediaAddAPIResponse) {
	v.Reset()
	poolAlibabaIdleItemMediaAddAPIResponse.Put(v)
}

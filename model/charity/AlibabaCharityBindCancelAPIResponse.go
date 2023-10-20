package charity

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCharityBindCancelAPIResponse 取消用户绑定 API返回值
// alibaba.charity.bind.cancel
//
// 取消用户绑定
type AlibabaCharityBindCancelAPIResponse struct {
	model.CommonResponse
	AlibabaCharityBindCancelAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCharityBindCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCharityBindCancelAPIResponseModel).Reset()
}

// AlibabaCharityBindCancelAPIResponseModel is 取消用户绑定 成功返回结果
type AlibabaCharityBindCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_charity_bind_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *ThreehoursResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCharityBindCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCharityBindCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCharityBindCancelAPIResponse)
	},
}

// GetAlibabaCharityBindCancelAPIResponse 从 sync.Pool 获取 AlibabaCharityBindCancelAPIResponse
func GetAlibabaCharityBindCancelAPIResponse() *AlibabaCharityBindCancelAPIResponse {
	return poolAlibabaCharityBindCancelAPIResponse.Get().(*AlibabaCharityBindCancelAPIResponse)
}

// ReleaseAlibabaCharityBindCancelAPIResponse 将 AlibabaCharityBindCancelAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCharityBindCancelAPIResponse(v *AlibabaCharityBindCancelAPIResponse) {
	v.Reset()
	poolAlibabaCharityBindCancelAPIResponse.Put(v)
}

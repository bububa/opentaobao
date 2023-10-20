package charity

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCharityCharitytimeCommonauthAPIResponse 通用授权 API返回值
// alibaba.charity.charitytime.commonauth
//
// 三小时和外部账号绑定通用top 返回跳转链接进行绑定
type AlibabaCharityCharitytimeCommonauthAPIResponse struct {
	model.CommonResponse
	AlibabaCharityCharitytimeCommonauthAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCharityCharitytimeCommonauthAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCharityCharitytimeCommonauthAPIResponseModel).Reset()
}

// AlibabaCharityCharitytimeCommonauthAPIResponseModel is 通用授权 成功返回结果
type AlibabaCharityCharitytimeCommonauthAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_charity_charitytime_commonauth_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *CsrResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCharityCharitytimeCommonauthAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCharityCharitytimeCommonauthAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCharityCharitytimeCommonauthAPIResponse)
	},
}

// GetAlibabaCharityCharitytimeCommonauthAPIResponse 从 sync.Pool 获取 AlibabaCharityCharitytimeCommonauthAPIResponse
func GetAlibabaCharityCharitytimeCommonauthAPIResponse() *AlibabaCharityCharitytimeCommonauthAPIResponse {
	return poolAlibabaCharityCharitytimeCommonauthAPIResponse.Get().(*AlibabaCharityCharitytimeCommonauthAPIResponse)
}

// ReleaseAlibabaCharityCharitytimeCommonauthAPIResponse 将 AlibabaCharityCharitytimeCommonauthAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCharityCharitytimeCommonauthAPIResponse(v *AlibabaCharityCharitytimeCommonauthAPIResponse) {
	v.Reset()
	poolAlibabaCharityCharitytimeCommonauthAPIResponse.Put(v)
}

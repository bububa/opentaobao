package middleclaims

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMiddleClaimsbillReceiveAPIResponse 国际化中台服务域接收理赔账单 API返回值
// alibaba.middle.claimsbill.receive
//
// 国际化中台服务域与保险公司交互对接一个订单在保险公司方对该订单进行理赔打款的处理后，将该打款结果返回至服务域
type AlibabaMiddleClaimsbillReceiveAPIResponse struct {
	model.CommonResponse
	AlibabaMiddleClaimsbillReceiveAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMiddleClaimsbillReceiveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMiddleClaimsbillReceiveAPIResponseModel).Reset()
}

// AlibabaMiddleClaimsbillReceiveAPIResponseModel is 国际化中台服务域接收理赔账单 成功返回结果
type AlibabaMiddleClaimsbillReceiveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_middle_claimsbill_receive_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果实体类
	Result *AlibabaMiddleClaimsbillReceiveResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMiddleClaimsbillReceiveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMiddleClaimsbillReceiveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMiddleClaimsbillReceiveAPIResponse)
	},
}

// GetAlibabaMiddleClaimsbillReceiveAPIResponse 从 sync.Pool 获取 AlibabaMiddleClaimsbillReceiveAPIResponse
func GetAlibabaMiddleClaimsbillReceiveAPIResponse() *AlibabaMiddleClaimsbillReceiveAPIResponse {
	return poolAlibabaMiddleClaimsbillReceiveAPIResponse.Get().(*AlibabaMiddleClaimsbillReceiveAPIResponse)
}

// ReleaseAlibabaMiddleClaimsbillReceiveAPIResponse 将 AlibabaMiddleClaimsbillReceiveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMiddleClaimsbillReceiveAPIResponse(v *AlibabaMiddleClaimsbillReceiveAPIResponse) {
	v.Reset()
	poolAlibabaMiddleClaimsbillReceiveAPIResponse.Put(v)
}

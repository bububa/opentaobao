package middleclaims

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMiddleClaimsresultReceiveAPIResponse 国际化中台服务域接收理赔结果 API返回值
// alibaba.middle.claimsresult.receive
//
// 国际化中台服务域与保险公司交互对接一个订单在保险公司方对该订单进行理赔结果的处理后，将该结果返回至服务域
type AlibabaMiddleClaimsresultReceiveAPIResponse struct {
	model.CommonResponse
	AlibabaMiddleClaimsresultReceiveAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMiddleClaimsresultReceiveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMiddleClaimsresultReceiveAPIResponseModel).Reset()
}

// AlibabaMiddleClaimsresultReceiveAPIResponseModel is 国际化中台服务域接收理赔结果 成功返回结果
type AlibabaMiddleClaimsresultReceiveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_middle_claimsresult_receive_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *AlibabaMiddleClaimsresultReceiveResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMiddleClaimsresultReceiveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMiddleClaimsresultReceiveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMiddleClaimsresultReceiveAPIResponse)
	},
}

// GetAlibabaMiddleClaimsresultReceiveAPIResponse 从 sync.Pool 获取 AlibabaMiddleClaimsresultReceiveAPIResponse
func GetAlibabaMiddleClaimsresultReceiveAPIResponse() *AlibabaMiddleClaimsresultReceiveAPIResponse {
	return poolAlibabaMiddleClaimsresultReceiveAPIResponse.Get().(*AlibabaMiddleClaimsresultReceiveAPIResponse)
}

// ReleaseAlibabaMiddleClaimsresultReceiveAPIResponse 将 AlibabaMiddleClaimsresultReceiveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMiddleClaimsresultReceiveAPIResponse(v *AlibabaMiddleClaimsresultReceiveAPIResponse) {
	v.Reset()
	poolAlibabaMiddleClaimsresultReceiveAPIResponse.Put(v)
}

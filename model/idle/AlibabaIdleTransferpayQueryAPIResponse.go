package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleTransferpayQueryAPIResponse 闲鱼转账结果查询 API返回值
// alibaba.idle.transferpay.query
//
// 商家业务 转账支付的结果查询
type AlibabaIdleTransferpayQueryAPIResponse struct {
	model.CommonResponse
	AlibabaIdleTransferpayQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleTransferpayQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleTransferpayQueryAPIResponseModel).Reset()
}

// AlibabaIdleTransferpayQueryAPIResponseModel is 闲鱼转账结果查询 成功返回结果
type AlibabaIdleTransferpayQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_transferpay_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaIdleTransferpayQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleTransferpayQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleTransferpayQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleTransferpayQueryAPIResponse)
	},
}

// GetAlibabaIdleTransferpayQueryAPIResponse 从 sync.Pool 获取 AlibabaIdleTransferpayQueryAPIResponse
func GetAlibabaIdleTransferpayQueryAPIResponse() *AlibabaIdleTransferpayQueryAPIResponse {
	return poolAlibabaIdleTransferpayQueryAPIResponse.Get().(*AlibabaIdleTransferpayQueryAPIResponse)
}

// ReleaseAlibabaIdleTransferpayQueryAPIResponse 将 AlibabaIdleTransferpayQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleTransferpayQueryAPIResponse(v *AlibabaIdleTransferpayQueryAPIResponse) {
	v.Reset()
	poolAlibabaIdleTransferpayQueryAPIResponse.Put(v)
}

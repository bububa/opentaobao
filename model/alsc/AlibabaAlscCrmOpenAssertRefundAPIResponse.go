package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmOpenAssertRefundAPIResponse 资产核销回退接口 API返回值
// alibaba.alsc.crm.open.assert.refund
//
// 回退已经核销的储值积分券资产
type AlibabaAlscCrmOpenAssertRefundAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmOpenAssertRefundAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmOpenAssertRefundAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmOpenAssertRefundAPIResponseModel).Reset()
}

// AlibabaAlscCrmOpenAssertRefundAPIResponseModel is 资产核销回退接口 成功返回结果
type AlibabaAlscCrmOpenAssertRefundAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_open_assert_refund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmOpenAssertRefundAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmOpenAssertRefundAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmOpenAssertRefundAPIResponse)
	},
}

// GetAlibabaAlscCrmOpenAssertRefundAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmOpenAssertRefundAPIResponse
func GetAlibabaAlscCrmOpenAssertRefundAPIResponse() *AlibabaAlscCrmOpenAssertRefundAPIResponse {
	return poolAlibabaAlscCrmOpenAssertRefundAPIResponse.Get().(*AlibabaAlscCrmOpenAssertRefundAPIResponse)
}

// ReleaseAlibabaAlscCrmOpenAssertRefundAPIResponse 将 AlibabaAlscCrmOpenAssertRefundAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmOpenAssertRefundAPIResponse(v *AlibabaAlscCrmOpenAssertRefundAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmOpenAssertRefundAPIResponse.Put(v)
}

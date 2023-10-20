package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleFengniaoTradeUpdateAPIResponse 更新蜂鸟扣费状态 API返回值
// alibaba.ele.fengniao.trade.update
//
// 汇金扣费成功后，回调该接口更新扣费状态
type AlibabaEleFengniaoTradeUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaEleFengniaoTradeUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEleFengniaoTradeUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEleFengniaoTradeUpdateAPIResponseModel).Reset()
}

// AlibabaEleFengniaoTradeUpdateAPIResponseModel is 更新蜂鸟扣费状态 成功返回结果
type AlibabaEleFengniaoTradeUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_fengniao_trade_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 无此交易记录
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 1:成功 0：失败
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEleFengniaoTradeUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.Status = 0
}

var poolAlibabaEleFengniaoTradeUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEleFengniaoTradeUpdateAPIResponse)
	},
}

// GetAlibabaEleFengniaoTradeUpdateAPIResponse 从 sync.Pool 获取 AlibabaEleFengniaoTradeUpdateAPIResponse
func GetAlibabaEleFengniaoTradeUpdateAPIResponse() *AlibabaEleFengniaoTradeUpdateAPIResponse {
	return poolAlibabaEleFengniaoTradeUpdateAPIResponse.Get().(*AlibabaEleFengniaoTradeUpdateAPIResponse)
}

// ReleaseAlibabaEleFengniaoTradeUpdateAPIResponse 将 AlibabaEleFengniaoTradeUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEleFengniaoTradeUpdateAPIResponse(v *AlibabaEleFengniaoTradeUpdateAPIResponse) {
	v.Reset()
	poolAlibabaEleFengniaoTradeUpdateAPIResponse.Put(v)
}

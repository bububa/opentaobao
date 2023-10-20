package nlife

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNlifeB2cTradeCancelAPIResponse 零售+平台取消订单 API返回值
// alibaba.nlife.b2c.trade.cancel
//
// 零售+平台取消订单接口
type AlibabaNlifeB2cTradeCancelAPIResponse struct {
	model.CommonResponse
	AlibabaNlifeB2cTradeCancelAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaNlifeB2cTradeCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaNlifeB2cTradeCancelAPIResponseModel).Reset()
}

// AlibabaNlifeB2cTradeCancelAPIResponseModel is 零售+平台取消订单 成功返回结果
type AlibabaNlifeB2cTradeCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_nlife_b2c_trade_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单取消时间
	GmtCancel string `json:"gmt_cancel,omitempty" xml:"gmt_cancel,omitempty"`
	// 扩展参数JSON
	ExtendParams string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaNlifeB2cTradeCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.GmtCancel = ""
	m.ExtendParams = ""
}

var poolAlibabaNlifeB2cTradeCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaNlifeB2cTradeCancelAPIResponse)
	},
}

// GetAlibabaNlifeB2cTradeCancelAPIResponse 从 sync.Pool 获取 AlibabaNlifeB2cTradeCancelAPIResponse
func GetAlibabaNlifeB2cTradeCancelAPIResponse() *AlibabaNlifeB2cTradeCancelAPIResponse {
	return poolAlibabaNlifeB2cTradeCancelAPIResponse.Get().(*AlibabaNlifeB2cTradeCancelAPIResponse)
}

// ReleaseAlibabaNlifeB2cTradeCancelAPIResponse 将 AlibabaNlifeB2cTradeCancelAPIResponse 保存到 sync.Pool
func ReleaseAlibabaNlifeB2cTradeCancelAPIResponse(v *AlibabaNlifeB2cTradeCancelAPIResponse) {
	v.Reset()
	poolAlibabaNlifeB2cTradeCancelAPIResponse.Put(v)
}

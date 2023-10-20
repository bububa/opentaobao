package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRdcAligeniusRefundsCheckAPIResponse 退款信息审核 API返回值
// taobao.rdc.aligenius.refunds.check
//
// 根据退款信息，对退款单进行审核
type TaobaoRdcAligeniusRefundsCheckAPIResponse struct {
	model.CommonResponse
	TaobaoRdcAligeniusRefundsCheckAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRdcAligeniusRefundsCheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRdcAligeniusRefundsCheckAPIResponseModel).Reset()
}

// TaobaoRdcAligeniusRefundsCheckAPIResponseModel is 退款信息审核 成功返回结果
type TaobaoRdcAligeniusRefundsCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"rdc_aligenius_refunds_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoRdcAligeniusRefundsCheckResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRdcAligeniusRefundsCheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoRdcAligeniusRefundsCheckAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRdcAligeniusRefundsCheckAPIResponse)
	},
}

// GetTaobaoRdcAligeniusRefundsCheckAPIResponse 从 sync.Pool 获取 TaobaoRdcAligeniusRefundsCheckAPIResponse
func GetTaobaoRdcAligeniusRefundsCheckAPIResponse() *TaobaoRdcAligeniusRefundsCheckAPIResponse {
	return poolTaobaoRdcAligeniusRefundsCheckAPIResponse.Get().(*TaobaoRdcAligeniusRefundsCheckAPIResponse)
}

// ReleaseTaobaoRdcAligeniusRefundsCheckAPIResponse 将 TaobaoRdcAligeniusRefundsCheckAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRdcAligeniusRefundsCheckAPIResponse(v *TaobaoRdcAligeniusRefundsCheckAPIResponse) {
	v.Reset()
	poolTaobaoRdcAligeniusRefundsCheckAPIResponse.Put(v)
}

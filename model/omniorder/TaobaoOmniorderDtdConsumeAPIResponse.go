package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderDtdConsumeAPIResponse 门店自送对码进行核销 API返回值
// taobao.omniorder.dtd.consume
//
// 该接口根据传入的码及订单信息，如果码与订单一致，则对门店自送服务进行核销。
type TaobaoOmniorderDtdConsumeAPIResponse struct {
	model.CommonResponse
	TaobaoOmniorderDtdConsumeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOmniorderDtdConsumeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOmniorderDtdConsumeAPIResponseModel).Reset()
}

// TaobaoOmniorderDtdConsumeAPIResponseModel is 门店自送对码进行核销 成功返回结果
type TaobaoOmniorderDtdConsumeAPIResponseModel struct {
	XMLName xml.Name `xml:"omniorder_dtd_consume_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码，为0表示成功，非0表示失败
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误西溪
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOmniorderDtdConsumeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.Message = ""
}

var poolTaobaoOmniorderDtdConsumeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOmniorderDtdConsumeAPIResponse)
	},
}

// GetTaobaoOmniorderDtdConsumeAPIResponse 从 sync.Pool 获取 TaobaoOmniorderDtdConsumeAPIResponse
func GetTaobaoOmniorderDtdConsumeAPIResponse() *TaobaoOmniorderDtdConsumeAPIResponse {
	return poolTaobaoOmniorderDtdConsumeAPIResponse.Get().(*TaobaoOmniorderDtdConsumeAPIResponse)
}

// ReleaseTaobaoOmniorderDtdConsumeAPIResponse 将 TaobaoOmniorderDtdConsumeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOmniorderDtdConsumeAPIResponse(v *TaobaoOmniorderDtdConsumeAPIResponse) {
	v.Reset()
	poolTaobaoOmniorderDtdConsumeAPIResponse.Put(v)
}

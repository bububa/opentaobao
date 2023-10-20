package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStorecollectConsumeAPIResponse 全渠道门店自提核销订单 API返回值
// taobao.omniorder.storecollect.consume
//
// 全渠道门店自提核销订单
type TaobaoOmniorderStorecollectConsumeAPIResponse struct {
	model.CommonResponse
	TaobaoOmniorderStorecollectConsumeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOmniorderStorecollectConsumeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOmniorderStorecollectConsumeAPIResponseModel).Reset()
}

// TaobaoOmniorderStorecollectConsumeAPIResponseModel is 全渠道门店自提核销订单 成功返回结果
type TaobaoOmniorderStorecollectConsumeAPIResponseModel struct {
	XMLName xml.Name `xml:"omniorder_storecollect_consume_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 0表示成功
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 核销错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOmniorderStorecollectConsumeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrCode = ""
	m.ErrMsg = ""
}

var poolTaobaoOmniorderStorecollectConsumeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOmniorderStorecollectConsumeAPIResponse)
	},
}

// GetTaobaoOmniorderStorecollectConsumeAPIResponse 从 sync.Pool 获取 TaobaoOmniorderStorecollectConsumeAPIResponse
func GetTaobaoOmniorderStorecollectConsumeAPIResponse() *TaobaoOmniorderStorecollectConsumeAPIResponse {
	return poolTaobaoOmniorderStorecollectConsumeAPIResponse.Get().(*TaobaoOmniorderStorecollectConsumeAPIResponse)
}

// ReleaseTaobaoOmniorderStorecollectConsumeAPIResponse 将 TaobaoOmniorderStorecollectConsumeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOmniorderStorecollectConsumeAPIResponse(v *TaobaoOmniorderStorecollectConsumeAPIResponse) {
	v.Reset()
	poolTaobaoOmniorderStorecollectConsumeAPIResponse.Put(v)
}

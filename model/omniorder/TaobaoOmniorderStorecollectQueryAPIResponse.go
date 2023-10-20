package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStorecollectQueryAPIResponse 全渠道门店自提根据核销码查询订单 API返回值
// taobao.omniorder.storecollect.query
//
// 全渠道门店自提根据核销码查询订单
type TaobaoOmniorderStorecollectQueryAPIResponse struct {
	model.CommonResponse
	TaobaoOmniorderStorecollectQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOmniorderStorecollectQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOmniorderStorecollectQueryAPIResponseModel).Reset()
}

// TaobaoOmniorderStorecollectQueryAPIResponseModel is 全渠道门店自提根据核销码查询订单 成功返回结果
type TaobaoOmniorderStorecollectQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"omniorder_storecollect_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoOmniorderStorecollectQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOmniorderStorecollectQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOmniorderStorecollectQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOmniorderStorecollectQueryAPIResponse)
	},
}

// GetTaobaoOmniorderStorecollectQueryAPIResponse 从 sync.Pool 获取 TaobaoOmniorderStorecollectQueryAPIResponse
func GetTaobaoOmniorderStorecollectQueryAPIResponse() *TaobaoOmniorderStorecollectQueryAPIResponse {
	return poolTaobaoOmniorderStorecollectQueryAPIResponse.Get().(*TaobaoOmniorderStorecollectQueryAPIResponse)
}

// ReleaseTaobaoOmniorderStorecollectQueryAPIResponse 将 TaobaoOmniorderStorecollectQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOmniorderStorecollectQueryAPIResponse(v *TaobaoOmniorderStorecollectQueryAPIResponse) {
	v.Reset()
	poolTaobaoOmniorderStorecollectQueryAPIResponse.Put(v)
}

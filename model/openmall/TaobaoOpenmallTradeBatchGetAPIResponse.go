package openmall

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallTradeBatchGetAPIResponse 批量获取openmall订单 API返回值
// taobao.openmall.trade.batch.get
//
// 批量获取openmall订单
// 注意：该接口数据存在延迟，实时数据请通过taobao.openmall.trade.get获取
type TaobaoOpenmallTradeBatchGetAPIResponse struct {
	model.CommonResponse
	TaobaoOpenmallTradeBatchGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenmallTradeBatchGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenmallTradeBatchGetAPIResponseModel).Reset()
}

// TaobaoOpenmallTradeBatchGetAPIResponseModel is 批量获取openmall订单 成功返回结果
type TaobaoOpenmallTradeBatchGetAPIResponseModel struct {
	XMLName xml.Name `xml:"openmall_trade_batch_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单列表
	Entities []TopTradeDetailVo `json:"entities,omitempty" xml:"entities>top_trade_detail_vo,omitempty"`
	// 范围内总订单数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenmallTradeBatchGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Entities = m.Entities[:0]
	m.TotalCount = 0
}

var poolTaobaoOpenmallTradeBatchGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenmallTradeBatchGetAPIResponse)
	},
}

// GetTaobaoOpenmallTradeBatchGetAPIResponse 从 sync.Pool 获取 TaobaoOpenmallTradeBatchGetAPIResponse
func GetTaobaoOpenmallTradeBatchGetAPIResponse() *TaobaoOpenmallTradeBatchGetAPIResponse {
	return poolTaobaoOpenmallTradeBatchGetAPIResponse.Get().(*TaobaoOpenmallTradeBatchGetAPIResponse)
}

// ReleaseTaobaoOpenmallTradeBatchGetAPIResponse 将 TaobaoOpenmallTradeBatchGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenmallTradeBatchGetAPIResponse(v *TaobaoOpenmallTradeBatchGetAPIResponse) {
	v.Reset()
	poolTaobaoOpenmallTradeBatchGetAPIResponse.Put(v)
}

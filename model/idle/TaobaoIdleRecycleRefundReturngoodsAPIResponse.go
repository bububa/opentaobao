package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoIdleRecycleRefundReturngoodsAPIResponse 闲鱼回收退货V2 API返回值
// taobao.idle.recycle.refund.returngoods
//
// 回收商买家退货，填写退货运单号
type TaobaoIdleRecycleRefundReturngoodsAPIResponse struct {
	model.CommonResponse
	TaobaoIdleRecycleRefundReturngoodsAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoIdleRecycleRefundReturngoodsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoIdleRecycleRefundReturngoodsAPIResponseModel).Reset()
}

// TaobaoIdleRecycleRefundReturngoodsAPIResponseModel is 闲鱼回收退货V2 成功返回结果
type TaobaoIdleRecycleRefundReturngoodsAPIResponseModel struct {
	XMLName xml.Name `xml:"idle_recycle_refund_returngoods_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 退货
	Result *IdleTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoIdleRecycleRefundReturngoodsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoIdleRecycleRefundReturngoodsAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoIdleRecycleRefundReturngoodsAPIResponse)
	},
}

// GetTaobaoIdleRecycleRefundReturngoodsAPIResponse 从 sync.Pool 获取 TaobaoIdleRecycleRefundReturngoodsAPIResponse
func GetTaobaoIdleRecycleRefundReturngoodsAPIResponse() *TaobaoIdleRecycleRefundReturngoodsAPIResponse {
	return poolTaobaoIdleRecycleRefundReturngoodsAPIResponse.Get().(*TaobaoIdleRecycleRefundReturngoodsAPIResponse)
}

// ReleaseTaobaoIdleRecycleRefundReturngoodsAPIResponse 将 TaobaoIdleRecycleRefundReturngoodsAPIResponse 保存到 sync.Pool
func ReleaseTaobaoIdleRecycleRefundReturngoodsAPIResponse(v *TaobaoIdleRecycleRefundReturngoodsAPIResponse) {
	v.Reset()
	poolTaobaoIdleRecycleRefundReturngoodsAPIResponse.Put(v)
}

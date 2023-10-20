package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoIdleRecycleRefundCancleapplyAPIResponse 闲鱼回收取消退款申请V2 API返回值
// taobao.idle.recycle.refund.cancleapply
//
// 回收商的回收订单取消退款申请
type TaobaoIdleRecycleRefundCancleapplyAPIResponse struct {
	model.CommonResponse
	TaobaoIdleRecycleRefundCancleapplyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoIdleRecycleRefundCancleapplyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoIdleRecycleRefundCancleapplyAPIResponseModel).Reset()
}

// TaobaoIdleRecycleRefundCancleapplyAPIResponseModel is 闲鱼回收取消退款申请V2 成功返回结果
type TaobaoIdleRecycleRefundCancleapplyAPIResponseModel struct {
	XMLName xml.Name `xml:"idle_recycle_refund_cancleapply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 撤销申请结果
	Result *IdleTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoIdleRecycleRefundCancleapplyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoIdleRecycleRefundCancleapplyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoIdleRecycleRefundCancleapplyAPIResponse)
	},
}

// GetTaobaoIdleRecycleRefundCancleapplyAPIResponse 从 sync.Pool 获取 TaobaoIdleRecycleRefundCancleapplyAPIResponse
func GetTaobaoIdleRecycleRefundCancleapplyAPIResponse() *TaobaoIdleRecycleRefundCancleapplyAPIResponse {
	return poolTaobaoIdleRecycleRefundCancleapplyAPIResponse.Get().(*TaobaoIdleRecycleRefundCancleapplyAPIResponse)
}

// ReleaseTaobaoIdleRecycleRefundCancleapplyAPIResponse 将 TaobaoIdleRecycleRefundCancleapplyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoIdleRecycleRefundCancleapplyAPIResponse(v *TaobaoIdleRecycleRefundCancleapplyAPIResponse) {
	v.Reset()
	poolTaobaoIdleRecycleRefundCancleapplyAPIResponse.Put(v)
}

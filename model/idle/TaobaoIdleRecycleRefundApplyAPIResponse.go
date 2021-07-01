package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoIdleRecycleRefundApplyAPIResponse
闲鱼回收交易退款申请V2 API返回值
taobao.idle.recycle.refund.apply

回收商买家申请退款 */
type TaobaoIdleRecycleRefundApplyAPIResponse struct {
	model.CommonResponse
	TaobaoIdleRecycleRefundApplyAPIResponseModel
}

// TaobaoIdleRecycleRefundApplyAPIResponseModel is 闲鱼回收交易退款申请V2 成功返回结果
type TaobaoIdleRecycleRefundApplyAPIResponseModel struct {
	XMLName xml.Name `xml:"idle_recycle_refund_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 退款申请结果
	Result *IdleTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

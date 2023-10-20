package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoidlerecyclerefundcancleapplyAPIResponse 闲鱼回收取消退款申请V2 API返回值
// taobao.idle.recycle.refund.cancleapply
//
// 回收商的回收订单取消退款申请
type TaobaoidlerecyclerefundcancleapplyAPIResponse struct {
	model.CommonResponse
	TaobaoidlerecyclerefundcancleapplyAPIResponseModel
}

// TaobaoidlerecyclerefundcancleapplyAPIResponseModel is 闲鱼回收取消退款申请V2 成功返回结果
type TaobaoidlerecyclerefundcancleapplyAPIResponseModel struct {
	XMLName xml.Name `xml:"idle_recycle_refund_cancleapply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 撤销申请结果
	Result *IdleTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

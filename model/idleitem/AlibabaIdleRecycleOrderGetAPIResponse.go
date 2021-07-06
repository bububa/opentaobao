package idleitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRecycleOrderGetAPIResponse 闲鱼回收订单查询V2 API返回值
// alibaba.idle.recycle.order.get
//
// 闲鱼回收业务中,外部回收商作为交易上买家,闲鱼用户下单后,需要回收商主动拉取交易订单
type AlibabaIdleRecycleOrderGetAPIResponse struct {
	model.CommonResponse
	AlibabaIdleRecycleOrderGetAPIResponseModel
}

// AlibabaIdleRecycleOrderGetAPIResponseModel is 闲鱼回收订单查询V2 成功返回结果
type AlibabaIdleRecycleOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_recycle_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 失败时候错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 订单信息
	Module *RecycleOrderTo `json:"module,omitempty" xml:"module,omitempty"`
	// 是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLogisticsOnlineCancelAPIResponse
取消物流订单接口 API返回值
taobao.logistics.online.cancel

调此接口取消发货的订单，重新选择物流公司发货。前提是物流公司未揽收货物。对未发货和已经被物流公司揽收的物流订单，是不能取消的。 */
type TaobaoLogisticsOnlineCancelAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsOnlineCancelAPIResponseModel
}

// TaobaoLogisticsOnlineCancelAPIResponseModel is 取消物流订单接口 成功返回结果
type TaobaoLogisticsOnlineCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_online_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功与失败
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 修改时间
	ModifyTime string `json:"modify_time,omitempty" xml:"modify_time,omitempty"`
	// 重新创建物流订单id
	RecreatedOrderId string `json:"recreated_order_id,omitempty" xml:"recreated_order_id,omitempty"`
}

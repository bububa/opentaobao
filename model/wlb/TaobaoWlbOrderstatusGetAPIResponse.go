package wlb

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbOrderstatusGetAPIResponse
物流宝订单流转状态查询 API返回值
taobao.wlb.orderstatus.get

根据物流宝订单号查询物流宝订单至目前为止的流转状态列表 */
type TaobaoWlbOrderstatusGetAPIResponse struct {
	model.CommonResponse
	TaobaoWlbOrderstatusGetAPIResponseModel
}

// TaobaoWlbOrderstatusGetAPIResponseModel is 物流宝订单流转状态查询 成功返回结果
type TaobaoWlbOrderstatusGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_orderstatus_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单流转信息状态列表
	WlbOrderStatus []WlbProcessStatus `json:"wlb_order_status,omitempty" xml:"wlb_order_status>wlb_process_status,omitempty"`
}

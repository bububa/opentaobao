package wlb

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbOrderdetailDateGetAPIResponse 按照日期范围查询物流订单详情 API返回值
// taobao.wlb.orderdetail.date.get
//
// 外部ERP可通过该接口查询一段时间内的物流宝订单，以及订单详情
type TaobaoWlbOrderdetailDateGetAPIResponse struct {
	model.CommonResponse
	TaobaoWlbOrderdetailDateGetAPIResponseModel
}

// TaobaoWlbOrderdetailDateGetAPIResponseModel is 按照日期范围查询物流订单详情 成功返回结果
type TaobaoWlbOrderdetailDateGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_orderdetail_date_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 物流宝订单，并且包含订单详情
	OrderDetailList []WlbOrderDetail `json:"order_detail_list,omitempty" xml:"order_detail_list>wlb_order_detail,omitempty"`
	// 总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

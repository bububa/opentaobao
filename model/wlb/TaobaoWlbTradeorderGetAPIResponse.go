package wlb

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbtradeordergetAPIResponse 根据交易号获取物流宝订单 API返回值
// taobao.wlb.tradeorder.get
//
// 根据交易类型和交易id查询物流宝订单详情
type TaobaowlbtradeordergetAPIResponse struct {
	model.CommonResponse
	TaobaowlbtradeordergetAPIResponseModel
}

// TaobaowlbtradeordergetAPIResponseModel is 根据交易号获取物流宝订单 成功返回结果
type TaobaowlbtradeordergetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_tradeorder_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 物流宝订单对象
	WlbOrderList []WlbOrder `json:"wlb_order_list,omitempty" xml:"wlb_order_list>wlb_order,omitempty"`
}

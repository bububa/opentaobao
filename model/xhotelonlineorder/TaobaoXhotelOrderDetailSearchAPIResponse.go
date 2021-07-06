package xhotelonlineorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderDetailSearchAPIResponse 订单详情查询 API返回值
// taobao.xhotel.order.detail.search
//
// 提供订单详情查询
type TaobaoXhotelOrderDetailSearchAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelOrderDetailSearchAPIResponseModel
}

// TaobaoXhotelOrderDetailSearchAPIResponseModel is 订单详情查询 成功返回结果
type TaobaoXhotelOrderDetailSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_order_detail_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误编号
	Error string `json:"error,omitempty" xml:"error,omitempty"`
	// 订单详情对象
	TopOrderDetail *TopOrderDetail `json:"top_order_detail,omitempty" xml:"top_order_detail,omitempty"`
}

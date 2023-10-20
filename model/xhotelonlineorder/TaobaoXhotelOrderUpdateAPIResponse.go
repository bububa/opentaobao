package xhotelonlineorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelorderupdateAPIResponse 酒店订单发货接口 API返回值
// taobao.xhotel.order.update
//
// 卖家确认订单或者取消订单，适用于预付、面付、信用住订单
type TaobaoxhotelorderupdateAPIResponse struct {
	model.CommonResponse
	TaobaoxhotelorderupdateAPIResponseModel
}

// TaobaoxhotelorderupdateAPIResponseModel is 酒店订单发货接口 成功返回结果
type TaobaoxhotelorderupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_order_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回提示信息
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

package xhotelonlineorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelorderalipayfaceupdateAPIResponse 酒店信用住订单状态更新 API返回值
// taobao.xhotel.order.alipayface.update
//
// 完成对信用住或者面付订单的状态的更新。包含订单状态的确认，入离店状态的更新等等。（不适用于预付订单）
type TaobaoxhotelorderalipayfaceupdateAPIResponse struct {
	model.CommonResponse
	TaobaoxhotelorderalipayfaceupdateAPIResponseModel
}

// TaobaoxhotelorderalipayfaceupdateAPIResponseModel is 酒店信用住订单状态更新 成功返回结果
type TaobaoxhotelorderalipayfaceupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_order_alipayface_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回提示信息
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

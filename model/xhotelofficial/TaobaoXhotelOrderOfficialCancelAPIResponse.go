package xhotelofficial

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderOfficialCancelAPIResponse 官网信用住订单取消 API返回值
// taobao.xhotel.order.official.cancel
//
// 官网信用住订单取消
type TaobaoXhotelOrderOfficialCancelAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelOrderOfficialCancelAPIResponseModel
}

// TaobaoXhotelOrderOfficialCancelAPIResponseModel is 官网信用住订单取消 成功返回结果
type TaobaoXhotelOrderOfficialCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_order_official_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回提示信息
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

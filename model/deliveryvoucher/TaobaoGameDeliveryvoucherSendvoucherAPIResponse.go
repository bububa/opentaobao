package deliveryvoucher

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoGameDeliveryvoucherSendvoucherAPIResponse 提货券发券接口 API返回值
// taobao.game.deliveryvoucher.sendvoucher
//
// 提货券发券接口：同步券和订单的关联信息
type TaobaoGameDeliveryvoucherSendvoucherAPIResponse struct {
	model.CommonResponse
	TaobaoGameDeliveryvoucherSendvoucherAPIResponseModel
}

// TaobaoGameDeliveryvoucherSendvoucherAPIResponseModel is 提货券发券接口 成功返回结果
type TaobaoGameDeliveryvoucherSendvoucherAPIResponseModel struct {
	XMLName xml.Name `xml:"game_deliveryvoucher_sendvoucher_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

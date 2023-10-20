package util

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimentradeuserdeleteAPIResponse 删除奇门订单链路用户 API返回值
// taobao.qimen.trade.user.delete
//
// 删除奇门订单链路用户
type TaobaoqimentradeuserdeleteAPIResponse struct {
	model.CommonResponse
	TaobaoqimentradeuserdeleteAPIResponseModel
}

// TaobaoqimentradeuserdeleteAPIResponseModel is 删除奇门订单链路用户 成功返回结果
type TaobaoqimentradeuserdeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_trade_user_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// modal
	Modal bool `json:"modal,omitempty" xml:"modal,omitempty"`
}

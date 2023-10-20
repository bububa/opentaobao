package rhino

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaorhinosupplychainoutboundpickingcompleteAPIResponse 【WMS005】接收成衣捡配完成通知 API返回值
// taobao.rhino.supplychain.outbound.pickingcomplete
//
// 接收成衣捡配完成通知,WMS005
type TaobaorhinosupplychainoutboundpickingcompleteAPIResponse struct {
	model.CommonResponse
	TaobaorhinosupplychainoutboundpickingcompleteAPIResponseModel
}

// TaobaorhinosupplychainoutboundpickingcompleteAPIResponseModel is 【WMS005】接收成衣捡配完成通知 成功返回结果
type TaobaorhinosupplychainoutboundpickingcompleteAPIResponseModel struct {
	XMLName xml.Name `xml:"rhino_supplychain_outbound_pickingcomplete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

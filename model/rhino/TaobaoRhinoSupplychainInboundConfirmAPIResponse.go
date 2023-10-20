package rhino

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaorhinosupplychaininboundconfirmAPIResponse WMS003成衣入库确认 API返回值
// taobao.rhino.supplychain.inbound.confirm
//
// 【WMS003】【同步成衣入库完成信息】
type TaobaorhinosupplychaininboundconfirmAPIResponse struct {
	model.CommonResponse
	TaobaorhinosupplychaininboundconfirmAPIResponseModel
}

// TaobaorhinosupplychaininboundconfirmAPIResponseModel is WMS003成衣入库确认 成功返回结果
type TaobaorhinosupplychaininboundconfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"rhino_supplychain_inbound_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

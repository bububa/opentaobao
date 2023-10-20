package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenreturnorderconfirmAPIResponse 退货入库单确认接口 API返回值
// taobao.qimen.returnorder.confirm
//
// taobao.qimen.returnorder.confirm
type TaobaoqimenreturnorderconfirmAPIResponse struct {
	model.CommonResponse
	TaobaoqimenreturnorderconfirmAPIResponseModel
}

// TaobaoqimenreturnorderconfirmAPIResponseModel is 退货入库单确认接口 成功返回结果
type TaobaoqimenreturnorderconfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_returnorder_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoqimenreturnorderconfirmResponse `json:"response,omitempty" xml:"response,omitempty"`
}

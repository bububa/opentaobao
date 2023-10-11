package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenReturnorderConfirmAPIResponse 退货入库单确认接口 API返回值
// taobao.qimen.returnorder.confirm
//
// taobao.qimen.returnorder.confirm
type TaobaoQimenReturnorderConfirmAPIResponse struct {
	model.CommonResponse
	TaobaoQimenReturnorderConfirmAPIResponseModel
}

// TaobaoQimenReturnorderConfirmAPIResponseModel is 退货入库单确认接口 成功返回结果
type TaobaoQimenReturnorderConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_returnorder_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenReturnorderConfirmResponse `json:"response,omitempty" xml:"response,omitempty"`
}

package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenentryorderconfirmAPIResponse 入库单确认接口 API返回值
// taobao.qimen.entryorder.confirm
//
// WMS调用接口，回传入库单信息;
type TaobaoqimenentryorderconfirmAPIResponse struct {
	model.CommonResponse
	TaobaoqimenentryorderconfirmAPIResponseModel
}

// TaobaoqimenentryorderconfirmAPIResponseModel is 入库单确认接口 成功返回结果
type TaobaoqimenentryorderconfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_entryorder_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoqimenentryorderconfirmResponse `json:"response,omitempty" xml:"response,omitempty"`
}

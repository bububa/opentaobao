package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenReturnorderCreateAPIResponse 退货入库单创建接口 API返回值
// taobao.qimen.returnorder.create
//
// taobao.qimen.returnorder.create
type TaobaoQimenReturnorderCreateAPIResponse struct {
	model.CommonResponse
	TaobaoQimenReturnorderCreateAPIResponseModel
}

// TaobaoQimenReturnorderCreateAPIResponseModel is 退货入库单创建接口 成功返回结果
type TaobaoQimenReturnorderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_returnorder_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenReturnorderCreateResponse `json:"response,omitempty" xml:"response,omitempty"`
}

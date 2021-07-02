package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenTransferorderQueryAPIResponse 调拨单查询 API返回值
// taobao.qimen.transferorder.query
//
// 调拨单查询
type TaobaoQimenTransferorderQueryAPIResponse struct {
	model.CommonResponse
	TaobaoQimenTransferorderQueryAPIResponseModel
}

// TaobaoQimenTransferorderQueryAPIResponseModel is 调拨单查询 成功返回结果
type TaobaoQimenTransferorderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_transferorder_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenTransferorderQueryStruct `json:"response,omitempty" xml:"response,omitempty"`
}

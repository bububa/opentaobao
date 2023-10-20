package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimentransferorderqueryAPIResponse 调拨单查询 API返回值
// taobao.qimen.transferorder.query
//
// 调拨单查询
type TaobaoqimentransferorderqueryAPIResponse struct {
	model.CommonResponse
	TaobaoqimentransferorderqueryAPIResponseModel
}

// TaobaoqimentransferorderqueryAPIResponseModel is 调拨单查询 成功返回结果
type TaobaoqimentransferorderqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_transferorder_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoqimentransferorderqueryStruct `json:"response,omitempty" xml:"response,omitempty"`
}

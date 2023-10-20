package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenentryorderqueryAPIResponse 入库单查询接口 API返回值
// taobao.qimen.entryorder.query
//
// ERP调用接口，查询入库单信息;
type TaobaoqimenentryorderqueryAPIResponse struct {
	model.CommonResponse
	TaobaoqimenentryorderqueryAPIResponseModel
}

// TaobaoqimenentryorderqueryAPIResponseModel is 入库单查询接口 成功返回结果
type TaobaoqimenentryorderqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_entryorder_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *EntryOrderQueryResponse `json:"response,omitempty" xml:"response,omitempty"`
}

package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenchannelinventoryqueryAPIResponse 渠道库存查询接口 API返回值
// taobao.qimen.channelinventory.query
//
// 渠道库存查询
type TaobaoqimenchannelinventoryqueryAPIResponse struct {
	model.CommonResponse
	TaobaoqimenchannelinventoryqueryAPIResponseModel
}

// TaobaoqimenchannelinventoryqueryAPIResponseModel is 渠道库存查询接口 成功返回结果
type TaobaoqimenchannelinventoryqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_channelinventory_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *ResponseDo `json:"response,omitempty" xml:"response,omitempty"`
}

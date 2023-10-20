package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelXitemQueryAPIResponse 查询 x 元素 API返回值
// taobao.xhotel.xitem.query
//
// 查询 x 元素
type TaobaoXhotelXitemQueryAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelXitemQueryAPIResponseModel
}

// TaobaoXhotelXitemQueryAPIResponseModel is 查询 x 元素 成功返回结果
type TaobaoXhotelXitemQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_xitem_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoXhotelXitemQueryResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

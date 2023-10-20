package travel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelitemelementqueryAPIResponse 【API3.0】资源元素查询接口 API返回值
// taobao.alitrip.travel.item.element.query
//
// 提供资源元素查询接口，支持商家查询已经发布过的资源元素
type TaobaoalitriptravelitemelementqueryAPIResponse struct {
	model.CommonResponse
	TaobaoalitriptravelitemelementqueryAPIResponseModel
}

// TaobaoalitriptravelitemelementqueryAPIResponseModel is 【API3.0】资源元素查询接口 成功返回结果
type TaobaoalitriptravelitemelementqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_item_element_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 资源元素列表
	Results []TopElementParam `json:"results,omitempty" xml:"results>top_element_param,omitempty"`
}

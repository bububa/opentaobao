package opentrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeGroupQueryAPIResponse 组团购场景查询商品组团详情 API返回值
// taobao.opentrade.group.query
//
// 组团购场景下，查询商品开团详情
type TaobaoOpentradeGroupQueryAPIResponse struct {
	model.CommonResponse
	TaobaoOpentradeGroupQueryAPIResponseModel
}

// TaobaoOpentradeGroupQueryAPIResponseModel is 组团购场景查询商品组团详情 成功返回结果
type TaobaoOpentradeGroupQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"opentrade_group_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 团信息
	Results []GroupDetailResponse `json:"results,omitempty" xml:"results>group_detail_response,omitempty"`
	// 总记录数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

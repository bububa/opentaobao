package inventory

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLocationRelationQueryAPIResponse 地点关联关系查询 API返回值
// taobao.location.relation.query
//
// 地点关联关系查询
// 门店和仓库关联关系查询
type TaobaoLocationRelationQueryAPIResponse struct {
	model.CommonResponse
	TaobaoLocationRelationQueryAPIResponseModel
}

// TaobaoLocationRelationQueryAPIResponseModel is 地点关联关系查询 成功返回结果
type TaobaoLocationRelationQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"location_relation_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *SingleResult `json:"result,omitempty" xml:"result,omitempty"`
}

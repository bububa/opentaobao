package inventory

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoLocationRelationQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLocationRelationQueryAPIResponseModel).Reset()
}

// TaobaoLocationRelationQueryAPIResponseModel is 地点关联关系查询 成功返回结果
type TaobaoLocationRelationQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"location_relation_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *SingleResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLocationRelationQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoLocationRelationQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLocationRelationQueryAPIResponse)
	},
}

// GetTaobaoLocationRelationQueryAPIResponse 从 sync.Pool 获取 TaobaoLocationRelationQueryAPIResponse
func GetTaobaoLocationRelationQueryAPIResponse() *TaobaoLocationRelationQueryAPIResponse {
	return poolTaobaoLocationRelationQueryAPIResponse.Get().(*TaobaoLocationRelationQueryAPIResponse)
}

// ReleaseTaobaoLocationRelationQueryAPIResponse 将 TaobaoLocationRelationQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLocationRelationQueryAPIResponse(v *TaobaoLocationRelationQueryAPIResponse) {
	v.Reset()
	poolTaobaoLocationRelationQueryAPIResponse.Put(v)
}

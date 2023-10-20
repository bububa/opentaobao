package nrt

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtStoreRelationQueryAPIResponse 喵零门店关系查询 API返回值
// tmall.nrt.store.relation.query
//
// 喵零门店关系查询
type TmallNrtStoreRelationQueryAPIResponse struct {
	model.CommonResponse
	TmallNrtStoreRelationQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrtStoreRelationQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrtStoreRelationQueryAPIResponseModel).Reset()
}

// TmallNrtStoreRelationQueryAPIResponseModel is 喵零门店关系查询 成功返回结果
type TmallNrtStoreRelationQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_store_relation_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果对象
	Result *ResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrtStoreRelationQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallNrtStoreRelationQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrtStoreRelationQueryAPIResponse)
	},
}

// GetTmallNrtStoreRelationQueryAPIResponse 从 sync.Pool 获取 TmallNrtStoreRelationQueryAPIResponse
func GetTmallNrtStoreRelationQueryAPIResponse() *TmallNrtStoreRelationQueryAPIResponse {
	return poolTmallNrtStoreRelationQueryAPIResponse.Get().(*TmallNrtStoreRelationQueryAPIResponse)
}

// ReleaseTmallNrtStoreRelationQueryAPIResponse 将 TmallNrtStoreRelationQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallNrtStoreRelationQueryAPIResponse(v *TmallNrtStoreRelationQueryAPIResponse) {
	v.Reset()
	poolTmallNrtStoreRelationQueryAPIResponse.Put(v)
}

package store

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStoreRelationQueryAPIResponse 门店关系查询 API返回值
// taobao.place.store.relation.query
//
// 查询门店关系
type TaobaoPlaceStoreRelationQueryAPIResponse struct {
	model.CommonResponse
	TaobaoPlaceStoreRelationQueryAPIResponseModel
}

// TaobaoPlaceStoreRelationQueryAPIResponseModel is 门店关系查询 成功返回结果
type TaobaoPlaceStoreRelationQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"place_store_relation_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TopBatchResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

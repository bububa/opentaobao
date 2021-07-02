package lifeservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStoreRelationAddAPIResponse 门店关系新增 API返回值
// taobao.place.store.relation.add
//
// 新增授权用户的门店关系信息
type TaobaoPlaceStoreRelationAddAPIResponse struct {
	model.CommonResponse
	TaobaoPlaceStoreRelationAddAPIResponseModel
}

// TaobaoPlaceStoreRelationAddAPIResponseModel is 门店关系新增 成功返回结果
type TaobaoPlaceStoreRelationAddAPIResponseModel struct {
	XMLName xml.Name `xml:"place_store_relation_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否失败
	Failure bool `json:"failure,omitempty" xml:"failure,omitempty"`
	// 返回结果：true成功；false失败
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
	// 个数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
}

package store

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStoreQueryAPIResponse 门店信息查询接口 API返回值
// taobao.place.store.query
//
// 根据用户授权信息，获取用户的门店公开信息
type TaobaoPlaceStoreQueryAPIResponse struct {
	model.CommonResponse
	TaobaoPlaceStoreQueryAPIResponseModel
}

// TaobaoPlaceStoreQueryAPIResponseModel is 门店信息查询接口 成功返回结果
type TaobaoPlaceStoreQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"place_store_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Top返回对象
	Result *TopResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

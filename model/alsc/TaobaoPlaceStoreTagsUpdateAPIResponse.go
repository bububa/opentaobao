package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStoreTagsUpdateAPIResponse 门店打标去标 API返回值
// taobao.place.store.tags.update
//
// 门店打标去标
type TaobaoPlaceStoreTagsUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoPlaceStoreTagsUpdateAPIResponseModel
}

// TaobaoPlaceStoreTagsUpdateAPIResponseModel is 门店打标去标 成功返回结果
type TaobaoPlaceStoreTagsUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"place_store_tags_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果：true成功；false失败
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

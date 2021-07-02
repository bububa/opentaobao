package shop

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoStoreFollowurlGetAPIResponse 获取店铺关注URL API返回值
// taobao.store.followurl.get
//
// 获取关注店铺的URL
type TaobaoStoreFollowurlGetAPIResponse struct {
	model.CommonResponse
	TaobaoStoreFollowurlGetAPIResponseModel
}

// TaobaoStoreFollowurlGetAPIResponseModel is 获取店铺关注URL 成功返回结果
type TaobaoStoreFollowurlGetAPIResponseModel struct {
	XMLName xml.Name `xml:"store_followurl_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 店铺关注URL
	Url string `json:"url,omitempty" xml:"url,omitempty"`
}

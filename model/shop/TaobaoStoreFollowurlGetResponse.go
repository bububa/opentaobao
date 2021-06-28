package shop

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取店铺关注URL APIResponse
taobao.store.followurl.get

获取关注店铺的URL
*/
type TaobaoStoreFollowurlGetAPIResponse struct {
    model.CommonResponse
    TaobaoStoreFollowurlGetResponse
}

type TaobaoStoreFollowurlGetResponse struct {
    XMLName xml.Name `xml:"store_followurl_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 店铺关注URL
    
    Url   string `json:"url,omitempty" xml:"url,omitempty"`

    
}

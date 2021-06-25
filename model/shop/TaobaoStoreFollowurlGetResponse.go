package shop

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取店铺关注URL APIResponse
taobao.store.followurl.get

获取关注店铺的URL
*/
type TaobaoStoreFollowurlGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoStoreFollowurlGetResponse `json:"taobao_store_followurl_get_response,omitempty"`
}

type TaobaoStoreFollowurlGetResponse struct {

    // 店铺关注URL
    Url   string `json:"url,omitempty"`

}

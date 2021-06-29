package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店打标去标 APIResponse
taobao.place.store.tags.update

门店打标去标
*/
type TaobaoPlaceStoreTagsUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoPlaceStoreTagsUpdateResponse
}

type TaobaoPlaceStoreTagsUpdateResponse struct {
    XMLName xml.Name `xml:"place_store_tags_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果：true成功；false失败
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}

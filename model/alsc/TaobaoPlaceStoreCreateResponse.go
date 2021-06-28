package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商户门店创建接口 APIResponse
taobao.place.store.create

用于商家创建线下门店
*/
type TaobaoPlaceStoreCreateAPIResponse struct {
    model.CommonResponse
    TaobaoPlaceStoreCreateResponse
}

type TaobaoPlaceStoreCreateResponse struct {
    XMLName xml.Name `xml:"place_store_create_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    StoreId   int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`

    
}

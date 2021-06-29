package store

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店和子门店关系查找 APIResponse
taobao.place.storerelatesub.get

门店和子门店关系查找
*/
type TaobaoPlaceStorerelatesubGetAPIResponse struct {
    model.CommonResponse
    TaobaoPlaceStorerelatesubGetResponse
}

type TaobaoPlaceStorerelatesubGetResponse struct {
    XMLName xml.Name `xml:"place_storerelatesub_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TopBatchResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}

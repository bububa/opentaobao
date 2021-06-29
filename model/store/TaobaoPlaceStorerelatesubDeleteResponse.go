package store

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店和子门店关系删除 APIResponse
taobao.place.storerelatesub.delete

门店和子门店关系删除
*/
type TaobaoPlaceStorerelatesubDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoPlaceStorerelatesubDeleteResponse
}

type TaobaoPlaceStorerelatesubDeleteResponse struct {
    XMLName xml.Name `xml:"place_storerelatesub_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TopBatchResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}

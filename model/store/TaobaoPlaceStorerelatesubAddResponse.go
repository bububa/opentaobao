package store

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店和子门店关系新增 APIResponse
taobao.place.storerelatesub.add

门店和子门店关系新增
*/
type TaobaoPlaceStorerelatesubAddAPIResponse struct {
    model.CommonResponse
    TaobaoPlaceStorerelatesubAddResponse
}

type TaobaoPlaceStorerelatesubAddResponse struct {
    XMLName xml.Name `xml:"place_storerelatesub_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *TopBatchResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}

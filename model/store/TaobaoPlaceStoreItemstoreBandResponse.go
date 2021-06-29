package store

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店商品关联绑定接口 APIResponse
taobao.place.store.itemstore.band

商品和多个门店关系绑定接口
*/
type TaobaoPlaceStoreItemstoreBandAPIResponse struct {
    model.CommonResponse
    TaobaoPlaceStoreItemstoreBandResponse
}

type TaobaoPlaceStoreItemstoreBandResponse struct {
    XMLName xml.Name `xml:"place_store_itemstore_band_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TopBatchResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}

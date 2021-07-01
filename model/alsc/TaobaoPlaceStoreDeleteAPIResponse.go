package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
线下门店删除 API返回值 
taobao.place.store.delete

用于商家删除线下门店
*/
type TaobaoPlaceStoreDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoPlaceStoreDeleteAPIResponseModel
}

// 线下门店删除 成功返回结果
type TaobaoPlaceStoreDeleteAPIResponseModel struct {
    XMLName xml.Name `xml:"place_store_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 门店删除结果
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`
}

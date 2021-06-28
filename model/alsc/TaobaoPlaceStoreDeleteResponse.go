package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
线下门店删除 APIResponse
taobao.place.store.delete

用于商家删除线下门店
*/
type TaobaoPlaceStoreDeleteAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"place_store_delete_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 门店删除结果
    
    Result   bool `json:"result,omitempty" xml:"
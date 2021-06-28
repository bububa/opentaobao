package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家修改线下门店 APIResponse
taobao.place.store.modify

用于商家修改线下门店信息
*/
type TaobaoPlaceStoreModifyAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"place_store_modify_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否修改成功
    
    Result   bool `json:"result,omitempty" xml:"
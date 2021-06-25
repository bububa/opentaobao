package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家修改线下门店 APIResponse
taobao.place.store.modify

用于商家修改线下门店信息
*/
type TaobaoPlaceStoreModifyAPIResponse struct {
    model.CommonResponse
    Response *TaobaoPlaceStoreModifyResponse `json:"taobao_place_store_modify_response,omitempty"`
}

type TaobaoPlaceStoreModifyResponse struct {

    // 是否修改成功
    Result   bool `json:"result,omitempty"`

}
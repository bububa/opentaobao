package wms

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品同步 APIResponse
taobao.wlb.wms.sku.create

商品同步
*/
type TaobaoWlbWmsSkuCreateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWlbWmsSkuCreateResponse `json:"wlb_wms_sku_create_response,omitempty"` 
    TaobaoWlbWmsSkuCreateResponse
}

/* model for simplify = false
type TaobaoWlbWmsSkuCreateResponse struct {

    // 系统自动生成
    
    ItemId   int64 `json:"item_id,omitempty"`
    

    // 错误信息
    
    WlErrorMsg   string `json:"wl_error_msg,omitempty"`
    

    // 是否成功
    
    WlSuccess   bool `json:"wl_success,omitempty"`
    

    // 错误码
    
    WlErrorCode   string `json:"wl_error_code,omitempty"`
    

}
*/

type TaobaoWlbWmsSkuCreateResponse struct {

    // 系统自动生成
    ItemId   int64 `json:"item_id,omitempty"`

    // 错误信息
    WlErrorMsg   string `json:"wl_error_msg,omitempty"`

    // 是否成功
    WlSuccess   bool `json:"wl_success,omitempty"`

    // 错误码
    WlErrorCode   string `json:"wl_error_code,omitempty"`

}

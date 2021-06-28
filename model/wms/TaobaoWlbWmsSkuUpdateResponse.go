package wms

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品信息的更新 APIResponse
taobao.wlb.wms.sku.update

商品信息的更新
*/
type TaobaoWlbWmsSkuUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWlbWmsSkuUpdateResponse `json:"wlb_wms_sku_update_response,omitempty"` 
    TaobaoWlbWmsSkuUpdateResponse
}

/* model for simplify = false
type TaobaoWlbWmsSkuUpdateResponse struct {

    // 错误信息
    
    WlErrorMsg   string `json:"wl_error_msg,omitempty"`
    

    // 错误编码
    
    WlErrorCode   string `json:"wl_error_code,omitempty"`
    

    // 是否成功
    
    WlSuccess   bool `json:"wl_success,omitempty"`
    

}
*/

type TaobaoWlbWmsSkuUpdateResponse struct {

    // 错误信息
    WlErrorMsg   string `json:"wl_error_msg,omitempty"`

    // 错误编码
    WlErrorCode   string `json:"wl_error_code,omitempty"`

    // 是否成功
    WlSuccess   bool `json:"wl_success,omitempty"`

}

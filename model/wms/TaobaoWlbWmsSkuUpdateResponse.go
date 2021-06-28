package wms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品信息的更新 APIResponse
taobao.wlb.wms.sku.update

商品信息的更新
*/
type TaobaoWlbWmsSkuUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_wms_sku_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误信息
    
    WlErrorMsg   string `json:"wl_error_msg,omitempty" xml:"
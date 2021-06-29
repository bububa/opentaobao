package wms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品信息的更新 API返回值 
taobao.wlb.wms.sku.update

商品信息的更新
*/
type TaobaoWlbWmsSkuUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoWlbWmsSkuUpdateResponse
}

// 商品信息的更新 成功返回结果
type TaobaoWlbWmsSkuUpdateResponse struct {
    XMLName xml.Name `xml:"wlb_wms_sku_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 错误信息
    WlErrorMsg   string `json:"wl_error_msg,omitempty" xml:"wl_error_msg,omitempty"`
    // 错误编码
    WlErrorCode   string `json:"wl_error_code,omitempty" xml:"wl_error_code,omitempty"`
    // 是否成功
    WlSuccess   bool `json:"wl_success,omitempty" xml:"wl_success,omitempty"`
}

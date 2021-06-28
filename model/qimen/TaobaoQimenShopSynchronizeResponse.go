package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
店铺同步接口 APIResponse
taobao.qimen.shop.synchronize

店铺同步接口描述
*/
type TaobaoQimenShopSynchronizeAPIResponse struct {
    model.CommonResponse
    TaobaoQimenShopSynchronizeResponse
}

type TaobaoQimenShopSynchronizeResponse struct {
    XMLName xml.Name `xml:"qimen_shop_synchronize_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // Response
    
    Response   *Response `json:"response,omitempty" xml:"response,omitempty"`

    
}

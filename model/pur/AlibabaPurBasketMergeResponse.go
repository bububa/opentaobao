package pur

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
合并购物车 APIResponse
alibaba.pur.basket.merge

采购商城接入第三方商家合并购物车接口服务
*/
type AlibabaPurBasketMergeAPIResponse struct {
    model.CommonResponse
    AlibabaPurBasketMergeResponse
}

type AlibabaPurBasketMergeResponse struct {
    XMLName xml.Name `xml:"alibaba_pur_basket_merge_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 获取url的出参
    
    Result   *ActionResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

package mtopopen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
allspark提供抽奖tida接口对应鉴权接口 APIResponse
alibaba.interact.allsparkisv.draw

该接口没有实际对外使用。只是内部鉴权使用，不会有三方应用调用
*/
type AlibabaInteractAllsparkisvDrawAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaInteractAllsparkisvDrawResponse `json:"alibaba_interact_allsparkisv_draw_response,omitempty"` 
    AlibabaInteractAllsparkisvDrawResponse
}

/* model for simplify = false
type AlibabaInteractAllsparkisvDrawResponse struct {

    // ddd
    
    Ddd   string `json:"ddd,omitempty"`
    

}
*/

type AlibabaInteractAllsparkisvDrawResponse struct {

    // ddd
    Ddd   string `json:"ddd,omitempty"`

}

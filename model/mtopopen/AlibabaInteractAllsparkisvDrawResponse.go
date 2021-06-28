package mtopopen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
allspark提供抽奖tida接口对应鉴权接口 APIResponse
alibaba.interact.allsparkisv.draw

该接口没有实际对外使用。只是内部鉴权使用，不会有三方应用调用
*/
type AlibabaInteractAllsparkisvDrawAPIResponse struct {
    model.CommonResponse
    AlibabaInteractAllsparkisvDrawResponse
}

type AlibabaInteractAllsparkisvDrawResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_allsparkisv_draw_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // ddd
    
    Ddd   string `json:"ddd,omitempty" xml:"ddd,omitempty"`

    
}

package pur

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
下pr单 APIResponse
alibaba.pur.pr.create

下pr单
*/
type AlibabaPurPrCreateAPIResponse struct {
    model.CommonResponse
    AlibabaPurPrCreateResponse
}

type AlibabaPurPrCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_pur_pr_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回对象
    
    Result   *MallReceivePrResponse `json:"result,omitempty" xml:"result,omitempty"`

    
}

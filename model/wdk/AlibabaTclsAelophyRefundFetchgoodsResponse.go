package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
saas 售后逆向 商户发起逆向取货 APIResponse
alibaba.tcls.aelophy.refund.fetchgoods

saas 售后逆向 商户发起逆向取货
*/
type AlibabaTclsAelophyRefundFetchgoodsAPIResponse struct {
    model.CommonResponse
    AlibabaTclsAelophyRefundFetchgoodsResponse
}

type AlibabaTclsAelophyRefundFetchgoodsResponse struct {
    XMLName xml.Name `xml:"alibaba_tcls_aelophy_refund_fetchgoods_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 出参
    
    Result   *AlibabaTclsAelophyRefundFetchgoodsApiResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

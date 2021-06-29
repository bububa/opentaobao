package alisports

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用户完成支付同步订单 APIResponse
alibaba.alisports.efsp.userplaceorder

用户完成支付同步订单
*/
type AlibabaAlisportsEfspUserplaceorderAPIResponse struct {
    model.CommonResponse
    AlibabaAlisportsEfspUserplaceorderResponse
}

type AlibabaAlisportsEfspUserplaceorderResponse struct {
    XMLName xml.Name `xml:"alibaba_alisports_efsp_userplaceorder_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 系统自动生成
    
    Result   *TrilateralResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

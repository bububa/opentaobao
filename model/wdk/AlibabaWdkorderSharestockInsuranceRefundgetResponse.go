package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
共享库存投保业务售后逆向订单数据获取 APIResponse
alibaba.wdkorder.sharestock.insurance.refundget

共享库存投保业务售后逆向订单数据获取
*/
type AlibabaWdkorderSharestockInsuranceRefundgetAPIResponse struct {
    model.CommonResponse
    AlibabaWdkorderSharestockInsuranceRefundgetResponse
}

type AlibabaWdkorderSharestockInsuranceRefundgetResponse struct {
    XMLName xml.Name `xml:"alibaba_wdkorder_sharestock_insurance_refundget_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *MaochaoOrderInsuranceRefundQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

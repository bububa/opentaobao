package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
饿了么对账单查询，带订单明细 APIResponse
alibaba.wdk.eleme.bill.detail.get

查询饿了么对账单信息，带订单明细
*/
type AlibabaWdkElemeBillDetailGetAPIResponse struct {
    model.CommonResponse
    AlibabaWdkElemeBillDetailGetResponse
}

type AlibabaWdkElemeBillDetailGetResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_eleme_bill_detail_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *AlibabaWdkElemeBillDetailGetApiResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

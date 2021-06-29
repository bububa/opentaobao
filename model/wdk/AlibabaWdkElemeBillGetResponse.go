package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
饿了么日维度对账单查询 APIResponse
alibaba.wdk.eleme.bill.get

查询饿了么日维度对账单信息
*/
type AlibabaWdkElemeBillGetAPIResponse struct {
    model.CommonResponse
    AlibabaWdkElemeBillGetResponse
}

type AlibabaWdkElemeBillGetResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_eleme_bill_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *AlibabaWdkElemeBillGetApiResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

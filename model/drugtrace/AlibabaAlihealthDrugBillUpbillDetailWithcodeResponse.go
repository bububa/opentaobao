package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询上游出库单明细(带追溯码信息) APIResponse
alibaba.alihealth.drug.bill.upbill.detail.withcode

查询上游出库单明细(带追溯码信息)
*/
type AlibabaAlihealthDrugBillUpbillDetailWithcodeAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugBillUpbillDetailWithcodeResponse
}

type AlibabaAlihealthDrugBillUpbillDetailWithcodeResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_bill_upbill_detail_withcode_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 监控宝推送网站监控信息，返回结果
    
    Result   *AlibabaAlihealthDrugBillUpbillDetailWithcodeResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}

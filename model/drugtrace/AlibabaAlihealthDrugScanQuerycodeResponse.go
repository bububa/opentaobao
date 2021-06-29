package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询药监码对应的有效期和包装规格 API返回值 
alibaba.alihealth.drug.scan.querycode

查询药监码对应的有效期和包装规格
*/
type AlibabaAlihealthDrugScanQuerycodeAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugScanQuerycodeResponse
}

// 查询药监码对应的有效期和包装规格 成功返回结果
type AlibabaAlihealthDrugScanQuerycodeResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_scan_querycode_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 和三方交互最外层model对象
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

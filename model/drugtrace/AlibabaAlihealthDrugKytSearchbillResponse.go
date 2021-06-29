package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
通过时间段批量查询入出库单信息 API返回值 
alibaba.alihealth.drug.kyt.searchbill

通过时间段批量查询入出库单信息
*/
type AlibabaAlihealthDrugKytSearchbillAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytSearchbillResponse
}

// 通过时间段批量查询入出库单信息 成功返回结果
type AlibabaAlihealthDrugKytSearchbillResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_searchbill_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 监控宝推送网站监控信息，返回结果
    Result   *AlibabaAlihealthDrugKytSearchbillResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

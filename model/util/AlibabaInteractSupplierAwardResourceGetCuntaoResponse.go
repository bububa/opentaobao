package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
权益池信息查询 API返回值 
alibaba.interact.supplier.award.resource.get.cuntao

农村淘宝营销互动权益池信息查询
*/
type AlibabaInteractSupplierAwardResourceGetCuntaoAPIResponse struct {
    model.CommonResponse
    AlibabaInteractSupplierAwardResourceGetCuntaoResponse
}

// 权益池信息查询 成功返回结果
type AlibabaInteractSupplierAwardResourceGetCuntaoResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_supplier_award_resource_get_cuntao_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 监控宝推送网站监控信息，返回结果
    Result   *AlibabaInteractSupplierAwardResourceGetCuntaoResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

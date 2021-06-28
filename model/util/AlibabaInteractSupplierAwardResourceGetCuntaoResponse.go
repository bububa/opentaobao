package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
权益池信息查询 APIResponse
alibaba.interact.supplier.award.resource.get.cuntao

农村淘宝营销互动权益池信息查询
*/
type AlibabaInteractSupplierAwardResourceGetCuntaoAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaInteractSupplierAwardResourceGetCuntaoResponse `json:"alibaba_interact_supplier_award_resource_get_cuntao_response,omitempty"` 
    AlibabaInteractSupplierAwardResourceGetCuntaoResponse
}

/* model for simplify = false
type AlibabaInteractSupplierAwardResourceGetCuntaoResponse struct {

    // 监控宝推送网站监控信息，返回结果
    
    Result  *struct {
        AlibabaInteractSupplierAwardResourceGetCuntaoResultModel  *AlibabaInteractSupplierAwardResourceGetCuntaoResultModel `json:"alibaba_interact_supplier_award_resource_get_cuntao_result_model,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaInteractSupplierAwardResourceGetCuntaoResponse struct {

    // 监控宝推送网站监控信息，返回结果
    Result   *AlibabaInteractSupplierAwardResourceGetCuntaoResultModel `json:"result,omitempty"`

}

package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据天猫id查询门店服务授权 APIResponse
alibaba.mallitemcenter.entitledservice.supplier.query

根据天猫id查询门店服务授权
*/
type AlibabaMallitemcenterEntitledserviceSupplierQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaMallitemcenterEntitledserviceSupplierQueryResponse `json:"alibaba_mallitemcenter_entitledservice_supplier_query_response,omitempty"` 
    AlibabaMallitemcenterEntitledserviceSupplierQueryResponse
}

/* model for simplify = false
type AlibabaMallitemcenterEntitledserviceSupplierQueryResponse struct {

    // 统一返回结果
    
    Result  *struct {
        AlibabaMallitemcenterEntitledserviceSupplierQueryResult  *AlibabaMallitemcenterEntitledserviceSupplierQueryResult `json:"alibaba_mallitemcenter_entitledservice_supplier_query_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaMallitemcenterEntitledserviceSupplierQueryResponse struct {

    // 统一返回结果
    Result   *AlibabaMallitemcenterEntitledserviceSupplierQueryResult `json:"result,omitempty"`

}

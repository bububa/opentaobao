package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据天猫id查询门店服务授权 APIResponse
alibaba.mallitemcenter.entitledservice.supplier.query

根据天猫id查询门店服务授权
*/
type AlibabaMallitemcenterEntitledserviceSupplierQueryAPIResponse struct {
    model.CommonResponse
    AlibabaMallitemcenterEntitledserviceSupplierQueryResponse
}

type AlibabaMallitemcenterEntitledserviceSupplierQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_mallitemcenter_entitledservice_supplier_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 统一返回结果
    
    Result   *AlibabaMallitemcenterEntitledserviceSupplierQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

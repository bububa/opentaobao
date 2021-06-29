package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询仓库api APIResponse
alibaba.alihealth.tracecodeseller.warehouse.search

查询仓库api
*/
type AlibabaAlihealthTracecodesellerWarehouseSearchAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthTracecodesellerWarehouseSearchResponse
}

type AlibabaAlihealthTracecodesellerWarehouseSearchResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_tracecodeseller_warehouse_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}

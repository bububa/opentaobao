package ascpchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
货品品类查询 APIResponse
alibaba.ascp.salecategory.query

根据货品ID查询对应销售品类ID
*/
type AlibabaAscpSalecategoryQueryAPIResponse struct {
    model.CommonResponse
    AlibabaAscpSalecategoryQueryResponse
}

type AlibabaAscpSalecategoryQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_ascp_salecategory_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 货品查询结构化对象
    
    DataList   []SalecategoryQueryResponse `json:"data_list,omitempty" xml:"data_list>salecategory_query_response,omitempty"`
    
    
}

package ascpffo

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
AliExpress发货单查询API APIResponse
aliexpress.ascp.ffo.query

AE 履约发货单分页查询接口
*/
type AliexpressAscpFfoQueryAPIResponse struct {
    model.CommonResponse
    AliexpressAscpFfoQueryResponse
}

type AliexpressAscpFfoQueryResponse struct {
    XMLName xml.Name `xml:"aliexpress_ascp_ffo_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // dto
    
    Result   *PageQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

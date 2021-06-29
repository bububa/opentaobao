package ascpffo

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
AliExpress货品查询查询API APIResponse
aliexpress.ascp.item.query

AE货品查询API
*/
type AliexpressAscpItemQueryAPIResponse struct {
    model.CommonResponse
    AliexpressAscpItemQueryResponse
}

type AliexpressAscpItemQueryResponse struct {
    XMLName xml.Name `xml:"aliexpress_ascp_item_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // DTO
    
    Result   *PageQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

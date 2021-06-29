package ascpffo

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
AliExpress销退单明细查询API APIResponse
aliexpress.ascp.fro.item.query

AE履约销退单明细查询API
*/
type AliexpressAscpFroItemQueryAPIResponse struct {
    model.CommonResponse
    AliexpressAscpFroItemQueryResponse
}

type AliexpressAscpFroItemQueryResponse struct {
    XMLName xml.Name `xml:"aliexpress_ascp_fro_item_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AliexpressAscpFroItemQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

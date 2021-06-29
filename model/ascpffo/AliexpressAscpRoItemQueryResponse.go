package ascpffo

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
AliExpress退供单明细查询API APIResponse
aliexpress.ascp.ro.item.query

AE仓发 单个退供单明细查询
*/
type AliexpressAscpRoItemQueryAPIResponse struct {
    model.CommonResponse
    AliexpressAscpRoItemQueryResponse
}

type AliexpressAscpRoItemQueryResponse struct {
    XMLName xml.Name `xml:"aliexpress_ascp_ro_item_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // dto
    
    Result   *PageQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

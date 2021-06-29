package ascpffo

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
AliExpress采购单明细查询API APIResponse
aliexpress.ascp.po.item.query

AE 供应链仓发 采购单明细查询
*/
type AliexpressAscpPoItemQueryAPIResponse struct {
    model.CommonResponse
    AliexpressAscpPoItemQueryResponse
}

type AliexpressAscpPoItemQueryResponse struct {
    XMLName xml.Name `xml:"aliexpress_ascp_po_item_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // demo
    
    Result   *PageQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

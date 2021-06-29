package ascpffo

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
AliExpress采购单查询API APIResponse
aliexpress.ascp.po.query

AE仓发业务采购单查询
*/
type AliexpressAscpPoQueryAPIResponse struct {
    model.CommonResponse
    AliexpressAscpPoQueryResponse
}

type AliexpressAscpPoQueryResponse struct {
    XMLName xml.Name `xml:"aliexpress_ascp_po_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 服务出参
    
    Result   *AliexpressAscpPoQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

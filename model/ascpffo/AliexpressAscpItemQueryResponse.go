package ascpffo

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
AliExpress货品查询查询API API返回值 
aliexpress.ascp.item.query

AE货品查询API
*/
type AliexpressAscpItemQueryAPIResponse struct {
    model.CommonResponse
    AliexpressAscpItemQueryResponse
}

// AliExpress货品查询查询API 成功返回结果
type AliexpressAscpItemQueryResponse struct {
    XMLName xml.Name `xml:"aliexpress_ascp_item_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // DTO
    Result   *PageQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

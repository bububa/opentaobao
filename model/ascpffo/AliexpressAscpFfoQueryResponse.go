package ascpffo

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
AliExpress发货单查询API API返回值 
aliexpress.ascp.ffo.query

AE 履约发货单分页查询接口
*/
type AliexpressAscpFfoQueryAPIResponse struct {
    model.CommonResponse
    AliexpressAscpFfoQueryResponse
}

// AliExpress发货单查询API 成功返回结果
type AliexpressAscpFfoQueryResponse struct {
    XMLName xml.Name `xml:"aliexpress_ascp_ffo_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // dto
    Result   *PageQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

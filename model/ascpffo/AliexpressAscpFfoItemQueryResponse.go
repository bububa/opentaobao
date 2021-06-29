package ascpffo

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
AliExpress发货单明细分页查询API API返回值 
aliexpress.ascp.ffo.item.query

AE履约发货单明细分页查询
*/
type AliexpressAscpFfoItemQueryAPIResponse struct {
    model.CommonResponse
    AliexpressAscpFfoItemQueryResponse
}

// AliExpress发货单明细分页查询API 成功返回结果
type AliexpressAscpFfoItemQueryResponse struct {
    XMLName xml.Name `xml:"aliexpress_ascp_ffo_item_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *AliexpressAscpFfoItemQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

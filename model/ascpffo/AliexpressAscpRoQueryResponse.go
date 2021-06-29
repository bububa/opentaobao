package ascpffo

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
AliExpress退供单查询API APIResponse
aliexpress.ascp.ro.query

AE仓发商家单个退供单查询接口
*/
type AliexpressAscpRoQueryAPIResponse struct {
    model.CommonResponse
    AliexpressAscpRoQueryResponse
}

type AliexpressAscpRoQueryResponse struct {
    XMLName xml.Name `xml:"aliexpress_ascp_ro_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AliexpressAscpRoQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

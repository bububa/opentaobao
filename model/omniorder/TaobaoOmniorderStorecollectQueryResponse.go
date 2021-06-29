package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
全渠道门店自提根据核销码查询订单 APIResponse
taobao.omniorder.storecollect.query

全渠道门店自提根据核销码查询订单
*/
type TaobaoOmniorderStorecollectQueryAPIResponse struct {
    model.CommonResponse
    TaobaoOmniorderStorecollectQueryResponse
}

type TaobaoOmniorderStorecollectQueryResponse struct {
    XMLName xml.Name `xml:"omniorder_storecollect_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoOmniorderStorecollectQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

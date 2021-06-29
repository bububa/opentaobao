package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
全渠道门店自提根据核销码查询订单 API返回值 
taobao.omniorder.storecollect.query

全渠道门店自提根据核销码查询订单
*/
type TaobaoOmniorderStorecollectQueryAPIResponse struct {
    model.CommonResponse
    TaobaoOmniorderStorecollectQueryResponse
}

// 全渠道门店自提根据核销码查询订单 成功返回结果
type TaobaoOmniorderStorecollectQueryResponse struct {
    XMLName xml.Name `xml:"omniorder_storecollect_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TaobaoOmniorderStorecollectQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

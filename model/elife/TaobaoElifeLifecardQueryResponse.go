package elife

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询交易结果 APIResponse
taobao.elife.lifecard.query

卖家在交易状态不明的情况下, 查询交易结果.
*/
type TaobaoElifeLifecardQueryAPIResponse struct {
    model.CommonResponse
    TaobaoElifeLifecardQueryResponse
}

type TaobaoElifeLifecardQueryResponse struct {
    XMLName xml.Name `xml:"elife_lifecard_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // resultMsg
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
    // amount
    
    Amount   int64 `json:"amount,omitempty" xml:"amount,omitempty"`

    
    // successed
    
    Successed   bool `json:"successed,omitempty" xml:"successed,omitempty"`

    
    // resultCode
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // inflateAmount
    
    InflateAmount   int64 `json:"inflate_amount,omitempty" xml:"inflate_amount,omitempty"`

    
}

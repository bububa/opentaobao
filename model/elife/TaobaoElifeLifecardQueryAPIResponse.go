package elife

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询交易结果 API返回值 
taobao.elife.lifecard.query

卖家在交易状态不明的情况下, 查询交易结果.
*/
type TaobaoElifeLifecardQueryAPIResponse struct {
    model.CommonResponse
    TaobaoElifeLifecardQueryAPIResponseModel
}

// 查询交易结果 成功返回结果
type TaobaoElifeLifecardQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"elife_lifecard_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
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

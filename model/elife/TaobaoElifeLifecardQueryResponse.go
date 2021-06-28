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
	RequestId     string         `json:"request_id,omitempty" xml:"elife_lifecard_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // resultMsg
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"
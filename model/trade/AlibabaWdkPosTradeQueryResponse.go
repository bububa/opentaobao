package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
轻pos品牌营销查询接口 APIResponse
alibaba.wdk.pos.trade.query

轻pos品牌营销场景，外部商家查询营销信息
*/
type AlibabaWdkPosTradeQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_pos_trade_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 查询返回结果
    
    Result   *FastBuyPosQueryResult `json:"result,omitempty" xml:"
package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌营销的订单活动信息查询 APIResponse
alibaba.wdk.bm.trade.activity.query

品牌营销的订单活动信息查询
*/
type AlibabaWdkBmTradeActivityQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkBmTradeActivityQueryResponse
}

type AlibabaWdkBmTradeActivityQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_bm_trade_activity_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果数据
    
    Result   *BmResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

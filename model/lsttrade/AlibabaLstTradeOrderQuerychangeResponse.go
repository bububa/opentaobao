package lsttrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单id批量查询（品牌商视角） APIResponse
alibaba.lst.trade.order.querychange

根据品牌和时间段查询有变更记录的订单id
*/
type AlibabaLstTradeOrderQuerychangeAPIResponse struct {
    model.CommonResponse
    AlibabaLstTradeOrderQuerychangeResponse
}

type AlibabaLstTradeOrderQuerychangeResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_trade_order_querychange_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 系统自动生成
    
    Result   *PagedResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}

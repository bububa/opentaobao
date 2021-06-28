package traderate

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索评价信息 APIResponse
taobao.traderates.get

搜索评价信息，只能获取距今180天内的评价记录(只支持查询卖家给出或得到的评价)。
*/
type TaobaoTraderatesGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"traderates_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 评价列表。返回的TradeRate包含的具体信息为入参fields请求的字段信息
    
    TradeRates   []TradeRate `json:"trade_rates,omitempty" xml:"
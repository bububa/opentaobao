package traveltrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
添加一笔交易备注 API返回值 
taobao.alitrip.travel.trade.memo.add

对一笔交易添加备注
*/
type TaobaoAlitripTravelTradeMemoAddAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripTravelTradeMemoAddResponse
}

// 添加一笔交易备注 成功返回结果
type TaobaoAlitripTravelTradeMemoAddResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_trade_memo_add_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 交易添加备注返回值
    MemoAdd   *MemoCreate `json:"memo_add,omitempty" xml:"memo_add,omitempty"`
}

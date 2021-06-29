package traveltrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改一笔交易备注 APIResponse
taobao.alitrip.travel.trade.memo.update

更新一笔交易备注
*/
type TaobaoAlitripTravelTradeMemoUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripTravelTradeMemoUpdateResponse
}

type TaobaoAlitripTravelTradeMemoUpdateResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_trade_memo_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 修改交易备注返回值
    
    MemoUpdate   *MemoUpdate `json:"memo_update,omitempty" xml:"memo_update,omitempty"`

    
}

package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取网厅号卡垂直标信息 APIResponse
taobao.wtt.trade.service.get

查询网厅订单信息
*/
type TaobaoWttTradeServiceGetAPIResponse struct {
    model.CommonResponse
    TaobaoWttTradeServiceGetResponse
}

type TaobaoWttTradeServiceGetResponse struct {
    XMLName xml.Name `xml:"wtt_trade_service_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回实例
    
    Modules   *WtverticalDto `json:"modules,omitempty" xml:"modules,omitempty"`

    
}

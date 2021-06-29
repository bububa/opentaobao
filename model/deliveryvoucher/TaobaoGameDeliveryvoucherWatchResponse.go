package deliveryvoucher

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
监控预约数据 APIResponse
taobao.game.deliveryvoucher.watch

监控预约数据
*/
type TaobaoGameDeliveryvoucherWatchAPIResponse struct {
    model.CommonResponse
    TaobaoGameDeliveryvoucherWatchResponse
}

type TaobaoGameDeliveryvoucherWatchResponse struct {
    XMLName xml.Name `xml:"game_deliveryvoucher_watch_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // code
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // isSuccess
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}

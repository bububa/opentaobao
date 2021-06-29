package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵酒店播放暂停 APIResponse
taobao.tmallgenie.hotelplayerpause

酒店推送指令给天猫精灵停止播放音乐
*/
type TaobaoTmallgenieHotelplayerpauseAPIResponse struct {
    model.CommonResponse
    TaobaoTmallgenieHotelplayerpauseResponse
}

type TaobaoTmallgenieHotelplayerpauseResponse struct {
    XMLName xml.Name `xml:"tmallgenie_hotelplayerpause_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // statusCode
    
    StatusCode   int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`

    
    // message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
}

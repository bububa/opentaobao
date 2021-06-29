package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店欢迎词推送 APIResponse
taobao.tmallgenie.hotelwelcome

推送欢迎词，让天猫精灵播放
*/
type TaobaoTmallgenieHotelwelcomeAPIResponse struct {
    model.CommonResponse
    TaobaoTmallgenieHotelwelcomeResponse
}

type TaobaoTmallgenieHotelwelcomeResponse struct {
    XMLName xml.Name `xml:"tmallgenie_hotelwelcome_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // statusCode
    
    StatusCode   int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`

    
    // message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
}

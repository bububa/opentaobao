package larkiot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取iot渠道开放的影院 APIResponse
taobao.lark.iot.order.getcinemas

iot渠道拉取有权限访问的影院
*/
type TaobaoLarkIotOrderGetcinemasAPIResponse struct {
    model.CommonResponse
    TaobaoLarkIotOrderGetcinemasResponse
}

type TaobaoLarkIotOrderGetcinemasResponse struct {
    XMLName xml.Name `xml:"lark_iot_order_getcinemas_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 有权限的影院列表
    
    Data   *BizListResult `json:"data,omitempty" xml:"data,omitempty"`

    
}

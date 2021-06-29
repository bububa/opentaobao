package xhotelonlineorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分销渠道各类通知接口 APIResponse
alitrip.xhotel.channel.notify

分销渠道支付通知
*/
type AlitripXhotelChannelNotifyAPIResponse struct {
    model.CommonResponse
    AlitripXhotelChannelNotifyResponse
}

type AlitripXhotelChannelNotifyResponse struct {
    XMLName xml.Name `xml:"alitrip_xhotel_channel_notify_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果信息
    
    Result   *HbsResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

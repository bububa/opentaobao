package xhotelonlineorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
渠道分销创建订单接口 APIResponse
alitrip.xhotel.channel.order.create

创建订单接口服务（如菲住等其他渠道分销提供）
*/
type AlitripXhotelChannelOrderCreateAPIResponse struct {
    model.CommonResponse
    AlitripXhotelChannelOrderCreateResponse
}

type AlitripXhotelChannelOrderCreateResponse struct {
    XMLName xml.Name `xml:"alitrip_xhotel_channel_order_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *HbsResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

package ascpchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘外分销逆向创单（未发货整单退） APIResponse
alibaba.ascp.channel.main.refund.create

淘外分销解决方案--订单--逆向创单（未发货整单退）
*/
type AlibabaAscpChannelMainRefundCreateAPIResponse struct {
    model.CommonResponse
    AlibabaAscpChannelMainRefundCreateResponse
}

type AlibabaAscpChannelMainRefundCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_ascp_channel_main_refund_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回值包装,result为返回具体消息内容
    
    Result   *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`

    
}

package ascpchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘外分销逆向创单（子单退） APIResponse
alibaba.ascp.channel.sub.refund.create

淘外分销逆向创单（子单退）
*/
type AlibabaAscpChannelSubRefundCreateAPIResponse struct {
    model.CommonResponse
    AlibabaAscpChannelSubRefundCreateResponse
}

type AlibabaAscpChannelSubRefundCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_ascp_channel_sub_refund_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 异步获取历史数据接口返回结果
    
    Result   *AlibabaAscpChannelSubRefundCreateResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}

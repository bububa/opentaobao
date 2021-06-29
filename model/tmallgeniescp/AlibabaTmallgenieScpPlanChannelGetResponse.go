package tmallgeniescp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
5-IBP同步渠道接口 APIResponse
alibaba.tmallgenie.scp.plan.channel.get

IBP同步渠道接口
*/
type AlibabaTmallgenieScpPlanChannelGetAPIResponse struct {
    model.CommonResponse
    AlibabaTmallgenieScpPlanChannelGetResponse
}

type AlibabaTmallgenieScpPlanChannelGetResponse struct {
    XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_channel_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回对象封装
    
    Result   *DataResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

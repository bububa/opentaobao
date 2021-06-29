package ioti

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
线上巡店Agent获取配置 APIResponse
alibaba.it.cloudlive.getagentconfig

线上巡店应用，外部Agent设备获取设备配置信息，根据配置信息链接mqtt，跟云端进行进一步的消息通信。
*/
type AlibabaItCloudliveGetagentconfigAPIResponse struct {
    model.CommonResponse
    AlibabaItCloudliveGetagentconfigResponse
}

type AlibabaItCloudliveGetagentconfigResponse struct {
    XMLName xml.Name `xml:"alibaba_it_cloudlive_getagentconfig_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 消息内容
    
    RespContent   *AgentConfigResponse `json:"resp_content,omitempty" xml:"resp_content,omitempty"`

    
}

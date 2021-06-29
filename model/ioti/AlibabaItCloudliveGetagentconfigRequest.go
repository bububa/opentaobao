package ioti

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线上巡店Agent获取配置 APIRequest
alibaba.it.cloudlive.getagentconfig

线上巡店应用，外部Agent设备获取设备配置信息，根据配置信息链接mqtt，跟云端进行进一步的消息通信。
*/
type AlibabaItCloudliveGetagentconfigRequest struct {
    model.Params

    // agent标识信息
    agentId   string 

    // 时间戳
    timeStamp   int64 

    // 签名
    signature   string 

    // 设备所在IP地址
    agentIp   string 

}

func NewAlibabaItCloudliveGetagentconfigRequest() *AlibabaItCloudliveGetagentconfigRequest{
    return &AlibabaItCloudliveGetagentconfigRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaItCloudliveGetagentconfigRequest) GetApiMethodName() string {
    return "alibaba.it.cloudlive.getagentconfig"
}

func (r AlibabaItCloudliveGetagentconfigRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaItCloudliveGetagentconfigRequest) SetAgentId(agentId string) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

func (r AlibabaItCloudliveGetagentconfigRequest) GetAgentId() string {
    return r.agentId
}

func (r *AlibabaItCloudliveGetagentconfigRequest) SetTimeStamp(timeStamp int64) error {
    r.timeStamp = timeStamp
    r.Set("time_stamp", timeStamp)
    return nil
}

func (r AlibabaItCloudliveGetagentconfigRequest) GetTimeStamp() int64 {
    return r.timeStamp
}

func (r *AlibabaItCloudliveGetagentconfigRequest) SetSignature(signature string) error {
    r.signature = signature
    r.Set("signature", signature)
    return nil
}

func (r AlibabaItCloudliveGetagentconfigRequest) GetSignature() string {
    return r.signature
}

func (r *AlibabaItCloudliveGetagentconfigRequest) SetAgentIp(agentIp string) error {
    r.agentIp = agentIp
    r.Set("agent_ip", agentIp)
    return nil
}

func (r AlibabaItCloudliveGetagentconfigRequest) GetAgentIp() string {
    return r.agentIp
}


package ioti

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线上巡店Agent获取配置 API请求
alibaba.it.cloudlive.getagentconfig

线上巡店应用，外部Agent设备获取设备配置信息，根据配置信息链接mqtt，跟云端进行进一步的消息通信。
*/
type AlibabaItCloudliveGetagentconfigRequest struct {
    model.Params
    // agent标识信息
    _agentId   string
    // 时间戳
    _timeStamp   int64
    // 签名
    _signature   string
    // 设备所在IP地址
    _agentIp   string
}

// 初始化AlibabaItCloudliveGetagentconfigRequest对象
func NewAlibabaItCloudliveGetagentconfigRequest() *AlibabaItCloudliveGetagentconfigRequest{
    return &AlibabaItCloudliveGetagentconfigRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaItCloudliveGetagentconfigRequest) GetApiMethodName() string {
    return "alibaba.it.cloudlive.getagentconfig"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaItCloudliveGetagentconfigRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AgentId Setter
// agent标识信息
func (r *AlibabaItCloudliveGetagentconfigRequest) SetAgentId(_agentId string) error {
    r._agentId = _agentId
    r.Set("agent_id", _agentId)
    return nil
}

// AgentId Getter
func (r AlibabaItCloudliveGetagentconfigRequest) GetAgentId() string {
    return r._agentId
}
// TimeStamp Setter
// 时间戳
func (r *AlibabaItCloudliveGetagentconfigRequest) SetTimeStamp(_timeStamp int64) error {
    r._timeStamp = _timeStamp
    r.Set("time_stamp", _timeStamp)
    return nil
}

// TimeStamp Getter
func (r AlibabaItCloudliveGetagentconfigRequest) GetTimeStamp() int64 {
    return r._timeStamp
}
// Signature Setter
// 签名
func (r *AlibabaItCloudliveGetagentconfigRequest) SetSignature(_signature string) error {
    r._signature = _signature
    r.Set("signature", _signature)
    return nil
}

// Signature Getter
func (r AlibabaItCloudliveGetagentconfigRequest) GetSignature() string {
    return r._signature
}
// AgentIp Setter
// 设备所在IP地址
func (r *AlibabaItCloudliveGetagentconfigRequest) SetAgentIp(_agentIp string) error {
    r._agentIp = _agentIp
    r.Set("agent_ip", _agentIp)
    return nil
}

// AgentIp Getter
func (r AlibabaItCloudliveGetagentconfigRequest) GetAgentIp() string {
    return r._agentIp
}

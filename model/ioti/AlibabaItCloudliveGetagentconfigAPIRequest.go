package ioti

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItCloudliveGetagentconfigAPIRequest 线上巡店Agent获取配置 API请求
// alibaba.it.cloudlive.getagentconfig
//
// 线上巡店应用，外部Agent设备获取设备配置信息，根据配置信息链接mqtt，跟云端进行进一步的消息通信。
type AlibabaItCloudliveGetagentconfigAPIRequest struct {
	model.Params
	// agent标识信息
	_agentId string
	// 签名
	_signature string
	// 设备所在IP地址
	_agentIp string
	// 时间戳
	_timeStamp int64
}

// NewAlibabaItCloudliveGetagentconfigRequest 初始化AlibabaItCloudliveGetagentconfigAPIRequest对象
func NewAlibabaItCloudliveGetagentconfigRequest() *AlibabaItCloudliveGetagentconfigAPIRequest {
	return &AlibabaItCloudliveGetagentconfigAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaItCloudliveGetagentconfigAPIRequest) GetApiMethodName() string {
	return "alibaba.it.cloudlive.getagentconfig"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaItCloudliveGetagentconfigAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAgentId is AgentId Setter
// agent标识信息
func (r *AlibabaItCloudliveGetagentconfigAPIRequest) SetAgentId(_agentId string) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r AlibabaItCloudliveGetagentconfigAPIRequest) GetAgentId() string {
	return r._agentId
}

// SetSignature is Signature Setter
// 签名
func (r *AlibabaItCloudliveGetagentconfigAPIRequest) SetSignature(_signature string) error {
	r._signature = _signature
	r.Set("signature", _signature)
	return nil
}

// GetSignature Signature Getter
func (r AlibabaItCloudliveGetagentconfigAPIRequest) GetSignature() string {
	return r._signature
}

// SetAgentIp is AgentIp Setter
// 设备所在IP地址
func (r *AlibabaItCloudliveGetagentconfigAPIRequest) SetAgentIp(_agentIp string) error {
	r._agentIp = _agentIp
	r.Set("agent_ip", _agentIp)
	return nil
}

// GetAgentIp AgentIp Getter
func (r AlibabaItCloudliveGetagentconfigAPIRequest) GetAgentIp() string {
	return r._agentIp
}

// SetTimeStamp is TimeStamp Setter
// 时间戳
func (r *AlibabaItCloudliveGetagentconfigAPIRequest) SetTimeStamp(_timeStamp int64) error {
	r._timeStamp = _timeStamp
	r.Set("time_stamp", _timeStamp)
	return nil
}

// GetTimeStamp TimeStamp Getter
func (r AlibabaItCloudliveGetagentconfigAPIRequest) GetTimeStamp() int64 {
	return r._timeStamp
}

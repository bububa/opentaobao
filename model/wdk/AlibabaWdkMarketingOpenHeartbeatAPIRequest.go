package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingOpenHeartbeatAPIRequest 心跳服务【10s一次】 API请求
// alibaba.wdk.marketing.open.heartbeat
//
// 商家数据同步心跳服务
type AlibabaWdkMarketingOpenHeartbeatAPIRequest struct {
	model.Params
	// 心跳信息
	_heartBeat *HeartBeatBo
}

// NewAlibabaWdkMarketingOpenHeartbeatRequest 初始化AlibabaWdkMarketingOpenHeartbeatAPIRequest对象
func NewAlibabaWdkMarketingOpenHeartbeatRequest() *AlibabaWdkMarketingOpenHeartbeatAPIRequest {
	return &AlibabaWdkMarketingOpenHeartbeatAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingOpenHeartbeatAPIRequest) Reset() {
	r._heartBeat = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingOpenHeartbeatAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.open.heartbeat"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingOpenHeartbeatAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingOpenHeartbeatAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHeartBeat is HeartBeat Setter
// 心跳信息
func (r *AlibabaWdkMarketingOpenHeartbeatAPIRequest) SetHeartBeat(_heartBeat *HeartBeatBo) error {
	r._heartBeat = _heartBeat
	r.Set("heart_beat", _heartBeat)
	return nil
}

// GetHeartBeat HeartBeat Getter
func (r AlibabaWdkMarketingOpenHeartbeatAPIRequest) GetHeartBeat() *HeartBeatBo {
	return r._heartBeat
}

var poolAlibabaWdkMarketingOpenHeartbeatAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingOpenHeartbeatRequest()
	},
}

// GetAlibabaWdkMarketingOpenHeartbeatRequest 从 sync.Pool 获取 AlibabaWdkMarketingOpenHeartbeatAPIRequest
func GetAlibabaWdkMarketingOpenHeartbeatAPIRequest() *AlibabaWdkMarketingOpenHeartbeatAPIRequest {
	return poolAlibabaWdkMarketingOpenHeartbeatAPIRequest.Get().(*AlibabaWdkMarketingOpenHeartbeatAPIRequest)
}

// ReleaseAlibabaWdkMarketingOpenHeartbeatAPIRequest 将 AlibabaWdkMarketingOpenHeartbeatAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingOpenHeartbeatAPIRequest(v *AlibabaWdkMarketingOpenHeartbeatAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingOpenHeartbeatAPIRequest.Put(v)
}

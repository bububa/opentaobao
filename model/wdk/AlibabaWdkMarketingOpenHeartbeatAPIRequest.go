package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingOpenHeartbeatAPIRequest
心跳服务【10s一次】 API请求
alibaba.wdk.marketing.open.heartbeat

商家数据同步心跳服务 */
type AlibabaWdkMarketingOpenHeartbeatAPIRequest struct {
	model.Params
	// 心跳信息
	_heartBeat *HeartBeatBo
}

// NewAlibabaWdkMarketingOpenHeartbeatRequest 初始化AlibabaWdkMarketingOpenHeartbeatAPIRequest对象
func NewAlibabaWdkMarketingOpenHeartbeatRequest() *AlibabaWdkMarketingOpenHeartbeatAPIRequest {
	return &AlibabaWdkMarketingOpenHeartbeatAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingOpenHeartbeatAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.open.heartbeat"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingOpenHeartbeatAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is HeartBeat Setter
// 心跳信息
func (r *AlibabaWdkMarketingOpenHeartbeatAPIRequest) SetHeartBeat(_heartBeat *HeartBeatBo) error {
	r._heartBeat = _heartBeat
	r.Set("heart_beat", _heartBeat)
	return nil
}

// Get HeartBeat Getter
func (r AlibabaWdkMarketingOpenHeartbeatAPIRequest) GetHeartBeat() *HeartBeatBo {
	return r._heartBeat
}

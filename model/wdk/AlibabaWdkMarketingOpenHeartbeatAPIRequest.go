package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingopenheartbeatAPIRequest 心跳服务【10s一次】 API请求
// alibaba.wdk.marketing.open.heartbeat
//
// 商家数据同步心跳服务
type AlibabawdkmarketingopenheartbeatAPIRequest struct {
	model.Params
	// 心跳信息
	_heartBeat *HeartBeatBo
}

// NewAlibabawdkmarketingopenheartbeatRequest 初始化AlibabawdkmarketingopenheartbeatAPIRequest对象
func NewAlibabawdkmarketingopenheartbeatRequest() *AlibabawdkmarketingopenheartbeatAPIRequest {
	return &AlibabawdkmarketingopenheartbeatAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingopenheartbeatAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.open.heartbeat"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingopenheartbeatAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingopenheartbeatAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHeartBeat is HeartBeat Setter
// 心跳信息
func (r *AlibabawdkmarketingopenheartbeatAPIRequest) SetHeartBeat(_heartBeat *HeartBeatBo) error {
	r._heartBeat = _heartBeat
	r.Set("heart_beat", _heartBeat)
	return nil
}

// GetHeartBeat HeartBeat Getter
func (r AlibabawdkmarketingopenheartbeatAPIRequest) GetHeartBeat() *HeartBeatBo {
	return r._heartBeat
}

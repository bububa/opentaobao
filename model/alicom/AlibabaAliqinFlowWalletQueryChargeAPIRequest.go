package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinflowwalletquerychargeAPIRequest 查询流量充值状态 API请求
// alibaba.aliqin.flow.wallet.query.charge
//
// 查询流量充值状态
type AlibabaaliqinflowwalletquerychargeAPIRequest struct {
	model.Params
	// 唯一流水号
	_outRechargeId string
	// 渠道id
	_channelId string
}

// NewAlibabaaliqinflowwalletquerychargeRequest 初始化AlibabaaliqinflowwalletquerychargeAPIRequest对象
func NewAlibabaaliqinflowwalletquerychargeRequest() *AlibabaaliqinflowwalletquerychargeAPIRequest {
	return &AlibabaaliqinflowwalletquerychargeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaliqinflowwalletquerychargeAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.flow.wallet.query.charge"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaliqinflowwalletquerychargeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaliqinflowwalletquerychargeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutRechargeId is OutRechargeId Setter
// 唯一流水号
func (r *AlibabaaliqinflowwalletquerychargeAPIRequest) SetOutRechargeId(_outRechargeId string) error {
	r._outRechargeId = _outRechargeId
	r.Set("out_recharge_id", _outRechargeId)
	return nil
}

// GetOutRechargeId OutRechargeId Getter
func (r AlibabaaliqinflowwalletquerychargeAPIRequest) GetOutRechargeId() string {
	return r._outRechargeId
}

// SetChannelId is ChannelId Setter
// 渠道id
func (r *AlibabaaliqinflowwalletquerychargeAPIRequest) SetChannelId(_channelId string) error {
	r._channelId = _channelId
	r.Set("channel_id", _channelId)
	return nil
}

// GetChannelId ChannelId Getter
func (r AlibabaaliqinflowwalletquerychargeAPIRequest) GetChannelId() string {
	return r._channelId
}

package alicom

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFlowWalletQueryChargeAPIRequest 查询流量充值状态 API请求
// alibaba.aliqin.flow.wallet.query.charge
//
// 查询流量充值状态
type AlibabaAliqinFlowWalletQueryChargeAPIRequest struct {
	model.Params
	// 唯一流水号
	_outRechargeId string
	// 渠道id
	_channelId string
}

// NewAlibabaAliqinFlowWalletQueryChargeRequest 初始化AlibabaAliqinFlowWalletQueryChargeAPIRequest对象
func NewAlibabaAliqinFlowWalletQueryChargeRequest() *AlibabaAliqinFlowWalletQueryChargeAPIRequest {
	return &AlibabaAliqinFlowWalletQueryChargeAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAliqinFlowWalletQueryChargeAPIRequest) Reset() {
	r._outRechargeId = ""
	r._channelId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFlowWalletQueryChargeAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.flow.wallet.query.charge"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFlowWalletQueryChargeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAliqinFlowWalletQueryChargeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutRechargeId is OutRechargeId Setter
// 唯一流水号
func (r *AlibabaAliqinFlowWalletQueryChargeAPIRequest) SetOutRechargeId(_outRechargeId string) error {
	r._outRechargeId = _outRechargeId
	r.Set("out_recharge_id", _outRechargeId)
	return nil
}

// GetOutRechargeId OutRechargeId Getter
func (r AlibabaAliqinFlowWalletQueryChargeAPIRequest) GetOutRechargeId() string {
	return r._outRechargeId
}

// SetChannelId is ChannelId Setter
// 渠道id
func (r *AlibabaAliqinFlowWalletQueryChargeAPIRequest) SetChannelId(_channelId string) error {
	r._channelId = _channelId
	r.Set("channel_id", _channelId)
	return nil
}

// GetChannelId ChannelId Getter
func (r AlibabaAliqinFlowWalletQueryChargeAPIRequest) GetChannelId() string {
	return r._channelId
}

var poolAlibabaAliqinFlowWalletQueryChargeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAliqinFlowWalletQueryChargeRequest()
	},
}

// GetAlibabaAliqinFlowWalletQueryChargeRequest 从 sync.Pool 获取 AlibabaAliqinFlowWalletQueryChargeAPIRequest
func GetAlibabaAliqinFlowWalletQueryChargeAPIRequest() *AlibabaAliqinFlowWalletQueryChargeAPIRequest {
	return poolAlibabaAliqinFlowWalletQueryChargeAPIRequest.Get().(*AlibabaAliqinFlowWalletQueryChargeAPIRequest)
}

// ReleaseAlibabaAliqinFlowWalletQueryChargeAPIRequest 将 AlibabaAliqinFlowWalletQueryChargeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAliqinFlowWalletQueryChargeAPIRequest(v *AlibabaAliqinFlowWalletQueryChargeAPIRequest) {
	v.Reset()
	poolAlibabaAliqinFlowWalletQueryChargeAPIRequest.Put(v)
}

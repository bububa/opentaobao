package tmallchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoChannelTradePrepayOfflineReduceAPIRequest
渠道分销供应商上传线下流水预存款（减少） API请求
taobao.channel.trade.prepay.offline.reduce

渠道分销供应商上传线下流水预存款（减少） */
type TaobaoChannelTradePrepayOfflineReduceAPIRequest struct {
	model.Params
	// 减少流水
	_offlineReducePrepayParam *TopOfflineReducePrepayDto
}

// NewTaobaoChannelTradePrepayOfflineReduceRequest 初始化TaobaoChannelTradePrepayOfflineReduceAPIRequest对象
func NewTaobaoChannelTradePrepayOfflineReduceRequest() *TaobaoChannelTradePrepayOfflineReduceAPIRequest {
	return &TaobaoChannelTradePrepayOfflineReduceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoChannelTradePrepayOfflineReduceAPIRequest) GetApiMethodName() string {
	return "taobao.channel.trade.prepay.offline.reduce"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoChannelTradePrepayOfflineReduceAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OfflineReducePrepayParam Setter
// 减少流水
func (r *TaobaoChannelTradePrepayOfflineReduceAPIRequest) SetOfflineReducePrepayParam(_offlineReducePrepayParam *TopOfflineReducePrepayDto) error {
	r._offlineReducePrepayParam = _offlineReducePrepayParam
	r.Set("offline_reduce_prepay_param", _offlineReducePrepayParam)
	return nil
}

// Get OfflineReducePrepayParam Getter
func (r TaobaoChannelTradePrepayOfflineReduceAPIRequest) GetOfflineReducePrepayParam() *TopOfflineReducePrepayDto {
	return r._offlineReducePrepayParam
}

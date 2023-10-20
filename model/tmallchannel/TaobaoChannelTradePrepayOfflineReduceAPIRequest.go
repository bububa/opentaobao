package tmallchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoChannelTradePrepayOfflineReduceAPIRequest 渠道分销供应商上传线下流水预存款（减少） API请求
// taobao.channel.trade.prepay.offline.reduce
//
// 渠道分销供应商上传线下流水预存款（减少）
type TaobaoChannelTradePrepayOfflineReduceAPIRequest struct {
	model.Params
	// 减少流水
	_offlineReducePrepayParam *TopOfflineReducePrepayDto
}

// NewTaobaoChannelTradePrepayOfflineReduceRequest 初始化TaobaoChannelTradePrepayOfflineReduceAPIRequest对象
func NewTaobaoChannelTradePrepayOfflineReduceRequest() *TaobaoChannelTradePrepayOfflineReduceAPIRequest {
	return &TaobaoChannelTradePrepayOfflineReduceAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoChannelTradePrepayOfflineReduceAPIRequest) Reset() {
	r._offlineReducePrepayParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoChannelTradePrepayOfflineReduceAPIRequest) GetApiMethodName() string {
	return "taobao.channel.trade.prepay.offline.reduce"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoChannelTradePrepayOfflineReduceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoChannelTradePrepayOfflineReduceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOfflineReducePrepayParam is OfflineReducePrepayParam Setter
// 减少流水
func (r *TaobaoChannelTradePrepayOfflineReduceAPIRequest) SetOfflineReducePrepayParam(_offlineReducePrepayParam *TopOfflineReducePrepayDto) error {
	r._offlineReducePrepayParam = _offlineReducePrepayParam
	r.Set("offline_reduce_prepay_param", _offlineReducePrepayParam)
	return nil
}

// GetOfflineReducePrepayParam OfflineReducePrepayParam Getter
func (r TaobaoChannelTradePrepayOfflineReduceAPIRequest) GetOfflineReducePrepayParam() *TopOfflineReducePrepayDto {
	return r._offlineReducePrepayParam
}

var poolTaobaoChannelTradePrepayOfflineReduceAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoChannelTradePrepayOfflineReduceRequest()
	},
}

// GetTaobaoChannelTradePrepayOfflineReduceRequest 从 sync.Pool 获取 TaobaoChannelTradePrepayOfflineReduceAPIRequest
func GetTaobaoChannelTradePrepayOfflineReduceAPIRequest() *TaobaoChannelTradePrepayOfflineReduceAPIRequest {
	return poolTaobaoChannelTradePrepayOfflineReduceAPIRequest.Get().(*TaobaoChannelTradePrepayOfflineReduceAPIRequest)
}

// ReleaseTaobaoChannelTradePrepayOfflineReduceAPIRequest 将 TaobaoChannelTradePrepayOfflineReduceAPIRequest 放入 sync.Pool
func ReleaseTaobaoChannelTradePrepayOfflineReduceAPIRequest(v *TaobaoChannelTradePrepayOfflineReduceAPIRequest) {
	v.Reset()
	poolTaobaoChannelTradePrepayOfflineReduceAPIRequest.Put(v)
}

package tmallchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoChannelTradePrepayOfflineAddAPIRequest 渠道分销供应商上传线下流水预存款（增加） API请求
// taobao.channel.trade.prepay.offline.add
//
// 渠道分销供应商上传线下流水预存款（增加）
type TaobaoChannelTradePrepayOfflineAddAPIRequest struct {
	model.Params
	// 增加流水
	_offlineAddPrepayParam *TopOfflineAddPrepayDto
}

// NewTaobaoChannelTradePrepayOfflineAddRequest 初始化TaobaoChannelTradePrepayOfflineAddAPIRequest对象
func NewTaobaoChannelTradePrepayOfflineAddRequest() *TaobaoChannelTradePrepayOfflineAddAPIRequest {
	return &TaobaoChannelTradePrepayOfflineAddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoChannelTradePrepayOfflineAddAPIRequest) Reset() {
	r._offlineAddPrepayParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoChannelTradePrepayOfflineAddAPIRequest) GetApiMethodName() string {
	return "taobao.channel.trade.prepay.offline.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoChannelTradePrepayOfflineAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoChannelTradePrepayOfflineAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOfflineAddPrepayParam is OfflineAddPrepayParam Setter
// 增加流水
func (r *TaobaoChannelTradePrepayOfflineAddAPIRequest) SetOfflineAddPrepayParam(_offlineAddPrepayParam *TopOfflineAddPrepayDto) error {
	r._offlineAddPrepayParam = _offlineAddPrepayParam
	r.Set("offline_add_prepay_param", _offlineAddPrepayParam)
	return nil
}

// GetOfflineAddPrepayParam OfflineAddPrepayParam Getter
func (r TaobaoChannelTradePrepayOfflineAddAPIRequest) GetOfflineAddPrepayParam() *TopOfflineAddPrepayDto {
	return r._offlineAddPrepayParam
}

var poolTaobaoChannelTradePrepayOfflineAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoChannelTradePrepayOfflineAddRequest()
	},
}

// GetTaobaoChannelTradePrepayOfflineAddRequest 从 sync.Pool 获取 TaobaoChannelTradePrepayOfflineAddAPIRequest
func GetTaobaoChannelTradePrepayOfflineAddAPIRequest() *TaobaoChannelTradePrepayOfflineAddAPIRequest {
	return poolTaobaoChannelTradePrepayOfflineAddAPIRequest.Get().(*TaobaoChannelTradePrepayOfflineAddAPIRequest)
}

// ReleaseTaobaoChannelTradePrepayOfflineAddAPIRequest 将 TaobaoChannelTradePrepayOfflineAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoChannelTradePrepayOfflineAddAPIRequest(v *TaobaoChannelTradePrepayOfflineAddAPIRequest) {
	v.Reset()
	poolTaobaoChannelTradePrepayOfflineAddAPIRequest.Put(v)
}

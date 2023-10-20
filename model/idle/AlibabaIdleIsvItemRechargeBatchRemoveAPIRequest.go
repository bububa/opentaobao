package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleisvitemrechargebatchremoveAPIRequest 闲鱼商品直充功能移除 API请求
// alibaba.idle.isv.item.recharge.batch.remove
//
// 闲鱼商品直充功能移除
type AlibabaidleisvitemrechargebatchremoveAPIRequest struct {
	model.Params
	// 请求参数
	_idleItemAutoRechargeBatchRemoveApiDo *IdleItemAutoRechargeBatchRemoveApiDo
}

// NewAlibabaidleisvitemrechargebatchremoveRequest 初始化AlibabaidleisvitemrechargebatchremoveAPIRequest对象
func NewAlibabaidleisvitemrechargebatchremoveRequest() *AlibabaidleisvitemrechargebatchremoveAPIRequest {
	return &AlibabaidleisvitemrechargebatchremoveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleisvitemrechargebatchremoveAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.item.recharge.batch.remove"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleisvitemrechargebatchremoveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleisvitemrechargebatchremoveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIdleItemAutoRechargeBatchRemoveApiDo is IdleItemAutoRechargeBatchRemoveApiDo Setter
// 请求参数
func (r *AlibabaidleisvitemrechargebatchremoveAPIRequest) SetIdleItemAutoRechargeBatchRemoveApiDo(_idleItemAutoRechargeBatchRemoveApiDo *IdleItemAutoRechargeBatchRemoveApiDo) error {
	r._idleItemAutoRechargeBatchRemoveApiDo = _idleItemAutoRechargeBatchRemoveApiDo
	r.Set("idle_item_auto_recharge_batch_remove_api_do", _idleItemAutoRechargeBatchRemoveApiDo)
	return nil
}

// GetIdleItemAutoRechargeBatchRemoveApiDo IdleItemAutoRechargeBatchRemoveApiDo Getter
func (r AlibabaidleisvitemrechargebatchremoveAPIRequest) GetIdleItemAutoRechargeBatchRemoveApiDo() *IdleItemAutoRechargeBatchRemoveApiDo {
	return r._idleItemAutoRechargeBatchRemoveApiDo
}

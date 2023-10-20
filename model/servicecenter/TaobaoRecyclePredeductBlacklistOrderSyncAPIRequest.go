package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRecyclePredeductBlacklistOrderSyncAPIRequest 同步服务商黑名单 API请求
// taobao.recycle.prededuct.blacklist.order.sync
//
// 同步服务商黑名单
type TaobaoRecyclePredeductBlacklistOrderSyncAPIRequest struct {
	model.Params
	// 需要加入黑名单的回收单订单 ID
	_blackOrderId int64
}

// NewTaobaoRecyclePredeductBlacklistOrderSyncRequest 初始化TaobaoRecyclePredeductBlacklistOrderSyncAPIRequest对象
func NewTaobaoRecyclePredeductBlacklistOrderSyncRequest() *TaobaoRecyclePredeductBlacklistOrderSyncAPIRequest {
	return &TaobaoRecyclePredeductBlacklistOrderSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRecyclePredeductBlacklistOrderSyncAPIRequest) GetApiMethodName() string {
	return "taobao.recycle.prededuct.blacklist.order.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRecyclePredeductBlacklistOrderSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRecyclePredeductBlacklistOrderSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBlackOrderId is BlackOrderId Setter
// 需要加入黑名单的回收单订单 ID
func (r *TaobaoRecyclePredeductBlacklistOrderSyncAPIRequest) SetBlackOrderId(_blackOrderId int64) error {
	r._blackOrderId = _blackOrderId
	r.Set("black_order_id", _blackOrderId)
	return nil
}

// GetBlackOrderId BlackOrderId Getter
func (r TaobaoRecyclePredeductBlacklistOrderSyncAPIRequest) GetBlackOrderId() int64 {
	return r._blackOrderId
}

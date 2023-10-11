package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRecyclePredeductOldGetAPIRequest 查询回收单前置抵扣详情 API请求
// taobao.recycle.prededuct.old.get
//
// 查询回收单前置抵扣详情
type TaobaoRecyclePredeductOldGetAPIRequest struct {
	model.Params
	// 旧机单 ID
	_oldOrderId int64
}

// NewTaobaoRecyclePredeductOldGetRequest 初始化TaobaoRecyclePredeductOldGetAPIRequest对象
func NewTaobaoRecyclePredeductOldGetRequest() *TaobaoRecyclePredeductOldGetAPIRequest {
	return &TaobaoRecyclePredeductOldGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRecyclePredeductOldGetAPIRequest) GetApiMethodName() string {
	return "taobao.recycle.prededuct.old.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRecyclePredeductOldGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRecyclePredeductOldGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOldOrderId is OldOrderId Setter
// 旧机单 ID
func (r *TaobaoRecyclePredeductOldGetAPIRequest) SetOldOrderId(_oldOrderId int64) error {
	r._oldOrderId = _oldOrderId
	r.Set("old_order_id", _oldOrderId)
	return nil
}

// GetOldOrderId OldOrderId Getter
func (r TaobaoRecyclePredeductOldGetAPIRequest) GetOldOrderId() int64 {
	return r._oldOrderId
}

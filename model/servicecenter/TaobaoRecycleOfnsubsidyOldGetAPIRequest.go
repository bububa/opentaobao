package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRecycleOfnsubsidyOldGetAPIRequest 回收单旧机款及补贴查询 API请求
// taobao.recycle.ofnsubsidy.old.get
//
// 回收单旧机款及补贴查询
type TaobaoRecycleOfnsubsidyOldGetAPIRequest struct {
	model.Params
	// 旧机单 ID
	_oldOrderId int64
}

// NewTaobaoRecycleOfnsubsidyOldGetRequest 初始化TaobaoRecycleOfnsubsidyOldGetAPIRequest对象
func NewTaobaoRecycleOfnsubsidyOldGetRequest() *TaobaoRecycleOfnsubsidyOldGetAPIRequest {
	return &TaobaoRecycleOfnsubsidyOldGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRecycleOfnsubsidyOldGetAPIRequest) GetApiMethodName() string {
	return "taobao.recycle.ofnsubsidy.old.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRecycleOfnsubsidyOldGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRecycleOfnsubsidyOldGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOldOrderId is OldOrderId Setter
// 旧机单 ID
func (r *TaobaoRecycleOfnsubsidyOldGetAPIRequest) SetOldOrderId(_oldOrderId int64) error {
	r._oldOrderId = _oldOrderId
	r.Set("old_order_id", _oldOrderId)
	return nil
}

// GetOldOrderId OldOrderId Getter
func (r TaobaoRecycleOfnsubsidyOldGetAPIRequest) GetOldOrderId() int64 {
	return r._oldOrderId
}

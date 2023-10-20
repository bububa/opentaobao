package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaorecycleofnsubsidyoldgetAPIRequest 回收单旧机款及补贴查询 API请求
// taobao.recycle.ofnsubsidy.old.get
//
// 回收单旧机款及补贴查询
type TaobaorecycleofnsubsidyoldgetAPIRequest struct {
	model.Params
	// 旧机单 ID
	_oldOrderId int64
}

// NewTaobaorecycleofnsubsidyoldgetRequest 初始化TaobaorecycleofnsubsidyoldgetAPIRequest对象
func NewTaobaorecycleofnsubsidyoldgetRequest() *TaobaorecycleofnsubsidyoldgetAPIRequest {
	return &TaobaorecycleofnsubsidyoldgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaorecycleofnsubsidyoldgetAPIRequest) GetApiMethodName() string {
	return "taobao.recycle.ofnsubsidy.old.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaorecycleofnsubsidyoldgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaorecycleofnsubsidyoldgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOldOrderId is OldOrderId Setter
// 旧机单 ID
func (r *TaobaorecycleofnsubsidyoldgetAPIRequest) SetOldOrderId(_oldOrderId int64) error {
	r._oldOrderId = _oldOrderId
	r.Set("old_order_id", _oldOrderId)
	return nil
}

// GetOldOrderId OldOrderId Getter
func (r TaobaorecycleofnsubsidyoldgetAPIRequest) GetOldOrderId() int64 {
	return r._oldOrderId
}

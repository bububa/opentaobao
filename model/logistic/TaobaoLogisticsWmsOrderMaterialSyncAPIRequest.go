package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticswmsordermaterialsyncAPIRequest 仓服务商订单包材耗材信息同步 API请求
// taobao.logistics.wms.order.material.sync
//
// 仓服务商订单包材耗材信息同步
type TaobaologisticswmsordermaterialsyncAPIRequest struct {
	model.Params
	// 包材耗材信息
	_wmsMaterialRequest *WmsMaterialRequest
}

// NewTaobaologisticswmsordermaterialsyncRequest 初始化TaobaologisticswmsordermaterialsyncAPIRequest对象
func NewTaobaologisticswmsordermaterialsyncRequest() *TaobaologisticswmsordermaterialsyncAPIRequest {
	return &TaobaologisticswmsordermaterialsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticswmsordermaterialsyncAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.wms.order.material.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticswmsordermaterialsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticswmsordermaterialsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWmsMaterialRequest is WmsMaterialRequest Setter
// 包材耗材信息
func (r *TaobaologisticswmsordermaterialsyncAPIRequest) SetWmsMaterialRequest(_wmsMaterialRequest *WmsMaterialRequest) error {
	r._wmsMaterialRequest = _wmsMaterialRequest
	r.Set("wms_material_request", _wmsMaterialRequest)
	return nil
}

// GetWmsMaterialRequest WmsMaterialRequest Getter
func (r TaobaologisticswmsordermaterialsyncAPIRequest) GetWmsMaterialRequest() *WmsMaterialRequest {
	return r._wmsMaterialRequest
}

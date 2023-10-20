package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsWmsOrderMaterialSyncAPIRequest 仓服务商订单包材耗材信息同步 API请求
// taobao.logistics.wms.order.material.sync
//
// 仓服务商订单包材耗材信息同步
type TaobaoLogisticsWmsOrderMaterialSyncAPIRequest struct {
	model.Params
	// 包材耗材信息
	_wmsMaterialRequest *WmsMaterialRequest
}

// NewTaobaoLogisticsWmsOrderMaterialSyncRequest 初始化TaobaoLogisticsWmsOrderMaterialSyncAPIRequest对象
func NewTaobaoLogisticsWmsOrderMaterialSyncRequest() *TaobaoLogisticsWmsOrderMaterialSyncAPIRequest {
	return &TaobaoLogisticsWmsOrderMaterialSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsWmsOrderMaterialSyncAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.wms.order.material.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsWmsOrderMaterialSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsWmsOrderMaterialSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWmsMaterialRequest is WmsMaterialRequest Setter
// 包材耗材信息
func (r *TaobaoLogisticsWmsOrderMaterialSyncAPIRequest) SetWmsMaterialRequest(_wmsMaterialRequest *WmsMaterialRequest) error {
	r._wmsMaterialRequest = _wmsMaterialRequest
	r.Set("wms_material_request", _wmsMaterialRequest)
	return nil
}

// GetWmsMaterialRequest WmsMaterialRequest Getter
func (r TaobaoLogisticsWmsOrderMaterialSyncAPIRequest) GetWmsMaterialRequest() *WmsMaterialRequest {
	return r._wmsMaterialRequest
}

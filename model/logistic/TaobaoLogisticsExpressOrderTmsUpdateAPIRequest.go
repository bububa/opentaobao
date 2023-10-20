package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsexpressordertmsupdateAPIRequest 服务商修改上门取退时间接口 API请求
// taobao.logistics.express.order.tms.update
//
// 服务商修改上门取退时间接口
type TaobaologisticsexpressordertmsupdateAPIRequest struct {
	model.Params
	// tms更新参数
	_tmsToMscUpdateOrderRequest *Tms2mscUpdateOrderRequest
}

// NewTaobaologisticsexpressordertmsupdateRequest 初始化TaobaologisticsexpressordertmsupdateAPIRequest对象
func NewTaobaologisticsexpressordertmsupdateRequest() *TaobaologisticsexpressordertmsupdateAPIRequest {
	return &TaobaologisticsexpressordertmsupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsexpressordertmsupdateAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.order.tms.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsexpressordertmsupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsexpressordertmsupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTmsToMscUpdateOrderRequest is TmsToMscUpdateOrderRequest Setter
// tms更新参数
func (r *TaobaologisticsexpressordertmsupdateAPIRequest) SetTmsToMscUpdateOrderRequest(_tmsToMscUpdateOrderRequest *Tms2mscUpdateOrderRequest) error {
	r._tmsToMscUpdateOrderRequest = _tmsToMscUpdateOrderRequest
	r.Set("tms_to_msc_update_order_request", _tmsToMscUpdateOrderRequest)
	return nil
}

// GetTmsToMscUpdateOrderRequest TmsToMscUpdateOrderRequest Getter
func (r TaobaologisticsexpressordertmsupdateAPIRequest) GetTmsToMscUpdateOrderRequest() *Tms2mscUpdateOrderRequest {
	return r._tmsToMscUpdateOrderRequest
}

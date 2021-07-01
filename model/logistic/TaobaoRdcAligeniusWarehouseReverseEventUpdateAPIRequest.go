package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest
销退单事件回传接口 API请求
taobao.rdc.aligenius.warehouse.reverse.event.update

用于erp回传销退单相关信息到平台 */
type TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest struct {
	model.Params
	// 参数
	_param0 *ReverseEventInfoDto
}

// NewTaobaoRdcAligeniusWarehouseReverseEventUpdateRequest 初始化TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest对象
func NewTaobaoRdcAligeniusWarehouseReverseEventUpdateRequest() *TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest {
	return &TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.rdc.aligenius.warehouse.reverse.event.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 参数
func (r *TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest) SetParam0(_param0 *ReverseEventInfoDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest) GetParam0() *ReverseEventInfoDto {
	return r._param0
}

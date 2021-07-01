package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest
补发单状态回传 API请求
taobao.rdc.aligenius.warehouse.resend.update

补发单状态回传接口 */
type TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest struct {
	model.Params
	// 参数
	_param0 *UpdateResendStatusDto
}

// NewTaobaoRdcAligeniusWarehouseResendUpdateRequest 初始化TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest对象
func NewTaobaoRdcAligeniusWarehouseResendUpdateRequest() *TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest {
	return &TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.rdc.aligenius.warehouse.resend.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 参数
func (r *TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest) SetParam0(_param0 *UpdateResendStatusDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest) GetParam0() *UpdateResendStatusDto {
	return r._param0
}

package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest 补发单状态回传 API请求
// taobao.rdc.aligenius.warehouse.resend.update
//
// 补发单状态回传接口
type TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest struct {
	model.Params
	// 参数
	_param0 *UpdateResendStatusDto
}

// NewTaobaoRdcAligeniusWarehouseResendUpdateRequest 初始化TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest对象
func NewTaobaoRdcAligeniusWarehouseResendUpdateRequest() *TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest {
	return &TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.rdc.aligenius.warehouse.resend.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 参数
func (r *TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest) SetParam0(_param0 *UpdateResendStatusDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest) GetParam0() *UpdateResendStatusDto {
	return r._param0
}

var poolTaobaoRdcAligeniusWarehouseResendUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoRdcAligeniusWarehouseResendUpdateRequest()
	},
}

// GetTaobaoRdcAligeniusWarehouseResendUpdateRequest 从 sync.Pool 获取 TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest
func GetTaobaoRdcAligeniusWarehouseResendUpdateAPIRequest() *TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest {
	return poolTaobaoRdcAligeniusWarehouseResendUpdateAPIRequest.Get().(*TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest)
}

// ReleaseTaobaoRdcAligeniusWarehouseResendUpdateAPIRequest 将 TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoRdcAligeniusWarehouseResendUpdateAPIRequest(v *TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest) {
	v.Reset()
	poolTaobaoRdcAligeniusWarehouseResendUpdateAPIRequest.Put(v)
}

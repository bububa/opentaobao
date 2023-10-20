package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest 销退单事件回传接口 API请求
// taobao.rdc.aligenius.warehouse.reverse.event.update
//
// 用于erp回传销退单相关信息到平台
type TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest struct {
	model.Params
	// 参数
	_param0 *ReverseEventInfoDto
}

// NewTaobaoRdcAligeniusWarehouseReverseEventUpdateRequest 初始化TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest对象
func NewTaobaoRdcAligeniusWarehouseReverseEventUpdateRequest() *TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest {
	return &TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.rdc.aligenius.warehouse.reverse.event.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 参数
func (r *TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest) SetParam0(_param0 *ReverseEventInfoDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest) GetParam0() *ReverseEventInfoDto {
	return r._param0
}

var poolTaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoRdcAligeniusWarehouseReverseEventUpdateRequest()
	},
}

// GetTaobaoRdcAligeniusWarehouseReverseEventUpdateRequest 从 sync.Pool 获取 TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest
func GetTaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest() *TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest {
	return poolTaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest.Get().(*TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest)
}

// ReleaseTaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest 将 TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest(v *TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest) {
	v.Reset()
	poolTaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest.Put(v)
}

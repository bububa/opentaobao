package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaordcaligeniuswarehousereverseeventupdateAPIRequest 销退单事件回传接口 API请求
// taobao.rdc.aligenius.warehouse.reverse.event.update
//
// 用于erp回传销退单相关信息到平台
type TaobaordcaligeniuswarehousereverseeventupdateAPIRequest struct {
	model.Params
	// 参数
	_param0 *ReverseEventInfoDto
}

// NewTaobaordcaligeniuswarehousereverseeventupdateRequest 初始化TaobaordcaligeniuswarehousereverseeventupdateAPIRequest对象
func NewTaobaordcaligeniuswarehousereverseeventupdateRequest() *TaobaordcaligeniuswarehousereverseeventupdateAPIRequest {
	return &TaobaordcaligeniuswarehousereverseeventupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaordcaligeniuswarehousereverseeventupdateAPIRequest) GetApiMethodName() string {
	return "taobao.rdc.aligenius.warehouse.reverse.event.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaordcaligeniuswarehousereverseeventupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaordcaligeniuswarehousereverseeventupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 参数
func (r *TaobaordcaligeniuswarehousereverseeventupdateAPIRequest) SetParam0(_param0 *ReverseEventInfoDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaordcaligeniuswarehousereverseeventupdateAPIRequest) GetParam0() *ReverseEventInfoDto {
	return r._param0
}

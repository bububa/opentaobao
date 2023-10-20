package refund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaordcaligeniussendgoodscancelAPIRequest 取消发货 API请求
// taobao.rdc.aligenius.sendgoods.cancel
//
// 提供商家在仅退款中发送取消发货状态
type TaobaordcaligeniussendgoodscancelAPIRequest struct {
	model.Params
	// 请求参数
	_param *CancelGoodsDto
}

// NewTaobaordcaligeniussendgoodscancelRequest 初始化TaobaordcaligeniussendgoodscancelAPIRequest对象
func NewTaobaordcaligeniussendgoodscancelRequest() *TaobaordcaligeniussendgoodscancelAPIRequest {
	return &TaobaordcaligeniussendgoodscancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaordcaligeniussendgoodscancelAPIRequest) GetApiMethodName() string {
	return "taobao.rdc.aligenius.sendgoods.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaordcaligeniussendgoodscancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaordcaligeniussendgoodscancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 请求参数
func (r *TaobaordcaligeniussendgoodscancelAPIRequest) SetParam(_param *CancelGoodsDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaordcaligeniussendgoodscancelAPIRequest) GetParam() *CancelGoodsDto {
	return r._param
}

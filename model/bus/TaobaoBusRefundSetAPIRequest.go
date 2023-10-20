package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobusrefundsetAPIRequest B2B退票申请接口 API请求
// taobao.bus.refund.set
//
// B2B业务支持退票
type TaobaobusrefundsetAPIRequest struct {
	model.Params
	// 入参
	_param0 *B2brefundOrderRq
}

// NewTaobaobusrefundsetRequest 初始化TaobaobusrefundsetAPIRequest对象
func NewTaobaobusrefundsetRequest() *TaobaobusrefundsetAPIRequest {
	return &TaobaobusrefundsetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobusrefundsetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.refund.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobusrefundsetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobusrefundsetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参
func (r *TaobaobusrefundsetAPIRequest) SetParam0(_param0 *B2brefundOrderRq) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaobusrefundsetAPIRequest) GetParam0() *B2brefundOrderRq {
	return r._param0
}

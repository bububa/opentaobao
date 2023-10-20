package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelorderupdateconfirmcodeAPIRequest 推送及更新订单确认号 API请求
// taobao.xhotel.order.update.confirmcode
//
// 商家拿到订单确认号后，异步推送给飞猪或更新给飞猪。订单确认号用于到店查无时的紧急查单依据。
type TaobaoxhotelorderupdateconfirmcodeAPIRequest struct {
	model.Params
	// 系统入参
	_param *UpdateOrderConfirmCodeParam
}

// NewTaobaoxhotelorderupdateconfirmcodeRequest 初始化TaobaoxhotelorderupdateconfirmcodeAPIRequest对象
func NewTaobaoxhotelorderupdateconfirmcodeRequest() *TaobaoxhotelorderupdateconfirmcodeAPIRequest {
	return &TaobaoxhotelorderupdateconfirmcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelorderupdateconfirmcodeAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.order.update.confirmcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelorderupdateconfirmcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelorderupdateconfirmcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 系统入参
func (r *TaobaoxhotelorderupdateconfirmcodeAPIRequest) SetParam(_param *UpdateOrderConfirmCodeParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoxhotelorderupdateconfirmcodeAPIRequest) GetParam() *UpdateOrderConfirmCodeParam {
	return r._param
}

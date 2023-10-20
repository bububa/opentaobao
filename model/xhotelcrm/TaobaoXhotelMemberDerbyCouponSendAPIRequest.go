package xhotelcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelmemberderbycouponsendAPIRequest 发券 API请求
// taobao.xhotel.member.derby.coupon.send
//
// 发券
type TaobaoxhotelmemberderbycouponsendAPIRequest struct {
	model.Params
	// 入参模型
	_param *SendCouponParam
}

// NewTaobaoxhotelmemberderbycouponsendRequest 初始化TaobaoxhotelmemberderbycouponsendAPIRequest对象
func NewTaobaoxhotelmemberderbycouponsendRequest() *TaobaoxhotelmemberderbycouponsendAPIRequest {
	return &TaobaoxhotelmemberderbycouponsendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelmemberderbycouponsendAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.member.derby.coupon.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelmemberderbycouponsendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelmemberderbycouponsendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参模型
func (r *TaobaoxhotelmemberderbycouponsendAPIRequest) SetParam(_param *SendCouponParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoxhotelmemberderbycouponsendAPIRequest) GetParam() *SendCouponParam {
	return r._param
}

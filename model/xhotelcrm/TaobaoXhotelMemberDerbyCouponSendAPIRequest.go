package xhotelcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelMemberDerbyCouponSendAPIRequest 发券 API请求
// taobao.xhotel.member.derby.coupon.send
//
// 发券
type TaobaoXhotelMemberDerbyCouponSendAPIRequest struct {
	model.Params
	// 入参模型
	_param *SendCouponParam
}

// NewTaobaoXhotelMemberDerbyCouponSendRequest 初始化TaobaoXhotelMemberDerbyCouponSendAPIRequest对象
func NewTaobaoXhotelMemberDerbyCouponSendRequest() *TaobaoXhotelMemberDerbyCouponSendAPIRequest {
	return &TaobaoXhotelMemberDerbyCouponSendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelMemberDerbyCouponSendAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.member.derby.coupon.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelMemberDerbyCouponSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelMemberDerbyCouponSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参模型
func (r *TaobaoXhotelMemberDerbyCouponSendAPIRequest) SetParam(_param *SendCouponParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoXhotelMemberDerbyCouponSendAPIRequest) GetParam() *SendCouponParam {
	return r._param
}

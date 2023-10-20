package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtcouponinstancequeryAPIRequest 查询用户券实例 API请求
// tmall.nrt.coupon.instance.query
//
// 查询用户券实例
type TmallnrtcouponinstancequeryAPIRequest struct {
	model.Params
	// 券实例id
	_instanceIds string
	// 业务身份
	_bizCode string
	// 券类型
	_couponTypes string
	// 加密后的uid
	_userId string
}

// NewTmallnrtcouponinstancequeryRequest 初始化TmallnrtcouponinstancequeryAPIRequest对象
func NewTmallnrtcouponinstancequeryRequest() *TmallnrtcouponinstancequeryAPIRequest {
	return &TmallnrtcouponinstancequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrtcouponinstancequeryAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.coupon.instance.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrtcouponinstancequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrtcouponinstancequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInstanceIds is InstanceIds Setter
// 券实例id
func (r *TmallnrtcouponinstancequeryAPIRequest) SetInstanceIds(_instanceIds string) error {
	r._instanceIds = _instanceIds
	r.Set("instance_ids", _instanceIds)
	return nil
}

// GetInstanceIds InstanceIds Getter
func (r TmallnrtcouponinstancequeryAPIRequest) GetInstanceIds() string {
	return r._instanceIds
}

// SetBizCode is BizCode Setter
// 业务身份
func (r *TmallnrtcouponinstancequeryAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// GetBizCode BizCode Getter
func (r TmallnrtcouponinstancequeryAPIRequest) GetBizCode() string {
	return r._bizCode
}

// SetCouponTypes is CouponTypes Setter
// 券类型
func (r *TmallnrtcouponinstancequeryAPIRequest) SetCouponTypes(_couponTypes string) error {
	r._couponTypes = _couponTypes
	r.Set("coupon_types", _couponTypes)
	return nil
}

// GetCouponTypes CouponTypes Getter
func (r TmallnrtcouponinstancequeryAPIRequest) GetCouponTypes() string {
	return r._couponTypes
}

// SetUserId is UserId Setter
// 加密后的uid
func (r *TmallnrtcouponinstancequeryAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TmallnrtcouponinstancequeryAPIRequest) GetUserId() string {
	return r._userId
}

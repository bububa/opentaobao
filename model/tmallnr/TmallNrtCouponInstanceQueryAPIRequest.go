package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtCouponInstanceQueryAPIRequest 查询用户券实例 API请求
// tmall.nrt.coupon.instance.query
//
// 查询用户券实例
type TmallNrtCouponInstanceQueryAPIRequest struct {
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

// NewTmallNrtCouponInstanceQueryRequest 初始化TmallNrtCouponInstanceQueryAPIRequest对象
func NewTmallNrtCouponInstanceQueryRequest() *TmallNrtCouponInstanceQueryAPIRequest {
	return &TmallNrtCouponInstanceQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtCouponInstanceQueryAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.coupon.instance.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtCouponInstanceQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetInstanceIds is InstanceIds Setter
// 券实例id
func (r *TmallNrtCouponInstanceQueryAPIRequest) SetInstanceIds(_instanceIds string) error {
	r._instanceIds = _instanceIds
	r.Set("instance_ids", _instanceIds)
	return nil
}

// GetInstanceIds InstanceIds Getter
func (r TmallNrtCouponInstanceQueryAPIRequest) GetInstanceIds() string {
	return r._instanceIds
}

// SetBizCode is BizCode Setter
// 业务身份
func (r *TmallNrtCouponInstanceQueryAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// GetBizCode BizCode Getter
func (r TmallNrtCouponInstanceQueryAPIRequest) GetBizCode() string {
	return r._bizCode
}

// SetCouponTypes is CouponTypes Setter
// 券类型
func (r *TmallNrtCouponInstanceQueryAPIRequest) SetCouponTypes(_couponTypes string) error {
	r._couponTypes = _couponTypes
	r.Set("coupon_types", _couponTypes)
	return nil
}

// GetCouponTypes CouponTypes Getter
func (r TmallNrtCouponInstanceQueryAPIRequest) GetCouponTypes() string {
	return r._couponTypes
}

// SetUserId is UserId Setter
// 加密后的uid
func (r *TmallNrtCouponInstanceQueryAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TmallNrtCouponInstanceQueryAPIRequest) GetUserId() string {
	return r._userId
}

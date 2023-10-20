package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtcoupontemplatequeryAPIRequest 查询券模版 API请求
// tmall.nrt.coupon.template.query
//
// 查询券模版
type TmallnrtcoupontemplatequeryAPIRequest struct {
	model.Params
	// 业务编码
	_bizCode string
	// 券模版id集合
	_couponTemplateIds string
	// 券类型
	_couponType int64
}

// NewTmallnrtcoupontemplatequeryRequest 初始化TmallnrtcoupontemplatequeryAPIRequest对象
func NewTmallnrtcoupontemplatequeryRequest() *TmallnrtcoupontemplatequeryAPIRequest {
	return &TmallnrtcoupontemplatequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrtcoupontemplatequeryAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.coupon.template.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrtcoupontemplatequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrtcoupontemplatequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizCode is BizCode Setter
// 业务编码
func (r *TmallnrtcoupontemplatequeryAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// GetBizCode BizCode Getter
func (r TmallnrtcoupontemplatequeryAPIRequest) GetBizCode() string {
	return r._bizCode
}

// SetCouponTemplateIds is CouponTemplateIds Setter
// 券模版id集合
func (r *TmallnrtcoupontemplatequeryAPIRequest) SetCouponTemplateIds(_couponTemplateIds string) error {
	r._couponTemplateIds = _couponTemplateIds
	r.Set("coupon_template_ids", _couponTemplateIds)
	return nil
}

// GetCouponTemplateIds CouponTemplateIds Getter
func (r TmallnrtcoupontemplatequeryAPIRequest) GetCouponTemplateIds() string {
	return r._couponTemplateIds
}

// SetCouponType is CouponType Setter
// 券类型
func (r *TmallnrtcoupontemplatequeryAPIRequest) SetCouponType(_couponType int64) error {
	r._couponType = _couponType
	r.Set("coupon_type", _couponType)
	return nil
}

// GetCouponType CouponType Getter
func (r TmallnrtcoupontemplatequeryAPIRequest) GetCouponType() int64 {
	return r._couponType
}

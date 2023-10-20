package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtcoupontemplatesynAPIRequest 喵零券同步 API请求
// tmall.nrt.coupon.template.syn
//
// 查询券模版
type TmallnrtcoupontemplatesynAPIRequest struct {
	model.Params
	// 业务编码
	_bizCode string
	// 券模版id集合
	_couponTemplateId string
	// 券类型
	_couponType int64
}

// NewTmallnrtcoupontemplatesynRequest 初始化TmallnrtcoupontemplatesynAPIRequest对象
func NewTmallnrtcoupontemplatesynRequest() *TmallnrtcoupontemplatesynAPIRequest {
	return &TmallnrtcoupontemplatesynAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrtcoupontemplatesynAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.coupon.template.syn"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrtcoupontemplatesynAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrtcoupontemplatesynAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizCode is BizCode Setter
// 业务编码
func (r *TmallnrtcoupontemplatesynAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// GetBizCode BizCode Getter
func (r TmallnrtcoupontemplatesynAPIRequest) GetBizCode() string {
	return r._bizCode
}

// SetCouponTemplateId is CouponTemplateId Setter
// 券模版id集合
func (r *TmallnrtcoupontemplatesynAPIRequest) SetCouponTemplateId(_couponTemplateId string) error {
	r._couponTemplateId = _couponTemplateId
	r.Set("coupon_template_id", _couponTemplateId)
	return nil
}

// GetCouponTemplateId CouponTemplateId Getter
func (r TmallnrtcoupontemplatesynAPIRequest) GetCouponTemplateId() string {
	return r._couponTemplateId
}

// SetCouponType is CouponType Setter
// 券类型
func (r *TmallnrtcoupontemplatesynAPIRequest) SetCouponType(_couponType int64) error {
	r._couponType = _couponType
	r.Set("coupon_type", _couponType)
	return nil
}

// GetCouponType CouponType Getter
func (r TmallnrtcoupontemplatesynAPIRequest) GetCouponType() int64 {
	return r._couponType
}

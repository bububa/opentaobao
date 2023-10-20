package nrt

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtCouponTemplateSynAPIRequest 喵零券同步 API请求
// tmall.nrt.coupon.template.syn
//
// 查询券模版
type TmallNrtCouponTemplateSynAPIRequest struct {
	model.Params
	// 业务编码
	_bizCode string
	// 券模版id集合
	_couponTemplateId string
	// 券类型
	_couponType int64
}

// NewTmallNrtCouponTemplateSynRequest 初始化TmallNrtCouponTemplateSynAPIRequest对象
func NewTmallNrtCouponTemplateSynRequest() *TmallNrtCouponTemplateSynAPIRequest {
	return &TmallNrtCouponTemplateSynAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallNrtCouponTemplateSynAPIRequest) Reset() {
	r._bizCode = ""
	r._couponTemplateId = ""
	r._couponType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtCouponTemplateSynAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.coupon.template.syn"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtCouponTemplateSynAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrtCouponTemplateSynAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizCode is BizCode Setter
// 业务编码
func (r *TmallNrtCouponTemplateSynAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// GetBizCode BizCode Getter
func (r TmallNrtCouponTemplateSynAPIRequest) GetBizCode() string {
	return r._bizCode
}

// SetCouponTemplateId is CouponTemplateId Setter
// 券模版id集合
func (r *TmallNrtCouponTemplateSynAPIRequest) SetCouponTemplateId(_couponTemplateId string) error {
	r._couponTemplateId = _couponTemplateId
	r.Set("coupon_template_id", _couponTemplateId)
	return nil
}

// GetCouponTemplateId CouponTemplateId Getter
func (r TmallNrtCouponTemplateSynAPIRequest) GetCouponTemplateId() string {
	return r._couponTemplateId
}

// SetCouponType is CouponType Setter
// 券类型
func (r *TmallNrtCouponTemplateSynAPIRequest) SetCouponType(_couponType int64) error {
	r._couponType = _couponType
	r.Set("coupon_type", _couponType)
	return nil
}

// GetCouponType CouponType Getter
func (r TmallNrtCouponTemplateSynAPIRequest) GetCouponType() int64 {
	return r._couponType
}

var poolTmallNrtCouponTemplateSynAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallNrtCouponTemplateSynRequest()
	},
}

// GetTmallNrtCouponTemplateSynRequest 从 sync.Pool 获取 TmallNrtCouponTemplateSynAPIRequest
func GetTmallNrtCouponTemplateSynAPIRequest() *TmallNrtCouponTemplateSynAPIRequest {
	return poolTmallNrtCouponTemplateSynAPIRequest.Get().(*TmallNrtCouponTemplateSynAPIRequest)
}

// ReleaseTmallNrtCouponTemplateSynAPIRequest 将 TmallNrtCouponTemplateSynAPIRequest 放入 sync.Pool
func ReleaseTmallNrtCouponTemplateSynAPIRequest(v *TmallNrtCouponTemplateSynAPIRequest) {
	v.Reset()
	poolTmallNrtCouponTemplateSynAPIRequest.Put(v)
}

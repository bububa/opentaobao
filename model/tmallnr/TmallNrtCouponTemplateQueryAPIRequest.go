package tmallnr

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtCouponTemplateQueryAPIRequest 查询券模版 API请求
// tmall.nrt.coupon.template.query
//
// 查询券模版
type TmallNrtCouponTemplateQueryAPIRequest struct {
	model.Params
	// 业务编码
	_bizCode string
	// 券模版id集合
	_couponTemplateIds string
	// 券类型
	_couponType int64
}

// NewTmallNrtCouponTemplateQueryRequest 初始化TmallNrtCouponTemplateQueryAPIRequest对象
func NewTmallNrtCouponTemplateQueryRequest() *TmallNrtCouponTemplateQueryAPIRequest {
	return &TmallNrtCouponTemplateQueryAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallNrtCouponTemplateQueryAPIRequest) Reset() {
	r._bizCode = ""
	r._couponTemplateIds = ""
	r._couponType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtCouponTemplateQueryAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.coupon.template.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtCouponTemplateQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrtCouponTemplateQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizCode is BizCode Setter
// 业务编码
func (r *TmallNrtCouponTemplateQueryAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// GetBizCode BizCode Getter
func (r TmallNrtCouponTemplateQueryAPIRequest) GetBizCode() string {
	return r._bizCode
}

// SetCouponTemplateIds is CouponTemplateIds Setter
// 券模版id集合
func (r *TmallNrtCouponTemplateQueryAPIRequest) SetCouponTemplateIds(_couponTemplateIds string) error {
	r._couponTemplateIds = _couponTemplateIds
	r.Set("coupon_template_ids", _couponTemplateIds)
	return nil
}

// GetCouponTemplateIds CouponTemplateIds Getter
func (r TmallNrtCouponTemplateQueryAPIRequest) GetCouponTemplateIds() string {
	return r._couponTemplateIds
}

// SetCouponType is CouponType Setter
// 券类型
func (r *TmallNrtCouponTemplateQueryAPIRequest) SetCouponType(_couponType int64) error {
	r._couponType = _couponType
	r.Set("coupon_type", _couponType)
	return nil
}

// GetCouponType CouponType Getter
func (r TmallNrtCouponTemplateQueryAPIRequest) GetCouponType() int64 {
	return r._couponType
}

var poolTmallNrtCouponTemplateQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallNrtCouponTemplateQueryRequest()
	},
}

// GetTmallNrtCouponTemplateQueryRequest 从 sync.Pool 获取 TmallNrtCouponTemplateQueryAPIRequest
func GetTmallNrtCouponTemplateQueryAPIRequest() *TmallNrtCouponTemplateQueryAPIRequest {
	return poolTmallNrtCouponTemplateQueryAPIRequest.Get().(*TmallNrtCouponTemplateQueryAPIRequest)
}

// ReleaseTmallNrtCouponTemplateQueryAPIRequest 将 TmallNrtCouponTemplateQueryAPIRequest 放入 sync.Pool
func ReleaseTmallNrtCouponTemplateQueryAPIRequest(v *TmallNrtCouponTemplateQueryAPIRequest) {
	v.Reset()
	poolTmallNrtCouponTemplateQueryAPIRequest.Put(v)
}

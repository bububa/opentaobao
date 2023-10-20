package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPricePromotionActivityDeleteAPIRequest 删除档期活动 API请求
// alibaba.price.promotion.activity.delete
//
// 删除盒马帮档期活动
type AlibabaPricePromotionActivityDeleteAPIRequest struct {
	model.Params
	// 外部主键
	_outerPromotionCode string
	// 经营店OU
	_ouCode string
}

// NewAlibabaPricePromotionActivityDeleteRequest 初始化AlibabaPricePromotionActivityDeleteAPIRequest对象
func NewAlibabaPricePromotionActivityDeleteRequest() *AlibabaPricePromotionActivityDeleteAPIRequest {
	return &AlibabaPricePromotionActivityDeleteAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaPricePromotionActivityDeleteAPIRequest) Reset() {
	r._outerPromotionCode = ""
	r._ouCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPricePromotionActivityDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.price.promotion.activity.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPricePromotionActivityDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaPricePromotionActivityDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterPromotionCode is OuterPromotionCode Setter
// 外部主键
func (r *AlibabaPricePromotionActivityDeleteAPIRequest) SetOuterPromotionCode(_outerPromotionCode string) error {
	r._outerPromotionCode = _outerPromotionCode
	r.Set("outer_promotion_code", _outerPromotionCode)
	return nil
}

// GetOuterPromotionCode OuterPromotionCode Getter
func (r AlibabaPricePromotionActivityDeleteAPIRequest) GetOuterPromotionCode() string {
	return r._outerPromotionCode
}

// SetOuCode is OuCode Setter
// 经营店OU
func (r *AlibabaPricePromotionActivityDeleteAPIRequest) SetOuCode(_ouCode string) error {
	r._ouCode = _ouCode
	r.Set("ou_code", _ouCode)
	return nil
}

// GetOuCode OuCode Getter
func (r AlibabaPricePromotionActivityDeleteAPIRequest) GetOuCode() string {
	return r._ouCode
}

var poolAlibabaPricePromotionActivityDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaPricePromotionActivityDeleteRequest()
	},
}

// GetAlibabaPricePromotionActivityDeleteRequest 从 sync.Pool 获取 AlibabaPricePromotionActivityDeleteAPIRequest
func GetAlibabaPricePromotionActivityDeleteAPIRequest() *AlibabaPricePromotionActivityDeleteAPIRequest {
	return poolAlibabaPricePromotionActivityDeleteAPIRequest.Get().(*AlibabaPricePromotionActivityDeleteAPIRequest)
}

// ReleaseAlibabaPricePromotionActivityDeleteAPIRequest 将 AlibabaPricePromotionActivityDeleteAPIRequest 放入 sync.Pool
func ReleaseAlibabaPricePromotionActivityDeleteAPIRequest(v *AlibabaPricePromotionActivityDeleteAPIRequest) {
	v.Reset()
	poolAlibabaPricePromotionActivityDeleteAPIRequest.Put(v)
}

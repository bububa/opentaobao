package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPricePromotionItemDeleteAPIRequest 批量删除档期 API请求
// alibaba.price.promotion.item.delete
//
// 盒马帮批量删除档期商品
type AlibabaPricePromotionItemDeleteAPIRequest struct {
	model.Params
	// 商品code
	_skuCodes []string
	// toB渠道店OU
	_ouCode string
	// 外部档期编码
	_outerPromotionCode string
	// 盒马唯一标识
	_uniqueId string
}

// NewAlibabaPricePromotionItemDeleteRequest 初始化AlibabaPricePromotionItemDeleteAPIRequest对象
func NewAlibabaPricePromotionItemDeleteRequest() *AlibabaPricePromotionItemDeleteAPIRequest {
	return &AlibabaPricePromotionItemDeleteAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaPricePromotionItemDeleteAPIRequest) Reset() {
	r._skuCodes = r._skuCodes[:0]
	r._ouCode = ""
	r._outerPromotionCode = ""
	r._uniqueId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPricePromotionItemDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.price.promotion.item.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPricePromotionItemDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaPricePromotionItemDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuCodes is SkuCodes Setter
// 商品code
func (r *AlibabaPricePromotionItemDeleteAPIRequest) SetSkuCodes(_skuCodes []string) error {
	r._skuCodes = _skuCodes
	r.Set("sku_codes", _skuCodes)
	return nil
}

// GetSkuCodes SkuCodes Getter
func (r AlibabaPricePromotionItemDeleteAPIRequest) GetSkuCodes() []string {
	return r._skuCodes
}

// SetOuCode is OuCode Setter
// toB渠道店OU
func (r *AlibabaPricePromotionItemDeleteAPIRequest) SetOuCode(_ouCode string) error {
	r._ouCode = _ouCode
	r.Set("ou_code", _ouCode)
	return nil
}

// GetOuCode OuCode Getter
func (r AlibabaPricePromotionItemDeleteAPIRequest) GetOuCode() string {
	return r._ouCode
}

// SetOuterPromotionCode is OuterPromotionCode Setter
// 外部档期编码
func (r *AlibabaPricePromotionItemDeleteAPIRequest) SetOuterPromotionCode(_outerPromotionCode string) error {
	r._outerPromotionCode = _outerPromotionCode
	r.Set("outer_promotion_code", _outerPromotionCode)
	return nil
}

// GetOuterPromotionCode OuterPromotionCode Getter
func (r AlibabaPricePromotionItemDeleteAPIRequest) GetOuterPromotionCode() string {
	return r._outerPromotionCode
}

// SetUniqueId is UniqueId Setter
// 盒马唯一标识
func (r *AlibabaPricePromotionItemDeleteAPIRequest) SetUniqueId(_uniqueId string) error {
	r._uniqueId = _uniqueId
	r.Set("unique_id", _uniqueId)
	return nil
}

// GetUniqueId UniqueId Getter
func (r AlibabaPricePromotionItemDeleteAPIRequest) GetUniqueId() string {
	return r._uniqueId
}

var poolAlibabaPricePromotionItemDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaPricePromotionItemDeleteRequest()
	},
}

// GetAlibabaPricePromotionItemDeleteRequest 从 sync.Pool 获取 AlibabaPricePromotionItemDeleteAPIRequest
func GetAlibabaPricePromotionItemDeleteAPIRequest() *AlibabaPricePromotionItemDeleteAPIRequest {
	return poolAlibabaPricePromotionItemDeleteAPIRequest.Get().(*AlibabaPricePromotionItemDeleteAPIRequest)
}

// ReleaseAlibabaPricePromotionItemDeleteAPIRequest 将 AlibabaPricePromotionItemDeleteAPIRequest 放入 sync.Pool
func ReleaseAlibabaPricePromotionItemDeleteAPIRequest(v *AlibabaPricePromotionItemDeleteAPIRequest) {
	v.Reset()
	poolAlibabaPricePromotionItemDeleteAPIRequest.Put(v)
}

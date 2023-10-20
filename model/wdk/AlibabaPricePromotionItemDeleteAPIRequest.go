package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabapricepromotionitemdeleteAPIRequest 批量删除档期 API请求
// alibaba.price.promotion.item.delete
//
// 盒马帮批量删除档期商品
type AlibabapricepromotionitemdeleteAPIRequest struct {
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

// NewAlibabapricepromotionitemdeleteRequest 初始化AlibabapricepromotionitemdeleteAPIRequest对象
func NewAlibabapricepromotionitemdeleteRequest() *AlibabapricepromotionitemdeleteAPIRequest {
	return &AlibabapricepromotionitemdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabapricepromotionitemdeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.price.promotion.item.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabapricepromotionitemdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabapricepromotionitemdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuCodes is SkuCodes Setter
// 商品code
func (r *AlibabapricepromotionitemdeleteAPIRequest) SetSkuCodes(_skuCodes []string) error {
	r._skuCodes = _skuCodes
	r.Set("sku_codes", _skuCodes)
	return nil
}

// GetSkuCodes SkuCodes Getter
func (r AlibabapricepromotionitemdeleteAPIRequest) GetSkuCodes() []string {
	return r._skuCodes
}

// SetOuCode is OuCode Setter
// toB渠道店OU
func (r *AlibabapricepromotionitemdeleteAPIRequest) SetOuCode(_ouCode string) error {
	r._ouCode = _ouCode
	r.Set("ou_code", _ouCode)
	return nil
}

// GetOuCode OuCode Getter
func (r AlibabapricepromotionitemdeleteAPIRequest) GetOuCode() string {
	return r._ouCode
}

// SetOuterPromotionCode is OuterPromotionCode Setter
// 外部档期编码
func (r *AlibabapricepromotionitemdeleteAPIRequest) SetOuterPromotionCode(_outerPromotionCode string) error {
	r._outerPromotionCode = _outerPromotionCode
	r.Set("outer_promotion_code", _outerPromotionCode)
	return nil
}

// GetOuterPromotionCode OuterPromotionCode Getter
func (r AlibabapricepromotionitemdeleteAPIRequest) GetOuterPromotionCode() string {
	return r._outerPromotionCode
}

// SetUniqueId is UniqueId Setter
// 盒马唯一标识
func (r *AlibabapricepromotionitemdeleteAPIRequest) SetUniqueId(_uniqueId string) error {
	r._uniqueId = _uniqueId
	r.Set("unique_id", _uniqueId)
	return nil
}

// GetUniqueId UniqueId Getter
func (r AlibabapricepromotionitemdeleteAPIRequest) GetUniqueId() string {
	return r._uniqueId
}

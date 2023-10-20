package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabapricepromotionactivitydeleteAPIRequest 删除档期活动 API请求
// alibaba.price.promotion.activity.delete
//
// 删除盒马帮档期活动
type AlibabapricepromotionactivitydeleteAPIRequest struct {
	model.Params
	// 外部主键
	_outerPromotionCode string
	// 经营店OU
	_ouCode string
}

// NewAlibabapricepromotionactivitydeleteRequest 初始化AlibabapricepromotionactivitydeleteAPIRequest对象
func NewAlibabapricepromotionactivitydeleteRequest() *AlibabapricepromotionactivitydeleteAPIRequest {
	return &AlibabapricepromotionactivitydeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabapricepromotionactivitydeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.price.promotion.activity.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabapricepromotionactivitydeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabapricepromotionactivitydeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterPromotionCode is OuterPromotionCode Setter
// 外部主键
func (r *AlibabapricepromotionactivitydeleteAPIRequest) SetOuterPromotionCode(_outerPromotionCode string) error {
	r._outerPromotionCode = _outerPromotionCode
	r.Set("outer_promotion_code", _outerPromotionCode)
	return nil
}

// GetOuterPromotionCode OuterPromotionCode Getter
func (r AlibabapricepromotionactivitydeleteAPIRequest) GetOuterPromotionCode() string {
	return r._outerPromotionCode
}

// SetOuCode is OuCode Setter
// 经营店OU
func (r *AlibabapricepromotionactivitydeleteAPIRequest) SetOuCode(_ouCode string) error {
	r._ouCode = _ouCode
	r.Set("ou_code", _ouCode)
	return nil
}

// GetOuCode OuCode Getter
func (r AlibabapricepromotionactivitydeleteAPIRequest) GetOuCode() string {
	return r._ouCode
}

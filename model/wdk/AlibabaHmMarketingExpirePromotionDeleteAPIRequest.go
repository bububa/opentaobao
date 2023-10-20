package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingexpirepromotiondeleteAPIRequest 短保优惠删除 API请求
// alibaba.hm.marketing.expire.promotion.delete
//
// 短保优惠删除
type AlibabahmmarketingexpirepromotiondeleteAPIRequest struct {
	model.Params
	// 删除短保优惠
	_param0 *ExpirePromotionBo
}

// NewAlibabahmmarketingexpirepromotiondeleteRequest 初始化AlibabahmmarketingexpirepromotiondeleteAPIRequest对象
func NewAlibabahmmarketingexpirepromotiondeleteRequest() *AlibabahmmarketingexpirepromotiondeleteAPIRequest {
	return &AlibabahmmarketingexpirepromotiondeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingexpirepromotiondeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.expire.promotion.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingexpirepromotiondeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingexpirepromotiondeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 删除短保优惠
func (r *AlibabahmmarketingexpirepromotiondeleteAPIRequest) SetParam0(_param0 *ExpirePromotionBo) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabahmmarketingexpirepromotiondeleteAPIRequest) GetParam0() *ExpirePromotionBo {
	return r._param0
}

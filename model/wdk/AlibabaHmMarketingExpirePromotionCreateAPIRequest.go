package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingexpirepromotioncreateAPIRequest 短保优惠创建 API请求
// alibaba.hm.marketing.expire.promotion.create
//
// 过期优惠优惠信息录入
type AlibabahmmarketingexpirepromotioncreateAPIRequest struct {
	model.Params
	// 创建短保优惠
	_param0 *ExpirePromotionBo
}

// NewAlibabahmmarketingexpirepromotioncreateRequest 初始化AlibabahmmarketingexpirepromotioncreateAPIRequest对象
func NewAlibabahmmarketingexpirepromotioncreateRequest() *AlibabahmmarketingexpirepromotioncreateAPIRequest {
	return &AlibabahmmarketingexpirepromotioncreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingexpirepromotioncreateAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.expire.promotion.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingexpirepromotioncreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingexpirepromotioncreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 创建短保优惠
func (r *AlibabahmmarketingexpirepromotioncreateAPIRequest) SetParam0(_param0 *ExpirePromotionBo) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabahmmarketingexpirepromotioncreateAPIRequest) GetParam0() *ExpirePromotionBo {
	return r._param0
}

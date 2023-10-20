package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingexpirepromotioncreateAPIRequest 短保优惠创建 API请求
// alibaba.wdk.marketing.expire.promotion.create
//
// 过期优惠优惠信息录入
type AlibabawdkmarketingexpirepromotioncreateAPIRequest struct {
	model.Params
	// 创建短保优惠
	_param0 *ExpirePromotionBo
}

// NewAlibabawdkmarketingexpirepromotioncreateRequest 初始化AlibabawdkmarketingexpirepromotioncreateAPIRequest对象
func NewAlibabawdkmarketingexpirepromotioncreateRequest() *AlibabawdkmarketingexpirepromotioncreateAPIRequest {
	return &AlibabawdkmarketingexpirepromotioncreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingexpirepromotioncreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.expire.promotion.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingexpirepromotioncreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingexpirepromotioncreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 创建短保优惠
func (r *AlibabawdkmarketingexpirepromotioncreateAPIRequest) SetParam0(_param0 *ExpirePromotionBo) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabawdkmarketingexpirepromotioncreateAPIRequest) GetParam0() *ExpirePromotionBo {
	return r._param0
}

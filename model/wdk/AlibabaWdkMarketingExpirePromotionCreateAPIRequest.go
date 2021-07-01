package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingExpirePromotionCreateAPIRequest
短保优惠创建 API请求
alibaba.wdk.marketing.expire.promotion.create

过期优惠优惠信息录入 */
type AlibabaWdkMarketingExpirePromotionCreateAPIRequest struct {
	model.Params
	// 创建短保优惠
	_param0 *ExpirePromotionBo
}

// NewAlibabaWdkMarketingExpirePromotionCreateRequest 初始化AlibabaWdkMarketingExpirePromotionCreateAPIRequest对象
func NewAlibabaWdkMarketingExpirePromotionCreateRequest() *AlibabaWdkMarketingExpirePromotionCreateAPIRequest {
	return &AlibabaWdkMarketingExpirePromotionCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingExpirePromotionCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.expire.promotion.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingExpirePromotionCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 创建短保优惠
func (r *AlibabaWdkMarketingExpirePromotionCreateAPIRequest) SetParam0(_param0 *ExpirePromotionBo) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r AlibabaWdkMarketingExpirePromotionCreateAPIRequest) GetParam0() *ExpirePromotionBo {
	return r._param0
}

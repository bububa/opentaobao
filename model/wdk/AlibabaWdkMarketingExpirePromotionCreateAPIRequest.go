package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingExpirePromotionCreateAPIRequest 短保优惠创建 API请求
// alibaba.wdk.marketing.expire.promotion.create
//
// 过期优惠优惠信息录入
type AlibabaWdkMarketingExpirePromotionCreateAPIRequest struct {
	model.Params
	// 创建短保优惠
	_param0 *ExpirePromotionBo
}

// NewAlibabaWdkMarketingExpirePromotionCreateRequest 初始化AlibabaWdkMarketingExpirePromotionCreateAPIRequest对象
func NewAlibabaWdkMarketingExpirePromotionCreateRequest() *AlibabaWdkMarketingExpirePromotionCreateAPIRequest {
	return &AlibabaWdkMarketingExpirePromotionCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingExpirePromotionCreateAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingExpirePromotionCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.expire.promotion.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingExpirePromotionCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingExpirePromotionCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 创建短保优惠
func (r *AlibabaWdkMarketingExpirePromotionCreateAPIRequest) SetParam0(_param0 *ExpirePromotionBo) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaWdkMarketingExpirePromotionCreateAPIRequest) GetParam0() *ExpirePromotionBo {
	return r._param0
}

var poolAlibabaWdkMarketingExpirePromotionCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingExpirePromotionCreateRequest()
	},
}

// GetAlibabaWdkMarketingExpirePromotionCreateRequest 从 sync.Pool 获取 AlibabaWdkMarketingExpirePromotionCreateAPIRequest
func GetAlibabaWdkMarketingExpirePromotionCreateAPIRequest() *AlibabaWdkMarketingExpirePromotionCreateAPIRequest {
	return poolAlibabaWdkMarketingExpirePromotionCreateAPIRequest.Get().(*AlibabaWdkMarketingExpirePromotionCreateAPIRequest)
}

// ReleaseAlibabaWdkMarketingExpirePromotionCreateAPIRequest 将 AlibabaWdkMarketingExpirePromotionCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingExpirePromotionCreateAPIRequest(v *AlibabaWdkMarketingExpirePromotionCreateAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingExpirePromotionCreateAPIRequest.Put(v)
}

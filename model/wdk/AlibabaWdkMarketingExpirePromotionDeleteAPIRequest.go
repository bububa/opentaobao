package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingExpirePromotionDeleteAPIRequest 短保优惠删除 API请求
// alibaba.wdk.marketing.expire.promotion.delete
//
// 短保优惠删除
type AlibabaWdkMarketingExpirePromotionDeleteAPIRequest struct {
	model.Params
	// 删除短保优惠
	_param0 *ExpirePromotionBo
}

// NewAlibabaWdkMarketingExpirePromotionDeleteRequest 初始化AlibabaWdkMarketingExpirePromotionDeleteAPIRequest对象
func NewAlibabaWdkMarketingExpirePromotionDeleteRequest() *AlibabaWdkMarketingExpirePromotionDeleteAPIRequest {
	return &AlibabaWdkMarketingExpirePromotionDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingExpirePromotionDeleteAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingExpirePromotionDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.expire.promotion.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingExpirePromotionDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingExpirePromotionDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 删除短保优惠
func (r *AlibabaWdkMarketingExpirePromotionDeleteAPIRequest) SetParam0(_param0 *ExpirePromotionBo) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaWdkMarketingExpirePromotionDeleteAPIRequest) GetParam0() *ExpirePromotionBo {
	return r._param0
}

var poolAlibabaWdkMarketingExpirePromotionDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingExpirePromotionDeleteRequest()
	},
}

// GetAlibabaWdkMarketingExpirePromotionDeleteRequest 从 sync.Pool 获取 AlibabaWdkMarketingExpirePromotionDeleteAPIRequest
func GetAlibabaWdkMarketingExpirePromotionDeleteAPIRequest() *AlibabaWdkMarketingExpirePromotionDeleteAPIRequest {
	return poolAlibabaWdkMarketingExpirePromotionDeleteAPIRequest.Get().(*AlibabaWdkMarketingExpirePromotionDeleteAPIRequest)
}

// ReleaseAlibabaWdkMarketingExpirePromotionDeleteAPIRequest 将 AlibabaWdkMarketingExpirePromotionDeleteAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingExpirePromotionDeleteAPIRequest(v *AlibabaWdkMarketingExpirePromotionDeleteAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingExpirePromotionDeleteAPIRequest.Put(v)
}

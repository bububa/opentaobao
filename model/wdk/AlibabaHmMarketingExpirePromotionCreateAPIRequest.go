package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingExpirePromotionCreateAPIRequest 短保优惠创建 API请求
// alibaba.hm.marketing.expire.promotion.create
//
// 过期优惠优惠信息录入
type AlibabaHmMarketingExpirePromotionCreateAPIRequest struct {
	model.Params
	// 创建短保优惠
	_param0 *ExpirePromotionBo
}

// NewAlibabaHmMarketingExpirePromotionCreateRequest 初始化AlibabaHmMarketingExpirePromotionCreateAPIRequest对象
func NewAlibabaHmMarketingExpirePromotionCreateRequest() *AlibabaHmMarketingExpirePromotionCreateAPIRequest {
	return &AlibabaHmMarketingExpirePromotionCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHmMarketingExpirePromotionCreateAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingExpirePromotionCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.expire.promotion.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingExpirePromotionCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingExpirePromotionCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 创建短保优惠
func (r *AlibabaHmMarketingExpirePromotionCreateAPIRequest) SetParam0(_param0 *ExpirePromotionBo) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaHmMarketingExpirePromotionCreateAPIRequest) GetParam0() *ExpirePromotionBo {
	return r._param0
}

var poolAlibabaHmMarketingExpirePromotionCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHmMarketingExpirePromotionCreateRequest()
	},
}

// GetAlibabaHmMarketingExpirePromotionCreateRequest 从 sync.Pool 获取 AlibabaHmMarketingExpirePromotionCreateAPIRequest
func GetAlibabaHmMarketingExpirePromotionCreateAPIRequest() *AlibabaHmMarketingExpirePromotionCreateAPIRequest {
	return poolAlibabaHmMarketingExpirePromotionCreateAPIRequest.Get().(*AlibabaHmMarketingExpirePromotionCreateAPIRequest)
}

// ReleaseAlibabaHmMarketingExpirePromotionCreateAPIRequest 将 AlibabaHmMarketingExpirePromotionCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaHmMarketingExpirePromotionCreateAPIRequest(v *AlibabaHmMarketingExpirePromotionCreateAPIRequest) {
	v.Reset()
	poolAlibabaHmMarketingExpirePromotionCreateAPIRequest.Put(v)
}

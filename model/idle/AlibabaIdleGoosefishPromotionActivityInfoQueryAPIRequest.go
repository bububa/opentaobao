package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleGoosefishPromotionActivityInfoQueryAPIRequest 闲鱼三方活动参与信息查询 API请求
// alibaba.idle.goosefish.promotion.activity.info.query
//
// 闲鱼三方活动参与信息查询
type AlibabaIdleGoosefishPromotionActivityInfoQueryAPIRequest struct {
	model.Params
	// 活动参与查询入参
	_promotionActivityQueryParam *PromotionActivityQueryParam
}

// NewAlibabaIdleGoosefishPromotionActivityInfoQueryRequest 初始化AlibabaIdleGoosefishPromotionActivityInfoQueryAPIRequest对象
func NewAlibabaIdleGoosefishPromotionActivityInfoQueryRequest() *AlibabaIdleGoosefishPromotionActivityInfoQueryAPIRequest {
	return &AlibabaIdleGoosefishPromotionActivityInfoQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleGoosefishPromotionActivityInfoQueryAPIRequest) Reset() {
	r._promotionActivityQueryParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleGoosefishPromotionActivityInfoQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.goosefish.promotion.activity.info.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleGoosefishPromotionActivityInfoQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleGoosefishPromotionActivityInfoQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPromotionActivityQueryParam is PromotionActivityQueryParam Setter
// 活动参与查询入参
func (r *AlibabaIdleGoosefishPromotionActivityInfoQueryAPIRequest) SetPromotionActivityQueryParam(_promotionActivityQueryParam *PromotionActivityQueryParam) error {
	r._promotionActivityQueryParam = _promotionActivityQueryParam
	r.Set("promotion_activity_query_param", _promotionActivityQueryParam)
	return nil
}

// GetPromotionActivityQueryParam PromotionActivityQueryParam Getter
func (r AlibabaIdleGoosefishPromotionActivityInfoQueryAPIRequest) GetPromotionActivityQueryParam() *PromotionActivityQueryParam {
	return r._promotionActivityQueryParam
}

var poolAlibabaIdleGoosefishPromotionActivityInfoQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleGoosefishPromotionActivityInfoQueryRequest()
	},
}

// GetAlibabaIdleGoosefishPromotionActivityInfoQueryRequest 从 sync.Pool 获取 AlibabaIdleGoosefishPromotionActivityInfoQueryAPIRequest
func GetAlibabaIdleGoosefishPromotionActivityInfoQueryAPIRequest() *AlibabaIdleGoosefishPromotionActivityInfoQueryAPIRequest {
	return poolAlibabaIdleGoosefishPromotionActivityInfoQueryAPIRequest.Get().(*AlibabaIdleGoosefishPromotionActivityInfoQueryAPIRequest)
}

// ReleaseAlibabaIdleGoosefishPromotionActivityInfoQueryAPIRequest 将 AlibabaIdleGoosefishPromotionActivityInfoQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleGoosefishPromotionActivityInfoQueryAPIRequest(v *AlibabaIdleGoosefishPromotionActivityInfoQueryAPIRequest) {
	v.Reset()
	poolAlibabaIdleGoosefishPromotionActivityInfoQueryAPIRequest.Put(v)
}

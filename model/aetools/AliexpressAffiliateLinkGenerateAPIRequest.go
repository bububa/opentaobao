package aetools

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressaffiliatelinkgenerateAPIRequest 联盟推广链接生成 API请求
// aliexpress.affiliate.link.generate
//
// AE联盟推广链接生成接口
type AliexpressaffiliatelinkgenerateAPIRequest struct {
	model.Params
	// API请求签名
	_appSignature string
	// 原始链接或者值
	_sourceValues string
	// 推广者原始trackingID
	_trackingId string
	// 转换的链接类型：0代表普通Link，1代表Search Link，2代表 hot link
	_promotionLinkType int64
}

// NewAliexpressaffiliatelinkgenerateRequest 初始化AliexpressaffiliatelinkgenerateAPIRequest对象
func NewAliexpressaffiliatelinkgenerateRequest() *AliexpressaffiliatelinkgenerateAPIRequest {
	return &AliexpressaffiliatelinkgenerateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressaffiliatelinkgenerateAPIRequest) GetApiMethodName() string {
	return "aliexpress.affiliate.link.generate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressaffiliatelinkgenerateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressaffiliatelinkgenerateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppSignature is AppSignature Setter
// API请求签名
func (r *AliexpressaffiliatelinkgenerateAPIRequest) SetAppSignature(_appSignature string) error {
	r._appSignature = _appSignature
	r.Set("app_signature", _appSignature)
	return nil
}

// GetAppSignature AppSignature Getter
func (r AliexpressaffiliatelinkgenerateAPIRequest) GetAppSignature() string {
	return r._appSignature
}

// SetSourceValues is SourceValues Setter
// 原始链接或者值
func (r *AliexpressaffiliatelinkgenerateAPIRequest) SetSourceValues(_sourceValues string) error {
	r._sourceValues = _sourceValues
	r.Set("source_values", _sourceValues)
	return nil
}

// GetSourceValues SourceValues Getter
func (r AliexpressaffiliatelinkgenerateAPIRequest) GetSourceValues() string {
	return r._sourceValues
}

// SetTrackingId is TrackingId Setter
// 推广者原始trackingID
func (r *AliexpressaffiliatelinkgenerateAPIRequest) SetTrackingId(_trackingId string) error {
	r._trackingId = _trackingId
	r.Set("tracking_id", _trackingId)
	return nil
}

// GetTrackingId TrackingId Getter
func (r AliexpressaffiliatelinkgenerateAPIRequest) GetTrackingId() string {
	return r._trackingId
}

// SetPromotionLinkType is PromotionLinkType Setter
// 转换的链接类型：0代表普通Link，1代表Search Link，2代表 hot link
func (r *AliexpressaffiliatelinkgenerateAPIRequest) SetPromotionLinkType(_promotionLinkType int64) error {
	r._promotionLinkType = _promotionLinkType
	r.Set("promotion_link_type", _promotionLinkType)
	return nil
}

// GetPromotionLinkType PromotionLinkType Getter
func (r AliexpressaffiliatelinkgenerateAPIRequest) GetPromotionLinkType() int64 {
	return r._promotionLinkType
}

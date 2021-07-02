package aetools

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAffiliateLinkGenerateAPIRequest 联盟推广链接生成 API请求
// aliexpress.affiliate.link.generate
//
// AE联盟推广链接生成接口
type AliexpressAffiliateLinkGenerateAPIRequest struct {
	model.Params
	// API请求签名
	_appSignature string
	// 转换的链接类型：0代表普通Link，1代表Search Link，2代表 hot link
	_promotionLinkType int64
	// 原始链接或者值
	_sourceValues string
	// 推广者原始trackingID
	_trackingId string
}

// NewAliexpressAffiliateLinkGenerateRequest 初始化AliexpressAffiliateLinkGenerateAPIRequest对象
func NewAliexpressAffiliateLinkGenerateRequest() *AliexpressAffiliateLinkGenerateAPIRequest {
	return &AliexpressAffiliateLinkGenerateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressAffiliateLinkGenerateAPIRequest) GetApiMethodName() string {
	return "aliexpress.affiliate.link.generate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressAffiliateLinkGenerateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AppSignature Setter
// API请求签名
func (r *AliexpressAffiliateLinkGenerateAPIRequest) SetAppSignature(_appSignature string) error {
	r._appSignature = _appSignature
	r.Set("app_signature", _appSignature)
	return nil
}

// Get AppSignature Getter
func (r AliexpressAffiliateLinkGenerateAPIRequest) GetAppSignature() string {
	return r._appSignature
}

// Set is PromotionLinkType Setter
// 转换的链接类型：0代表普通Link，1代表Search Link，2代表 hot link
func (r *AliexpressAffiliateLinkGenerateAPIRequest) SetPromotionLinkType(_promotionLinkType int64) error {
	r._promotionLinkType = _promotionLinkType
	r.Set("promotion_link_type", _promotionLinkType)
	return nil
}

// Get PromotionLinkType Getter
func (r AliexpressAffiliateLinkGenerateAPIRequest) GetPromotionLinkType() int64 {
	return r._promotionLinkType
}

// Set is SourceValues Setter
// 原始链接或者值
func (r *AliexpressAffiliateLinkGenerateAPIRequest) SetSourceValues(_sourceValues string) error {
	r._sourceValues = _sourceValues
	r.Set("source_values", _sourceValues)
	return nil
}

// Get SourceValues Getter
func (r AliexpressAffiliateLinkGenerateAPIRequest) GetSourceValues() string {
	return r._sourceValues
}

// Set is TrackingId Setter
// 推广者原始trackingID
func (r *AliexpressAffiliateLinkGenerateAPIRequest) SetTrackingId(_trackingId string) error {
	r._trackingId = _trackingId
	r.Set("tracking_id", _trackingId)
	return nil
}

// Get TrackingId Getter
func (r AliexpressAffiliateLinkGenerateAPIRequest) GetTrackingId() string {
	return r._trackingId
}

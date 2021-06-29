package aetools

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
联盟推广链接生成 API请求
aliexpress.affiliate.link.generate

AE联盟推广链接生成接口
*/
type AliexpressAffiliateLinkGenerateRequest struct {
    model.Params
    // API请求签名
    _appSignature   string
    // 转换的链接类型：0代表普通Link，1代表Search Link，2代表 hot link
    _promotionLinkType   int64
    // 原始链接或者值
    _sourceValues   string
    // 推广者原始trackingID
    _trackingId   string
}

// 初始化AliexpressAffiliateLinkGenerateRequest对象
func NewAliexpressAffiliateLinkGenerateRequest() *AliexpressAffiliateLinkGenerateRequest{
    return &AliexpressAffiliateLinkGenerateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressAffiliateLinkGenerateRequest) GetApiMethodName() string {
    return "aliexpress.affiliate.link.generate"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressAffiliateLinkGenerateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppSignature Setter
// API请求签名
func (r *AliexpressAffiliateLinkGenerateRequest) SetAppSignature(_appSignature string) error {
    r._appSignature = _appSignature
    r.Set("app_signature", _appSignature)
    return nil
}

// AppSignature Getter
func (r AliexpressAffiliateLinkGenerateRequest) GetAppSignature() string {
    return r._appSignature
}
// PromotionLinkType Setter
// 转换的链接类型：0代表普通Link，1代表Search Link，2代表 hot link
func (r *AliexpressAffiliateLinkGenerateRequest) SetPromotionLinkType(_promotionLinkType int64) error {
    r._promotionLinkType = _promotionLinkType
    r.Set("promotion_link_type", _promotionLinkType)
    return nil
}

// PromotionLinkType Getter
func (r AliexpressAffiliateLinkGenerateRequest) GetPromotionLinkType() int64 {
    return r._promotionLinkType
}
// SourceValues Setter
// 原始链接或者值
func (r *AliexpressAffiliateLinkGenerateRequest) SetSourceValues(_sourceValues string) error {
    r._sourceValues = _sourceValues
    r.Set("source_values", _sourceValues)
    return nil
}

// SourceValues Getter
func (r AliexpressAffiliateLinkGenerateRequest) GetSourceValues() string {
    return r._sourceValues
}
// TrackingId Setter
// 推广者原始trackingID
func (r *AliexpressAffiliateLinkGenerateRequest) SetTrackingId(_trackingId string) error {
    r._trackingId = _trackingId
    r.Set("tracking_id", _trackingId)
    return nil
}

// TrackingId Getter
func (r AliexpressAffiliateLinkGenerateRequest) GetTrackingId() string {
    return r._trackingId
}

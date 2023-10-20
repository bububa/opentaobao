package aecreatives

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressaffiliatefeaturedpromogetAPIRequest 联盟主题推广活动信息获取 API请求
// aliexpress.affiliate.featuredpromo.get
//
// 获取联盟主题推广活动信息
type AliexpressaffiliatefeaturedpromogetAPIRequest struct {
	model.Params
	// 请求签名
	_appSignature string
	// 返回字段列表
	_fields string
}

// NewAliexpressaffiliatefeaturedpromogetRequest 初始化AliexpressaffiliatefeaturedpromogetAPIRequest对象
func NewAliexpressaffiliatefeaturedpromogetRequest() *AliexpressaffiliatefeaturedpromogetAPIRequest {
	return &AliexpressaffiliatefeaturedpromogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressaffiliatefeaturedpromogetAPIRequest) GetApiMethodName() string {
	return "aliexpress.affiliate.featuredpromo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressaffiliatefeaturedpromogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressaffiliatefeaturedpromogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppSignature is AppSignature Setter
// 请求签名
func (r *AliexpressaffiliatefeaturedpromogetAPIRequest) SetAppSignature(_appSignature string) error {
	r._appSignature = _appSignature
	r.Set("app_signature", _appSignature)
	return nil
}

// GetAppSignature AppSignature Getter
func (r AliexpressaffiliatefeaturedpromogetAPIRequest) GetAppSignature() string {
	return r._appSignature
}

// SetFields is Fields Setter
// 返回字段列表
func (r *AliexpressaffiliatefeaturedpromogetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r AliexpressaffiliatefeaturedpromogetAPIRequest) GetFields() string {
	return r._fields
}

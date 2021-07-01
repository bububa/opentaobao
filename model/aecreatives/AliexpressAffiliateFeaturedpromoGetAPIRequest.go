package aecreatives

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressAffiliateFeaturedpromoGetAPIRequest
联盟主题推广活动信息获取 API请求
aliexpress.affiliate.featuredpromo.get

获取联盟主题推广活动信息 */
type AliexpressAffiliateFeaturedpromoGetAPIRequest struct {
	model.Params
	// 请求签名
	_appSignature string
	// 返回字段列表
	_fields string
}

// NewAliexpressAffiliateFeaturedpromoGetRequest 初始化AliexpressAffiliateFeaturedpromoGetAPIRequest对象
func NewAliexpressAffiliateFeaturedpromoGetRequest() *AliexpressAffiliateFeaturedpromoGetAPIRequest {
	return &AliexpressAffiliateFeaturedpromoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressAffiliateFeaturedpromoGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.affiliate.featuredpromo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressAffiliateFeaturedpromoGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AppSignature Setter
// 请求签名
func (r *AliexpressAffiliateFeaturedpromoGetAPIRequest) SetAppSignature(_appSignature string) error {
	r._appSignature = _appSignature
	r.Set("app_signature", _appSignature)
	return nil
}

// Get AppSignature Getter
func (r AliexpressAffiliateFeaturedpromoGetAPIRequest) GetAppSignature() string {
	return r._appSignature
}

// Set is Fields Setter
// 返回字段列表
func (r *AliexpressAffiliateFeaturedpromoGetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// Get Fields Getter
func (r AliexpressAffiliateFeaturedpromoGetAPIRequest) GetFields() string {
	return r._fields
}

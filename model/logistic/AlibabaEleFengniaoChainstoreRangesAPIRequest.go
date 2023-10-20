package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaelefengniaochainstorerangesAPIRequest 蜂鸟查询门店配送范围接口 API请求
// alibaba.ele.fengniao.chainstore.ranges
//
// 蜂鸟查询门店配送范围接口
type AlibabaelefengniaochainstorerangesAPIRequest struct {
	model.Params
	// 商户code
	_merchantCode string
	// appId
	_appId string
	// 门店code
	_chainstoreCode string
}

// NewAlibabaelefengniaochainstorerangesRequest 初始化AlibabaelefengniaochainstorerangesAPIRequest对象
func NewAlibabaelefengniaochainstorerangesRequest() *AlibabaelefengniaochainstorerangesAPIRequest {
	return &AlibabaelefengniaochainstorerangesAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaelefengniaochainstorerangesAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.fengniao.chainstore.ranges"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaelefengniaochainstorerangesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaelefengniaochainstorerangesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMerchantCode is MerchantCode Setter
// 商户code
func (r *AlibabaelefengniaochainstorerangesAPIRequest) SetMerchantCode(_merchantCode string) error {
	r._merchantCode = _merchantCode
	r.Set("merchant_code", _merchantCode)
	return nil
}

// GetMerchantCode MerchantCode Getter
func (r AlibabaelefengniaochainstorerangesAPIRequest) GetMerchantCode() string {
	return r._merchantCode
}

// SetAppId is AppId Setter
// appId
func (r *AlibabaelefengniaochainstorerangesAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r AlibabaelefengniaochainstorerangesAPIRequest) GetAppId() string {
	return r._appId
}

// SetChainstoreCode is ChainstoreCode Setter
// 门店code
func (r *AlibabaelefengniaochainstorerangesAPIRequest) SetChainstoreCode(_chainstoreCode string) error {
	r._chainstoreCode = _chainstoreCode
	r.Set("chainstore_code", _chainstoreCode)
	return nil
}

// GetChainstoreCode ChainstoreCode Getter
func (r AlibabaelefengniaochainstorerangesAPIRequest) GetChainstoreCode() string {
	return r._chainstoreCode
}

package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleFengniaoChainstoreRangesAPIRequest
蜂鸟查询门店配送范围接口 API请求
alibaba.ele.fengniao.chainstore.ranges

蜂鸟查询门店配送范围接口 */
type AlibabaEleFengniaoChainstoreRangesAPIRequest struct {
	model.Params
	// 商户code
	_merchantCode string
	// appId
	_appId string
	// 门店code
	_chainstoreCode string
}

// NewAlibabaEleFengniaoChainstoreRangesRequest 初始化AlibabaEleFengniaoChainstoreRangesAPIRequest对象
func NewAlibabaEleFengniaoChainstoreRangesRequest() *AlibabaEleFengniaoChainstoreRangesAPIRequest {
	return &AlibabaEleFengniaoChainstoreRangesAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoChainstoreRangesAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.fengniao.chainstore.ranges"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoChainstoreRangesAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MerchantCode Setter
// 商户code
func (r *AlibabaEleFengniaoChainstoreRangesAPIRequest) SetMerchantCode(_merchantCode string) error {
	r._merchantCode = _merchantCode
	r.Set("merchant_code", _merchantCode)
	return nil
}

// Get MerchantCode Getter
func (r AlibabaEleFengniaoChainstoreRangesAPIRequest) GetMerchantCode() string {
	return r._merchantCode
}

// Set is AppId Setter
// appId
func (r *AlibabaEleFengniaoChainstoreRangesAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// Get AppId Getter
func (r AlibabaEleFengniaoChainstoreRangesAPIRequest) GetAppId() string {
	return r._appId
}

// Set is ChainstoreCode Setter
// 门店code
func (r *AlibabaEleFengniaoChainstoreRangesAPIRequest) SetChainstoreCode(_chainstoreCode string) error {
	r._chainstoreCode = _chainstoreCode
	r.Set("chainstore_code", _chainstoreCode)
	return nil
}

// Get ChainstoreCode Getter
func (r AlibabaEleFengniaoChainstoreRangesAPIRequest) GetChainstoreCode() string {
	return r._chainstoreCode
}

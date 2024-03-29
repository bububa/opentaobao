package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleFengniaoChainstoreRangesAPIRequest 蜂鸟查询门店配送范围接口 API请求
// alibaba.ele.fengniao.chainstore.ranges
//
// 蜂鸟查询门店配送范围接口
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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEleFengniaoChainstoreRangesAPIRequest) Reset() {
	r._merchantCode = ""
	r._appId = ""
	r._chainstoreCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoChainstoreRangesAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.fengniao.chainstore.ranges"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoChainstoreRangesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEleFengniaoChainstoreRangesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMerchantCode is MerchantCode Setter
// 商户code
func (r *AlibabaEleFengniaoChainstoreRangesAPIRequest) SetMerchantCode(_merchantCode string) error {
	r._merchantCode = _merchantCode
	r.Set("merchant_code", _merchantCode)
	return nil
}

// GetMerchantCode MerchantCode Getter
func (r AlibabaEleFengniaoChainstoreRangesAPIRequest) GetMerchantCode() string {
	return r._merchantCode
}

// SetAppId is AppId Setter
// appId
func (r *AlibabaEleFengniaoChainstoreRangesAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r AlibabaEleFengniaoChainstoreRangesAPIRequest) GetAppId() string {
	return r._appId
}

// SetChainstoreCode is ChainstoreCode Setter
// 门店code
func (r *AlibabaEleFengniaoChainstoreRangesAPIRequest) SetChainstoreCode(_chainstoreCode string) error {
	r._chainstoreCode = _chainstoreCode
	r.Set("chainstore_code", _chainstoreCode)
	return nil
}

// GetChainstoreCode ChainstoreCode Getter
func (r AlibabaEleFengniaoChainstoreRangesAPIRequest) GetChainstoreCode() string {
	return r._chainstoreCode
}

var poolAlibabaEleFengniaoChainstoreRangesAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEleFengniaoChainstoreRangesRequest()
	},
}

// GetAlibabaEleFengniaoChainstoreRangesRequest 从 sync.Pool 获取 AlibabaEleFengniaoChainstoreRangesAPIRequest
func GetAlibabaEleFengniaoChainstoreRangesAPIRequest() *AlibabaEleFengniaoChainstoreRangesAPIRequest {
	return poolAlibabaEleFengniaoChainstoreRangesAPIRequest.Get().(*AlibabaEleFengniaoChainstoreRangesAPIRequest)
}

// ReleaseAlibabaEleFengniaoChainstoreRangesAPIRequest 将 AlibabaEleFengniaoChainstoreRangesAPIRequest 放入 sync.Pool
func ReleaseAlibabaEleFengniaoChainstoreRangesAPIRequest(v *AlibabaEleFengniaoChainstoreRangesAPIRequest) {
	v.Reset()
	poolAlibabaEleFengniaoChainstoreRangesAPIRequest.Put(v)
}

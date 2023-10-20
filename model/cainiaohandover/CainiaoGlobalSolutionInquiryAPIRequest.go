package cainiaohandover

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalSolutionInquiryAPIRequest 解决方案询盘 API请求
// cainiao.global.solution.inquiry
//
// 根据交易单号查询可用的解决方案
type CainiaoGlobalSolutionInquiryAPIRequest struct {
	model.Params
	// 包裹参数
	_packageParams []OpenPackageParam
	// 多语言，zh_CN中文、en_US:英文、ru_RU俄语
	_locale string
	// 交易单参数
	_tradeOrderParam *OpenTradeOrderParam
}

// NewCainiaoGlobalSolutionInquiryRequest 初始化CainiaoGlobalSolutionInquiryAPIRequest对象
func NewCainiaoGlobalSolutionInquiryRequest() *CainiaoGlobalSolutionInquiryAPIRequest {
	return &CainiaoGlobalSolutionInquiryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGlobalSolutionInquiryAPIRequest) GetApiMethodName() string {
	return "cainiao.global.solution.inquiry"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGlobalSolutionInquiryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoGlobalSolutionInquiryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPackageParams is PackageParams Setter
// 包裹参数
func (r *CainiaoGlobalSolutionInquiryAPIRequest) SetPackageParams(_packageParams []OpenPackageParam) error {
	r._packageParams = _packageParams
	r.Set("package_params", _packageParams)
	return nil
}

// GetPackageParams PackageParams Getter
func (r CainiaoGlobalSolutionInquiryAPIRequest) GetPackageParams() []OpenPackageParam {
	return r._packageParams
}

// SetLocale is Locale Setter
// 多语言，zh_CN中文、en_US:英文、ru_RU俄语
func (r *CainiaoGlobalSolutionInquiryAPIRequest) SetLocale(_locale string) error {
	r._locale = _locale
	r.Set("locale", _locale)
	return nil
}

// GetLocale Locale Getter
func (r CainiaoGlobalSolutionInquiryAPIRequest) GetLocale() string {
	return r._locale
}

// SetTradeOrderParam is TradeOrderParam Setter
// 交易单参数
func (r *CainiaoGlobalSolutionInquiryAPIRequest) SetTradeOrderParam(_tradeOrderParam *OpenTradeOrderParam) error {
	r._tradeOrderParam = _tradeOrderParam
	r.Set("trade_order_param", _tradeOrderParam)
	return nil
}

// GetTradeOrderParam TradeOrderParam Getter
func (r CainiaoGlobalSolutionInquiryAPIRequest) GetTradeOrderParam() *OpenTradeOrderParam {
	return r._tradeOrderParam
}

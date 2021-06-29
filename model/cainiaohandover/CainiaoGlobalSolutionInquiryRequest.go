package cainiaohandover

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
解决方案询盘 API请求
cainiao.global.solution.inquiry

根据交易单号查询可用的解决方案
*/
type CainiaoGlobalSolutionInquiryRequest struct {
    model.Params
    // 多语言，zh_CN中文、en_US:英文、ru_RU俄语
    _locale   string
    // 交易单参数
    _tradeOrderParam   *OpenTradeOrderParam
    // 商家信息
    _sellerInfoParam   *OpenSellerInfoParam
    // 包裹参数
    _packageParams   []OpenPackageParam
}

// 初始化CainiaoGlobalSolutionInquiryRequest对象
func NewCainiaoGlobalSolutionInquiryRequest() *CainiaoGlobalSolutionInquiryRequest{
    return &CainiaoGlobalSolutionInquiryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoGlobalSolutionInquiryRequest) GetApiMethodName() string {
    return "cainiao.global.solution.inquiry"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoGlobalSolutionInquiryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Locale Setter
// 多语言，zh_CN中文、en_US:英文、ru_RU俄语
func (r *CainiaoGlobalSolutionInquiryRequest) SetLocale(_locale string) error {
    r._locale = _locale
    r.Set("locale", _locale)
    return nil
}

// Locale Getter
func (r CainiaoGlobalSolutionInquiryRequest) GetLocale() string {
    return r._locale
}
// TradeOrderParam Setter
// 交易单参数
func (r *CainiaoGlobalSolutionInquiryRequest) SetTradeOrderParam(_tradeOrderParam *OpenTradeOrderParam) error {
    r._tradeOrderParam = _tradeOrderParam
    r.Set("trade_order_param", _tradeOrderParam)
    return nil
}

// TradeOrderParam Getter
func (r CainiaoGlobalSolutionInquiryRequest) GetTradeOrderParam() *OpenTradeOrderParam {
    return r._tradeOrderParam
}
// SellerInfoParam Setter
// 商家信息
func (r *CainiaoGlobalSolutionInquiryRequest) SetSellerInfoParam(_sellerInfoParam *OpenSellerInfoParam) error {
    r._sellerInfoParam = _sellerInfoParam
    r.Set("seller_info_param", _sellerInfoParam)
    return nil
}

// SellerInfoParam Getter
func (r CainiaoGlobalSolutionInquiryRequest) GetSellerInfoParam() *OpenSellerInfoParam {
    return r._sellerInfoParam
}
// PackageParams Setter
// 包裹参数
func (r *CainiaoGlobalSolutionInquiryRequest) SetPackageParams(_packageParams []OpenPackageParam) error {
    r._packageParams = _packageParams
    r.Set("package_params", _packageParams)
    return nil
}

// PackageParams Getter
func (r CainiaoGlobalSolutionInquiryRequest) GetPackageParams() []OpenPackageParam {
    return r._packageParams
}

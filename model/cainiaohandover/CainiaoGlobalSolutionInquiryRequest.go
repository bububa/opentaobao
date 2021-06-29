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
    locale   string
    // 交易单参数
    tradeOrderParam   *OpenTradeOrderParam
    // 商家信息
    sellerInfoParam   *OpenSellerInfoParam
    // 包裹参数
    packageParams   []OpenPackageParam
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
func (r *CainiaoGlobalSolutionInquiryRequest) SetLocale(locale string) error {
    r.locale = locale
    r.Set("locale", locale)
    return nil
}

// Locale Getter
func (r CainiaoGlobalSolutionInquiryRequest) GetLocale() string {
    return r.locale
}
// TradeOrderParam Setter
// 交易单参数
func (r *CainiaoGlobalSolutionInquiryRequest) SetTradeOrderParam(tradeOrderParam *OpenTradeOrderParam) error {
    r.tradeOrderParam = tradeOrderParam
    r.Set("trade_order_param", tradeOrderParam)
    return nil
}

// TradeOrderParam Getter
func (r CainiaoGlobalSolutionInquiryRequest) GetTradeOrderParam() *OpenTradeOrderParam {
    return r.tradeOrderParam
}
// SellerInfoParam Setter
// 商家信息
func (r *CainiaoGlobalSolutionInquiryRequest) SetSellerInfoParam(sellerInfoParam *OpenSellerInfoParam) error {
    r.sellerInfoParam = sellerInfoParam
    r.Set("seller_info_param", sellerInfoParam)
    return nil
}

// SellerInfoParam Getter
func (r CainiaoGlobalSolutionInquiryRequest) GetSellerInfoParam() *OpenSellerInfoParam {
    return r.sellerInfoParam
}
// PackageParams Setter
// 包裹参数
func (r *CainiaoGlobalSolutionInquiryRequest) SetPackageParams(packageParams []OpenPackageParam) error {
    r.packageParams = packageParams
    r.Set("package_params", packageParams)
    return nil
}

// PackageParams Getter
func (r CainiaoGlobalSolutionInquiryRequest) GetPackageParams() []OpenPackageParam {
    return r.packageParams
}

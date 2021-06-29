package cainiaohandover

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
解决方案询盘 APIRequest
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

func NewCainiaoGlobalSolutionInquiryRequest() *CainiaoGlobalSolutionInquiryRequest{
    return &CainiaoGlobalSolutionInquiryRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoGlobalSolutionInquiryRequest) GetApiMethodName() string {
    return "cainiao.global.solution.inquiry"
}

func (r CainiaoGlobalSolutionInquiryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoGlobalSolutionInquiryRequest) SetLocale(locale string) error {
    r.locale = locale
    r.Set("locale", locale)
    return nil
}

func (r CainiaoGlobalSolutionInquiryRequest) GetLocale() string {
    return r.locale
}

func (r *CainiaoGlobalSolutionInquiryRequest) SetTradeOrderParam(tradeOrderParam *OpenTradeOrderParam) error {
    r.tradeOrderParam = tradeOrderParam
    r.Set("trade_order_param", tradeOrderParam)
    return nil
}

func (r CainiaoGlobalSolutionInquiryRequest) GetTradeOrderParam() *OpenTradeOrderParam {
    return r.tradeOrderParam
}

func (r *CainiaoGlobalSolutionInquiryRequest) SetSellerInfoParam(sellerInfoParam *OpenSellerInfoParam) error {
    r.sellerInfoParam = sellerInfoParam
    r.Set("seller_info_param", sellerInfoParam)
    return nil
}

func (r CainiaoGlobalSolutionInquiryRequest) GetSellerInfoParam() *OpenSellerInfoParam {
    return r.sellerInfoParam
}

func (r *CainiaoGlobalSolutionInquiryRequest) SetPackageParams(packageParams []OpenPackageParam) error {
    r.packageParams = packageParams
    r.Set("package_params", packageParams)
    return nil
}

func (r CainiaoGlobalSolutionInquiryRequest) GetPackageParams() []OpenPackageParam {
    return r.packageParams
}


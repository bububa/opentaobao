package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
蜂鸟查询门店配送范围接口 API请求
alibaba.ele.fengniao.chainstore.ranges

蜂鸟查询门店配送范围接口
*/
type AlibabaEleFengniaoChainstoreRangesRequest struct {
    model.Params
    // 商户code
    merchantCode   string
    // appId
    appId   string
    // 门店code
    chainstoreCode   string
}

// 初始化AlibabaEleFengniaoChainstoreRangesRequest对象
func NewAlibabaEleFengniaoChainstoreRangesRequest() *AlibabaEleFengniaoChainstoreRangesRequest{
    return &AlibabaEleFengniaoChainstoreRangesRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoChainstoreRangesRequest) GetApiMethodName() string {
    return "alibaba.ele.fengniao.chainstore.ranges"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoChainstoreRangesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MerchantCode Setter
// 商户code
func (r *AlibabaEleFengniaoChainstoreRangesRequest) SetMerchantCode(merchantCode string) error {
    r.merchantCode = merchantCode
    r.Set("merchant_code", merchantCode)
    return nil
}

// MerchantCode Getter
func (r AlibabaEleFengniaoChainstoreRangesRequest) GetMerchantCode() string {
    return r.merchantCode
}
// AppId Setter
// appId
func (r *AlibabaEleFengniaoChainstoreRangesRequest) SetAppId(appId string) error {
    r.appId = appId
    r.Set("app_id", appId)
    return nil
}

// AppId Getter
func (r AlibabaEleFengniaoChainstoreRangesRequest) GetAppId() string {
    return r.appId
}
// ChainstoreCode Setter
// 门店code
func (r *AlibabaEleFengniaoChainstoreRangesRequest) SetChainstoreCode(chainstoreCode string) error {
    r.chainstoreCode = chainstoreCode
    r.Set("chainstore_code", chainstoreCode)
    return nil
}

// ChainstoreCode Getter
func (r AlibabaEleFengniaoChainstoreRangesRequest) GetChainstoreCode() string {
    return r.chainstoreCode
}

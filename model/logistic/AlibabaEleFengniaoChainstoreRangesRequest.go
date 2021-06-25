package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
蜂鸟查询门店配送范围接口 APIRequest
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

func NewAlibabaEleFengniaoChainstoreRangesRequest() *AlibabaEleFengniaoChainstoreRangesRequest{
    return &AlibabaEleFengniaoChainstoreRangesRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEleFengniaoChainstoreRangesRequest) GetApiMethodName() string {
    return "alibaba.ele.fengniao.chainstore.ranges"
}

func (r AlibabaEleFengniaoChainstoreRangesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEleFengniaoChainstoreRangesRequest) SetMerchantCode(merchantCode string) error {
    r.merchantCode = merchantCode
    r.Set("merchant_code", merchantCode)
    return nil
}

func (r AlibabaEleFengniaoChainstoreRangesRequest) GetMerchantCode() string {
    return r.merchantCode
}

func (r *AlibabaEleFengniaoChainstoreRangesRequest) SetAppId(appId string) error {
    r.appId = appId
    r.Set("app_id", appId)
    return nil
}

func (r AlibabaEleFengniaoChainstoreRangesRequest) GetAppId() string {
    return r.appId
}

func (r *AlibabaEleFengniaoChainstoreRangesRequest) SetChainstoreCode(chainstoreCode string) error {
    r.chainstoreCode = chainstoreCode
    r.Set("chainstore_code", chainstoreCode)
    return nil
}

func (r AlibabaEleFengniaoChainstoreRangesRequest) GetChainstoreCode() string {
    return r.chainstoreCode
}


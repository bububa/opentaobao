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
    _merchantCode   string
    // appId
    _appId   string
    // 门店code
    _chainstoreCode   string
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
func (r *AlibabaEleFengniaoChainstoreRangesRequest) SetMerchantCode(_merchantCode string) error {
    r._merchantCode = _merchantCode
    r.Set("merchant_code", _merchantCode)
    return nil
}

// MerchantCode Getter
func (r AlibabaEleFengniaoChainstoreRangesRequest) GetMerchantCode() string {
    return r._merchantCode
}
// AppId Setter
// appId
func (r *AlibabaEleFengniaoChainstoreRangesRequest) SetAppId(_appId string) error {
    r._appId = _appId
    r.Set("app_id", _appId)
    return nil
}

// AppId Getter
func (r AlibabaEleFengniaoChainstoreRangesRequest) GetAppId() string {
    return r._appId
}
// ChainstoreCode Setter
// 门店code
func (r *AlibabaEleFengniaoChainstoreRangesRequest) SetChainstoreCode(_chainstoreCode string) error {
    r._chainstoreCode = _chainstoreCode
    r.Set("chainstore_code", _chainstoreCode)
    return nil
}

// ChainstoreCode Getter
func (r AlibabaEleFengniaoChainstoreRangesRequest) GetChainstoreCode() string {
    return r._chainstoreCode
}

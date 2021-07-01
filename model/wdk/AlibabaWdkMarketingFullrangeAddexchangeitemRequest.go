package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
全场增加换购品 API请求
alibaba.wdk.marketing.fullrange.addexchangeitem

全场增加换购品
*/
type AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest struct {
    model.Params
    // 系统自动生成
    _param0   *ItemStairSku
    // 系统自动生成
    _param1   *CommonActivityParam
}

// 初始化AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest对象
func NewAlibabaWdkMarketingFullrangeAddexchangeitemRequest() *AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest{
    return &AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.fullrange.addexchangeitem"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 系统自动生成
func (r *AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest) SetParam0(_param0 *ItemStairSku) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest) GetParam0() *ItemStairSku {
    return r._param0
}
// Param1 Setter
// 系统自动生成
func (r *AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest) GetParam1() *CommonActivityParam {
    return r._param1
}

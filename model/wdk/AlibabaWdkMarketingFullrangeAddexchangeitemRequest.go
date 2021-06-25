package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
全场增加换购品 APIRequest
alibaba.wdk.marketing.fullrange.addexchangeitem

全场增加换购品
*/
type AlibabaWdkMarketingFullrangeAddexchangeitemRequest struct {
    model.Params

    // 系统自动生成
    param0   *ItemStairSku 

    // 系统自动生成
    param1   *CommonActivityParam 

}

func NewAlibabaWdkMarketingFullrangeAddexchangeitemRequest() *AlibabaWdkMarketingFullrangeAddexchangeitemRequest{
    return &AlibabaWdkMarketingFullrangeAddexchangeitemRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingFullrangeAddexchangeitemRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.fullrange.addexchangeitem"
}

func (r AlibabaWdkMarketingFullrangeAddexchangeitemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingFullrangeAddexchangeitemRequest) SetParam0(param0 *ItemStairSku) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaWdkMarketingFullrangeAddexchangeitemRequest) GetParam0() *ItemStairSku {
    return r.param0
}

func (r *AlibabaWdkMarketingFullrangeAddexchangeitemRequest) SetParam1(param1 *CommonActivityParam) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r AlibabaWdkMarketingFullrangeAddexchangeitemRequest) GetParam1() *CommonActivityParam {
    return r.param1
}


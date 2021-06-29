package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询特价商品 API请求
alibaba.wdk.marketing.itemdiscount.queryitems

查询参加特价活动的商品优惠详情
*/
type AlibabaWdkMarketingItemdiscountQueryitemsRequest struct {
    model.Params
    // 查询入参
    param   *ActivitySkuQuery
}

// 初始化AlibabaWdkMarketingItemdiscountQueryitemsRequest对象
func NewAlibabaWdkMarketingItemdiscountQueryitemsRequest() *AlibabaWdkMarketingItemdiscountQueryitemsRequest{
    return &AlibabaWdkMarketingItemdiscountQueryitemsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItemdiscountQueryitemsRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itemdiscount.queryitems"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItemdiscountQueryitemsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 查询入参
func (r *AlibabaWdkMarketingItemdiscountQueryitemsRequest) SetParam(param *ActivitySkuQuery) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaWdkMarketingItemdiscountQueryitemsRequest) GetParam() *ActivitySkuQuery {
    return r.param
}

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
type AlibabaWdkMarketingItemdiscountQueryitemsAPIRequest struct {
    model.Params
    // 查询入参
    _param   *ActivitySkuQuery
}

// 初始化AlibabaWdkMarketingItemdiscountQueryitemsAPIRequest对象
func NewAlibabaWdkMarketingItemdiscountQueryitemsRequest() *AlibabaWdkMarketingItemdiscountQueryitemsAPIRequest{
    return &AlibabaWdkMarketingItemdiscountQueryitemsAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItemdiscountQueryitemsAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itemdiscount.queryitems"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItemdiscountQueryitemsAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 查询入参
func (r *AlibabaWdkMarketingItemdiscountQueryitemsAPIRequest) SetParam(_param *ActivitySkuQuery) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaWdkMarketingItemdiscountQueryitemsAPIRequest) GetParam() *ActivitySkuQuery {
    return r._param
}

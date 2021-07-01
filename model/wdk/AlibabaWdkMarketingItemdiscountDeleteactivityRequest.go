package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除商品特价活动 API请求
alibaba.wdk.marketing.itemdiscount.deleteactivity

删除商品特价活动
*/
type AlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest struct {
    model.Params
    // 需要删除的活动的信息
    _param   *CommonActivityRequest
}

// 初始化AlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest对象
func NewAlibabaWdkMarketingItemdiscountDeleteactivityRequest() *AlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest{
    return &AlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itemdiscount.deleteactivity"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 需要删除的活动的信息
func (r *AlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest) SetParam(_param *CommonActivityRequest) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest) GetParam() *CommonActivityRequest {
    return r._param
}

package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品池排除商品【品类优惠使用】 API请求
alibaba.wdk.marketing.itempool.excludeskucode

品类优惠新增排除池
*/
type AlibabaWdkMarketingItempoolExcludeskucodeRequest struct {
    model.Params
    // 商品对象
    param0   *ItemPoolSku
    // 活动基本信息
    param1   *CommonActivityParam
}

// 初始化AlibabaWdkMarketingItempoolExcludeskucodeRequest对象
func NewAlibabaWdkMarketingItempoolExcludeskucodeRequest() *AlibabaWdkMarketingItempoolExcludeskucodeRequest{
    return &AlibabaWdkMarketingItempoolExcludeskucodeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolExcludeskucodeRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itempool.excludeskucode"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolExcludeskucodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 商品对象
func (r *AlibabaWdkMarketingItempoolExcludeskucodeRequest) SetParam0(param0 *ItemPoolSku) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkMarketingItempoolExcludeskucodeRequest) GetParam0() *ItemPoolSku {
    return r.param0
}
// Param1 Setter
// 活动基本信息
func (r *AlibabaWdkMarketingItempoolExcludeskucodeRequest) SetParam1(param1 *CommonActivityParam) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

// Param1 Getter
func (r AlibabaWdkMarketingItempoolExcludeskucodeRequest) GetParam1() *CommonActivityParam {
    return r.param1
}

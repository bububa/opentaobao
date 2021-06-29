package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除单品特价活动【同城零售】 API请求
alibaba.retail.marketing.itemdiscount.activity.delete

同城零售单品特价活动删除
*/
type AlibabaRetailMarketingItemdiscountActivityDeleteRequest struct {
    model.Params
    // 删除活动参数
    param   *ItemDiscountActivityOperateRequest
}

// 初始化AlibabaRetailMarketingItemdiscountActivityDeleteRequest对象
func NewAlibabaRetailMarketingItemdiscountActivityDeleteRequest() *AlibabaRetailMarketingItemdiscountActivityDeleteRequest{
    return &AlibabaRetailMarketingItemdiscountActivityDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItemdiscountActivityDeleteRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.itemdiscount.activity.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItemdiscountActivityDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 删除活动参数
func (r *AlibabaRetailMarketingItemdiscountActivityDeleteRequest) SetParam(param *ItemDiscountActivityOperateRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaRetailMarketingItemdiscountActivityDeleteRequest) GetParam() *ItemDiscountActivityOperateRequest {
    return r.param
}

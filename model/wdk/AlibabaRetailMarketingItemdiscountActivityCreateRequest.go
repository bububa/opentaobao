package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建单品特价活动【同城零售】 API请求
alibaba.retail.marketing.itemdiscount.activity.create

同城零售单品特价活动创建
*/
type AlibabaRetailMarketingItemdiscountActivityCreateRequest struct {
    model.Params
    // 创建活动参数
    _param   *ItemDiscountActivityOperateRequest
}

// 初始化AlibabaRetailMarketingItemdiscountActivityCreateRequest对象
func NewAlibabaRetailMarketingItemdiscountActivityCreateRequest() *AlibabaRetailMarketingItemdiscountActivityCreateRequest{
    return &AlibabaRetailMarketingItemdiscountActivityCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItemdiscountActivityCreateRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.itemdiscount.activity.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItemdiscountActivityCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 创建活动参数
func (r *AlibabaRetailMarketingItemdiscountActivityCreateRequest) SetParam(_param *ItemDiscountActivityOperateRequest) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaRetailMarketingItemdiscountActivityCreateRequest) GetParam() *ItemDiscountActivityOperateRequest {
    return r._param
}

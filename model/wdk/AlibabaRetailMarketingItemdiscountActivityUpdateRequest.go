package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新单品特价活动【同城零售】 API请求
alibaba.retail.marketing.itemdiscount.activity.update

同城零售单品特价活动更新
*/
type AlibabaRetailMarketingItemdiscountActivityUpdateRequest struct {
    model.Params
    // 创建活动参数
    param   *ItemDiscountActivityOperateRequest
}

// 初始化AlibabaRetailMarketingItemdiscountActivityUpdateRequest对象
func NewAlibabaRetailMarketingItemdiscountActivityUpdateRequest() *AlibabaRetailMarketingItemdiscountActivityUpdateRequest{
    return &AlibabaRetailMarketingItemdiscountActivityUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItemdiscountActivityUpdateRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.itemdiscount.activity.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItemdiscountActivityUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 创建活动参数
func (r *AlibabaRetailMarketingItemdiscountActivityUpdateRequest) SetParam(param *ItemDiscountActivityOperateRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaRetailMarketingItemdiscountActivityUpdateRequest) GetParam() *ItemDiscountActivityOperateRequest {
    return r.param
}

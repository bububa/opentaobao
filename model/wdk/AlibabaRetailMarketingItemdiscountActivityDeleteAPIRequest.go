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
type AlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest struct {
    model.Params
    // 删除活动参数
    _param   *ItemDiscountActivityOperateRequest
}

// 初始化AlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest对象
func NewAlibabaRetailMarketingItemdiscountActivityDeleteRequest() *AlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest{
    return &AlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.itemdiscount.activity.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 删除活动参数
func (r *AlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest) SetParam(_param *ItemDiscountActivityOperateRequest) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest) GetParam() *ItemDiscountActivityOperateRequest {
    return r._param
}

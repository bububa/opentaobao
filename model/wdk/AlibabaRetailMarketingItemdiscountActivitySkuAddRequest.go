package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
特价活动新增商品 API请求
alibaba.retail.marketing.itemdiscount.activity.sku.add

新增或更新活动商品信息【同城零售】
*/
type AlibabaRetailMarketingItemdiscountActivitySkuAddRequest struct {
    model.Params
    // 添加活动商品参数
    param   *ItemDiscountActivityElementOperateRequest
}

// 初始化AlibabaRetailMarketingItemdiscountActivitySkuAddRequest对象
func NewAlibabaRetailMarketingItemdiscountActivitySkuAddRequest() *AlibabaRetailMarketingItemdiscountActivitySkuAddRequest{
    return &AlibabaRetailMarketingItemdiscountActivitySkuAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItemdiscountActivitySkuAddRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.itemdiscount.activity.sku.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItemdiscountActivitySkuAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 添加活动商品参数
func (r *AlibabaRetailMarketingItemdiscountActivitySkuAddRequest) SetParam(param *ItemDiscountActivityElementOperateRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaRetailMarketingItemdiscountActivitySkuAddRequest) GetParam() *ItemDiscountActivityElementOperateRequest {
    return r.param
}

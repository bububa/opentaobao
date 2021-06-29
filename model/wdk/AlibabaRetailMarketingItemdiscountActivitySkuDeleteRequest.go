package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除特价活动商品 API请求
alibaba.retail.marketing.itemdiscount.activity.sku.delete

删除活动商品信息【同城零售】
*/
type AlibabaRetailMarketingItemdiscountActivitySkuDeleteRequest struct {
    model.Params
    // 添加活动商品参数
    param   *ItemDiscountActivityElementOperateRequest
}

// 初始化AlibabaRetailMarketingItemdiscountActivitySkuDeleteRequest对象
func NewAlibabaRetailMarketingItemdiscountActivitySkuDeleteRequest() *AlibabaRetailMarketingItemdiscountActivitySkuDeleteRequest{
    return &AlibabaRetailMarketingItemdiscountActivitySkuDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItemdiscountActivitySkuDeleteRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.itemdiscount.activity.sku.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItemdiscountActivitySkuDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 添加活动商品参数
func (r *AlibabaRetailMarketingItemdiscountActivitySkuDeleteRequest) SetParam(param *ItemDiscountActivityElementOperateRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaRetailMarketingItemdiscountActivitySkuDeleteRequest) GetParam() *ItemDiscountActivityElementOperateRequest {
    return r.param
}

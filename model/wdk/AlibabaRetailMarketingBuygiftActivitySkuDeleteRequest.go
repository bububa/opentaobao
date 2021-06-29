package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除单品买赠活动商品 API请求
alibaba.retail.marketing.buygift.activity.sku.delete

删除单品买赠活动商品信息【同城零售】
*/
type AlibabaRetailMarketingBuygiftActivitySkuDeleteRequest struct {
    model.Params
    // 删除买赠活动商品参数
    param   *BuyGiftActivitySkuOperateRequest
}

// 初始化AlibabaRetailMarketingBuygiftActivitySkuDeleteRequest对象
func NewAlibabaRetailMarketingBuygiftActivitySkuDeleteRequest() *AlibabaRetailMarketingBuygiftActivitySkuDeleteRequest{
    return &AlibabaRetailMarketingBuygiftActivitySkuDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingBuygiftActivitySkuDeleteRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.buygift.activity.sku.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingBuygiftActivitySkuDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 删除买赠活动商品参数
func (r *AlibabaRetailMarketingBuygiftActivitySkuDeleteRequest) SetParam(param *BuyGiftActivitySkuOperateRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaRetailMarketingBuygiftActivitySkuDeleteRequest) GetParam() *BuyGiftActivitySkuOperateRequest {
    return r.param
}

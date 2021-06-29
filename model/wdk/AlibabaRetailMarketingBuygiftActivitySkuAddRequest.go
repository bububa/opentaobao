package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加单品买赠活动商品 API请求
alibaba.retail.marketing.buygift.activity.sku.add

新增或更新单品买赠活动商品信息【同城零售】
*/
type AlibabaRetailMarketingBuygiftActivitySkuAddRequest struct {
    model.Params
    // 添加活动商品参数
    param   *BuyGiftActivitySkuOperateRequest
}

// 初始化AlibabaRetailMarketingBuygiftActivitySkuAddRequest对象
func NewAlibabaRetailMarketingBuygiftActivitySkuAddRequest() *AlibabaRetailMarketingBuygiftActivitySkuAddRequest{
    return &AlibabaRetailMarketingBuygiftActivitySkuAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingBuygiftActivitySkuAddRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.buygift.activity.sku.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingBuygiftActivitySkuAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 添加活动商品参数
func (r *AlibabaRetailMarketingBuygiftActivitySkuAddRequest) SetParam(param *BuyGiftActivitySkuOperateRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaRetailMarketingBuygiftActivitySkuAddRequest) GetParam() *BuyGiftActivitySkuOperateRequest {
    return r.param
}

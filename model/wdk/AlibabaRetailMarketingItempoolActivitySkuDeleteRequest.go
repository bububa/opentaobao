package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除商品池活动商品【同城零售】 API请求
alibaba.retail.marketing.itempool.activity.sku.delete

删除商品池活动商品信息【同城零售】
*/
type AlibabaRetailMarketingItempoolActivitySkuDeleteRequest struct {
    model.Params
    // 入参
    param   *ItemPoolActivityElementOperateRequest
}

// 初始化AlibabaRetailMarketingItempoolActivitySkuDeleteRequest对象
func NewAlibabaRetailMarketingItempoolActivitySkuDeleteRequest() *AlibabaRetailMarketingItempoolActivitySkuDeleteRequest{
    return &AlibabaRetailMarketingItempoolActivitySkuDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItempoolActivitySkuDeleteRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.itempool.activity.sku.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItempoolActivitySkuDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参
func (r *AlibabaRetailMarketingItempoolActivitySkuDeleteRequest) SetParam(param *ItemPoolActivityElementOperateRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaRetailMarketingItempoolActivitySkuDeleteRequest) GetParam() *ItemPoolActivityElementOperateRequest {
    return r.param
}

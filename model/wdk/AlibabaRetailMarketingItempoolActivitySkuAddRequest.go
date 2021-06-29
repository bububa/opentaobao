package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品池活动新增商品 API请求
alibaba.retail.marketing.itempool.activity.sku.add

新增或更新商品池活动商品信息【同城零售】
*/
type AlibabaRetailMarketingItempoolActivitySkuAddRequest struct {
    model.Params
    // 入参
    param   *ItemPoolActivityElementOperateRequest
}

// 初始化AlibabaRetailMarketingItempoolActivitySkuAddRequest对象
func NewAlibabaRetailMarketingItempoolActivitySkuAddRequest() *AlibabaRetailMarketingItempoolActivitySkuAddRequest{
    return &AlibabaRetailMarketingItempoolActivitySkuAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItempoolActivitySkuAddRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.itempool.activity.sku.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItempoolActivitySkuAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参
func (r *AlibabaRetailMarketingItempoolActivitySkuAddRequest) SetParam(param *ItemPoolActivityElementOperateRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaRetailMarketingItempoolActivitySkuAddRequest) GetParam() *ItemPoolActivityElementOperateRequest {
    return r.param
}

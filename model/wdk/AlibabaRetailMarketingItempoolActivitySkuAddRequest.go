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
type AlibabaRetailMarketingItempoolActivitySkuAddAPIRequest struct {
    model.Params
    // 入参
    _param   *ItemPoolActivityElementOperateRequest
}

// 初始化AlibabaRetailMarketingItempoolActivitySkuAddAPIRequest对象
func NewAlibabaRetailMarketingItempoolActivitySkuAddRequest() *AlibabaRetailMarketingItempoolActivitySkuAddAPIRequest{
    return &AlibabaRetailMarketingItempoolActivitySkuAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItempoolActivitySkuAddAPIRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.itempool.activity.sku.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItempoolActivitySkuAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参
func (r *AlibabaRetailMarketingItempoolActivitySkuAddAPIRequest) SetParam(_param *ItemPoolActivityElementOperateRequest) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaRetailMarketingItempoolActivitySkuAddAPIRequest) GetParam() *ItemPoolActivityElementOperateRequest {
    return r._param
}

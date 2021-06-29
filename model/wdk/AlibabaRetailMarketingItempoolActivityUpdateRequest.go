package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新商品池活动【同城零售】 API请求
alibaba.retail.marketing.itempool.activity.update

同城零售商品池活动更新
*/
type AlibabaRetailMarketingItempoolActivityUpdateRequest struct {
    model.Params
    // 更新商品池活动参数
    param   *ItemPoolActivityOperateRequest
}

// 初始化AlibabaRetailMarketingItempoolActivityUpdateRequest对象
func NewAlibabaRetailMarketingItempoolActivityUpdateRequest() *AlibabaRetailMarketingItempoolActivityUpdateRequest{
    return &AlibabaRetailMarketingItempoolActivityUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItempoolActivityUpdateRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.itempool.activity.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItempoolActivityUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 更新商品池活动参数
func (r *AlibabaRetailMarketingItempoolActivityUpdateRequest) SetParam(param *ItemPoolActivityOperateRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaRetailMarketingItempoolActivityUpdateRequest) GetParam() *ItemPoolActivityOperateRequest {
    return r.param
}

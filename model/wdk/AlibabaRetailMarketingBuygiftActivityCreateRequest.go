package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建买赠活动 API请求
alibaba.retail.marketing.buygift.activity.create

同城供给买赠活动创建
*/
type AlibabaRetailMarketingBuygiftActivityCreateRequest struct {
    model.Params
    // 创建活动参数
    param   *BuyGiftActivityOperateRequest
}

// 初始化AlibabaRetailMarketingBuygiftActivityCreateRequest对象
func NewAlibabaRetailMarketingBuygiftActivityCreateRequest() *AlibabaRetailMarketingBuygiftActivityCreateRequest{
    return &AlibabaRetailMarketingBuygiftActivityCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingBuygiftActivityCreateRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.buygift.activity.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingBuygiftActivityCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 创建活动参数
func (r *AlibabaRetailMarketingBuygiftActivityCreateRequest) SetParam(param *BuyGiftActivityOperateRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaRetailMarketingBuygiftActivityCreateRequest) GetParam() *BuyGiftActivityOperateRequest {
    return r.param
}

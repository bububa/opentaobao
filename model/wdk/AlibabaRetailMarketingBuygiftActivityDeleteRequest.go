package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除单品买赠活动 API请求
alibaba.retail.marketing.buygift.activity.delete

同城零售单品特价活动删除
*/
type AlibabaRetailMarketingBuygiftActivityDeleteAPIRequest struct {
    model.Params
    // 删除活动参数
    _param   *ItemDiscountActivityOperateRequest
}

// 初始化AlibabaRetailMarketingBuygiftActivityDeleteAPIRequest对象
func NewAlibabaRetailMarketingBuygiftActivityDeleteRequest() *AlibabaRetailMarketingBuygiftActivityDeleteAPIRequest{
    return &AlibabaRetailMarketingBuygiftActivityDeleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingBuygiftActivityDeleteAPIRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.buygift.activity.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingBuygiftActivityDeleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 删除活动参数
func (r *AlibabaRetailMarketingBuygiftActivityDeleteAPIRequest) SetParam(_param *ItemDiscountActivityOperateRequest) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaRetailMarketingBuygiftActivityDeleteAPIRequest) GetParam() *ItemDiscountActivityOperateRequest {
    return r._param
}

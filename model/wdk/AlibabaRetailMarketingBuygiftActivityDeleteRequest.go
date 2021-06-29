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
type AlibabaRetailMarketingBuygiftActivityDeleteRequest struct {
    model.Params
    // 删除活动参数
    _param   *ItemDiscountActivityOperateRequest
}

// 初始化AlibabaRetailMarketingBuygiftActivityDeleteRequest对象
func NewAlibabaRetailMarketingBuygiftActivityDeleteRequest() *AlibabaRetailMarketingBuygiftActivityDeleteRequest{
    return &AlibabaRetailMarketingBuygiftActivityDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingBuygiftActivityDeleteRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.buygift.activity.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingBuygiftActivityDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 删除活动参数
func (r *AlibabaRetailMarketingBuygiftActivityDeleteRequest) SetParam(_param *ItemDiscountActivityOperateRequest) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaRetailMarketingBuygiftActivityDeleteRequest) GetParam() *ItemDiscountActivityOperateRequest {
    return r._param
}

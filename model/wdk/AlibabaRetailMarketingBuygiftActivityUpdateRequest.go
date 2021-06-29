package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新单品买赠活动 API请求
alibaba.retail.marketing.buygift.activity.update

同城零售单品买赠活动更新
*/
type AlibabaRetailMarketingBuygiftActivityUpdateRequest struct {
    model.Params
    // 更新单品买赠活动参数
    _param   *BuyGiftActivityOperateRequest
}

// 初始化AlibabaRetailMarketingBuygiftActivityUpdateRequest对象
func NewAlibabaRetailMarketingBuygiftActivityUpdateRequest() *AlibabaRetailMarketingBuygiftActivityUpdateRequest{
    return &AlibabaRetailMarketingBuygiftActivityUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingBuygiftActivityUpdateRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.buygift.activity.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingBuygiftActivityUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 更新单品买赠活动参数
func (r *AlibabaRetailMarketingBuygiftActivityUpdateRequest) SetParam(_param *BuyGiftActivityOperateRequest) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaRetailMarketingBuygiftActivityUpdateRequest) GetParam() *BuyGiftActivityOperateRequest {
    return r._param
}

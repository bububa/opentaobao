package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
应用风险概要信息查询接口 API请求
alibaba.security.jaq.app.risksummary.get

用户通过alibaba.security.jaq.app.risk.scan接口提交应用进行风险扫描后，用此接口获取风险概要信息，本接口不返回风险详细信息
*/
type AlibabaSecurityJaqAppRisksummaryGetRequest struct {
    model.Params
    // 任务唯一标识
    itemId   string
}

// 初始化AlibabaSecurityJaqAppRisksummaryGetRequest对象
func NewAlibabaSecurityJaqAppRisksummaryGetRequest() *AlibabaSecurityJaqAppRisksummaryGetRequest{
    return &AlibabaSecurityJaqAppRisksummaryGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqAppRisksummaryGetRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.app.risksummary.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqAppRisksummaryGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 任务唯一标识
func (r *AlibabaSecurityJaqAppRisksummaryGetRequest) SetItemId(itemId string) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r AlibabaSecurityJaqAppRisksummaryGetRequest) GetItemId() string {
    return r.itemId
}

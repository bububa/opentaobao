package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
存送业务查询奖品信息 API请求
alibaba.alicom.wtt.opentrade.getgiftdetails

话费宝充值送查询奖品信息
*/
type AlibabaAlicomWttOpentradeGetgiftdetailsRequest struct {
    model.Params
    // 活动ID
    _activityId   string
}

// 初始化AlibabaAlicomWttOpentradeGetgiftdetailsRequest对象
func NewAlibabaAlicomWttOpentradeGetgiftdetailsRequest() *AlibabaAlicomWttOpentradeGetgiftdetailsRequest{
    return &AlibabaAlicomWttOpentradeGetgiftdetailsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlicomWttOpentradeGetgiftdetailsRequest) GetApiMethodName() string {
    return "alibaba.alicom.wtt.opentrade.getgiftdetails"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlicomWttOpentradeGetgiftdetailsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivityId Setter
// 活动ID
func (r *AlibabaAlicomWttOpentradeGetgiftdetailsRequest) SetActivityId(_activityId string) error {
    r._activityId = _activityId
    r.Set("activity_id", _activityId)
    return nil
}

// ActivityId Getter
func (r AlibabaAlicomWttOpentradeGetgiftdetailsRequest) GetActivityId() string {
    return r._activityId
}

package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
存送业务查询奖品信息 APIRequest
alibaba.alicom.wtt.opentrade.getgiftdetails

话费宝充值送查询奖品信息
*/
type AlibabaAlicomWttOpentradeGetgiftdetailsRequest struct {
    model.Params

    // 活动ID
    activityId   string 

}

func NewAlibabaAlicomWttOpentradeGetgiftdetailsRequest() *AlibabaAlicomWttOpentradeGetgiftdetailsRequest{
    return &AlibabaAlicomWttOpentradeGetgiftdetailsRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlicomWttOpentradeGetgiftdetailsRequest) GetApiMethodName() string {
    return "alibaba.alicom.wtt.opentrade.getgiftdetails"
}

func (r AlibabaAlicomWttOpentradeGetgiftdetailsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlicomWttOpentradeGetgiftdetailsRequest) SetActivityId(activityId string) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r AlibabaAlicomWttOpentradeGetgiftdetailsRequest) GetActivityId() string {
    return r.activityId
}


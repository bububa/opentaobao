package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
心跳服务【10s一次】 APIRequest
alibaba.wdk.marketing.open.heartbeat

商家数据同步心跳服务
*/
type AlibabaWdkMarketingOpenHeartbeatRequest struct {
    model.Params

    // 心跳信息
    heartBeat   *HeartBeatBo 

}

func NewAlibabaWdkMarketingOpenHeartbeatRequest() *AlibabaWdkMarketingOpenHeartbeatRequest{
    return &AlibabaWdkMarketingOpenHeartbeatRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingOpenHeartbeatRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.open.heartbeat"
}

func (r AlibabaWdkMarketingOpenHeartbeatRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingOpenHeartbeatRequest) SetHeartBeat(heartBeat *HeartBeatBo) error {
    r.heartBeat = heartBeat
    r.Set("heart_beat", heartBeat)
    return nil
}

func (r AlibabaWdkMarketingOpenHeartbeatRequest) GetHeartBeat() *HeartBeatBo {
    return r.heartBeat
}


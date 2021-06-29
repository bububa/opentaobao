package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
心跳服务【10s一次】 API请求
alibaba.wdk.marketing.open.heartbeat

商家数据同步心跳服务
*/
type AlibabaWdkMarketingOpenHeartbeatRequest struct {
    model.Params
    // 心跳信息
    heartBeat   *HeartBeatBo
}

// 初始化AlibabaWdkMarketingOpenHeartbeatRequest对象
func NewAlibabaWdkMarketingOpenHeartbeatRequest() *AlibabaWdkMarketingOpenHeartbeatRequest{
    return &AlibabaWdkMarketingOpenHeartbeatRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingOpenHeartbeatRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.open.heartbeat"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingOpenHeartbeatRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// HeartBeat Setter
// 心跳信息
func (r *AlibabaWdkMarketingOpenHeartbeatRequest) SetHeartBeat(heartBeat *HeartBeatBo) error {
    r.heartBeat = heartBeat
    r.Set("heart_beat", heartBeat)
    return nil
}

// HeartBeat Getter
func (r AlibabaWdkMarketingOpenHeartbeatRequest) GetHeartBeat() *HeartBeatBo {
    return r.heartBeat
}

package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
心跳服务【10s一次】 
alibaba.wdk.marketing.open.heartbeat

商家数据同步心跳服务
*/
func AlibabaWdkMarketingOpenHeartbeat(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingOpenHeartbeatAPIRequest, session string) (*wdk.AlibabaWdkMarketingOpenHeartbeatAPIResponse, error) {
    var resp wdk.AlibabaWdkMarketingOpenHeartbeatAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

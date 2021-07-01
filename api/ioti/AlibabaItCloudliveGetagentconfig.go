package ioti

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ioti"
)

/* 
线上巡店Agent获取配置 
alibaba.it.cloudlive.getagentconfig

线上巡店应用，外部Agent设备获取设备配置信息，根据配置信息链接mqtt，跟云端进行进一步的消息通信。
*/
func AlibabaItCloudliveGetagentconfig(clt *core.SDKClient, req *ioti.AlibabaItCloudliveGetagentconfigAPIRequest, session string) (*ioti.AlibabaItCloudliveGetagentconfigAPIResponse, error) {
    var resp ioti.AlibabaItCloudliveGetagentconfigAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

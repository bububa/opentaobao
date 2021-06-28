package mtopopen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mtopopen"
)

/* 
社交组件 
alibaba.interact.sensor.social

赞，评论 ，关注 新增接口
*/
func AlibabaInteractSensorSocial(clt *core.SDKClient, req *mtopopen.AlibabaInteractSensorSocialRequest, session string) (*mtopopen.AlibabaInteractSensorSocialAPIResponse, error) {
    var resp mtopopen.AlibabaInteractSensorSocialAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

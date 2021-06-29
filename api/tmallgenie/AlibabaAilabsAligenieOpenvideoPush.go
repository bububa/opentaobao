package tmallgenie

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallgenie"
)

/* 
天猫精灵内容库视频分集数据推送接口 
alibaba.ailabs.aligenie.openvideo.push

天猫精灵内容库视频分集数据推送接口
*/
func AlibabaAilabsAligenieOpenvideoPush(clt *core.SDKClient, req *tmallgenie.AlibabaAilabsAligenieOpenvideoPushRequest, session string) (*tmallgenie.AlibabaAilabsAligenieOpenvideoPushAPIResponse, error) {
    var resp tmallgenie.AlibabaAilabsAligenieOpenvideoPushAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

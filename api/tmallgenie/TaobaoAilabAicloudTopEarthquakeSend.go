package tmallgenie

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallgenie"
)

/* 
地震局发送地震消息 
taobao.ailab.aicloud.top.earthquake.send

地震局发送地震消息给天猫精灵，天猫精灵根据地震消息判断发送地震消息给危险区域用户
*/
func TaobaoAilabAicloudTopEarthquakeSend(clt *core.SDKClient, req *tmallgenie.TaobaoAilabAicloudTopEarthquakeSendRequest, session string) (*tmallgenie.TaobaoAilabAicloudTopEarthquakeSendAPIResponse, error) {
    var resp tmallgenie.TaobaoAilabAicloudTopEarthquakeSendAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

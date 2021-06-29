package youkudsp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/youkudsp"
)

/* 
优酷实时批量获取可投放设备资源 
youku.dsp.delivery.resource.multiget

优酷实时获取可投放设备资源信息,为第三方渠道提供素材获取人群识别的api,支持批量获取
*/
func YoukuDspDeliveryResourceMultiget(clt *core.SDKClient, req *youkudsp.YoukuDspDeliveryResourceMultigetRequest, session string) (*youkudsp.YoukuDspDeliveryResourceMultigetAPIResponse, error) {
    var resp youkudsp.YoukuDspDeliveryResourceMultigetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

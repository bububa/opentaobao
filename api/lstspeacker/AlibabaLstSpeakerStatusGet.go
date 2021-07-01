package lstspeacker

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/lstspeacker"
)

/* 
音箱设备在线状态 
alibaba.lst.speaker.status.get

音箱设备在线状态查询
*/
func AlibabaLstSpeakerStatusGet(clt *core.SDKClient, req *lstspeacker.AlibabaLstSpeakerStatusGetAPIRequest, session string) (*lstspeacker.AlibabaLstSpeakerStatusGetAPIResponse, error) {
    var resp lstspeacker.AlibabaLstSpeakerStatusGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

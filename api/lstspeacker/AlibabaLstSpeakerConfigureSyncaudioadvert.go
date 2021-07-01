package lstspeacker

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/lstspeacker"
)

/* 
同步广告 
alibaba.lst.speaker.configure.syncaudioadvert

如意音箱广告同步
*/
func AlibabaLstSpeakerConfigureSyncaudioadvert(clt *core.SDKClient, req *lstspeacker.AlibabaLstSpeakerConfigureSyncaudioadvertAPIRequest, session string) (*lstspeacker.AlibabaLstSpeakerConfigureSyncaudioadvertAPIResponse, error) {
    var resp lstspeacker.AlibabaLstSpeakerConfigureSyncaudioadvertAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

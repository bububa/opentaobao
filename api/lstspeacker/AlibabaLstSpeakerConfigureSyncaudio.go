package lstspeacker

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/lstspeacker"
)

/* 
音频同步 
alibaba.lst.speaker.configure.syncaudio

音频同步
*/
func AlibabaLstSpeakerConfigureSyncaudio(clt *core.SDKClient, req *lstspeacker.AlibabaLstSpeakerConfigureSyncaudioRequest, session string) (*lstspeacker.AlibabaLstSpeakerConfigureSyncaudioAPIResponse, error) {
    var resp lstspeacker.AlibabaLstSpeakerConfigureSyncaudioAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

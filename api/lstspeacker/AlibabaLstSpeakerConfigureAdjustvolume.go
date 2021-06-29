package lstspeacker

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/lstspeacker"
)

/* 
音箱音量调节 
alibaba.lst.speaker.configure.adjustvolume

音箱音量调节
*/
func AlibabaLstSpeakerConfigureAdjustvolume(clt *core.SDKClient, req *lstspeacker.AlibabaLstSpeakerConfigureAdjustvolumeRequest, session string) (*lstspeacker.AlibabaLstSpeakerConfigureAdjustvolumeAPIResponse, error) {
    var resp lstspeacker.AlibabaLstSpeakerConfigureAdjustvolumeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

package lstspeacker

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/lstspeacker"
)

/* 
零售通音箱配置通用泛化调用接口 
alibaba.lst.speaker.configure

零售通音箱配置通用泛化调用接口，包括内容、音量、音频等内容
*/
func AlibabaLstSpeakerConfigure(clt *core.SDKClient, req *lstspeacker.AlibabaLstSpeakerConfigureAPIRequest, session string) (*lstspeacker.AlibabaLstSpeakerConfigureAPIResponse, error) {
    var resp lstspeacker.AlibabaLstSpeakerConfigureAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

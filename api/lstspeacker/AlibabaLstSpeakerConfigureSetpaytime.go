package lstspeacker

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstspeacker"
)

// AlibabaLstSpeakerConfigureSetpaytime 音箱播放配置
// alibaba.lst.speaker.configure.setpaytime
//
// 音箱播放配置
func AlibabaLstSpeakerConfigureSetpaytime(clt *core.SDKClient, req *lstspeacker.AlibabaLstSpeakerConfigureSetpaytimeAPIRequest, session string) (*lstspeacker.AlibabaLstSpeakerConfigureSetpaytimeAPIResponse, error) {
	var resp lstspeacker.AlibabaLstSpeakerConfigureSetpaytimeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

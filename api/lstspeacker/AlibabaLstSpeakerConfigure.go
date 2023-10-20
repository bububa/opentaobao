package lstspeacker

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstspeacker"
)

// AlibabaLstSpeakerConfigure 零售通音箱配置通用泛化调用接口
// alibaba.lst.speaker.configure
//
// 零售通音箱配置通用泛化调用接口，包括内容、音量、音频等内容
func AlibabaLstSpeakerConfigure(clt *core.SDKClient, req *lstspeacker.AlibabaLstSpeakerConfigureAPIRequest, resp *lstspeacker.AlibabaLstSpeakerConfigureAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

package lstspeacker

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstspeacker"
)

// AlibabaLstSpeakerConfigureAdjustvolume 音箱音量调节
// alibaba.lst.speaker.configure.adjustvolume
//
// 音箱音量调节
func AlibabaLstSpeakerConfigureAdjustvolume(clt *core.SDKClient, req *lstspeacker.AlibabaLstSpeakerConfigureAdjustvolumeAPIRequest, resp *lstspeacker.AlibabaLstSpeakerConfigureAdjustvolumeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

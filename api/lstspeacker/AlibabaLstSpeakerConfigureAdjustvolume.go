package lstspeacker

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstspeacker"
)

// Alibabalstspeakerconfigureadjustvolume 音箱音量调节
// alibaba.lst.speaker.configure.adjustvolume
//
// 音箱音量调节
func Alibabalstspeakerconfigureadjustvolume(clt *core.SDKClient, req *lstspeacker.AlibabalstspeakerconfigureadjustvolumeAPIRequest, session string) (*lstspeacker.AlibabalstspeakerconfigureadjustvolumeAPIResponse, error) {
	var resp lstspeacker.AlibabalstspeakerconfigureadjustvolumeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

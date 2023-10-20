package lstspeacker

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstspeacker"
)

// Alibabalstspeakerconfiguresyncaudio 音频同步
// alibaba.lst.speaker.configure.syncaudio
//
// 音频同步
func Alibabalstspeakerconfiguresyncaudio(clt *core.SDKClient, req *lstspeacker.AlibabalstspeakerconfiguresyncaudioAPIRequest, session string) (*lstspeacker.AlibabalstspeakerconfiguresyncaudioAPIResponse, error) {
	var resp lstspeacker.AlibabalstspeakerconfiguresyncaudioAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

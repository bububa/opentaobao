package lstspeacker

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstspeacker"
)

// Alibabalstspeakerconfiguresetpaytime 音箱播放配置
// alibaba.lst.speaker.configure.setpaytime
//
// 音箱播放配置
func Alibabalstspeakerconfiguresetpaytime(clt *core.SDKClient, req *lstspeacker.AlibabalstspeakerconfiguresetpaytimeAPIRequest, session string) (*lstspeacker.AlibabalstspeakerconfiguresetpaytimeAPIResponse, error) {
	var resp lstspeacker.AlibabalstspeakerconfiguresetpaytimeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

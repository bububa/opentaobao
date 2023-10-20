package lstspeacker

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstspeacker"
)

// Alibabalstspeakerconfiguresyncaudioadvert 同步广告
// alibaba.lst.speaker.configure.syncaudioadvert
//
// 如意音箱广告同步
func Alibabalstspeakerconfiguresyncaudioadvert(clt *core.SDKClient, req *lstspeacker.AlibabalstspeakerconfiguresyncaudioadvertAPIRequest, session string) (*lstspeacker.AlibabalstspeakerconfiguresyncaudioadvertAPIResponse, error) {
	var resp lstspeacker.AlibabalstspeakerconfiguresyncaudioadvertAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

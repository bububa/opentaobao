package lstspeacker

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstspeacker"
)

// Alibabalstspeakerstatusget 音箱设备在线状态
// alibaba.lst.speaker.status.get
//
// 音箱设备在线状态查询
func Alibabalstspeakerstatusget(clt *core.SDKClient, req *lstspeacker.AlibabalstspeakerstatusgetAPIRequest, session string) (*lstspeacker.AlibabalstspeakerstatusgetAPIResponse, error) {
	var resp lstspeacker.AlibabalstspeakerstatusgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

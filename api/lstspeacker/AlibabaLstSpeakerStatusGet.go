package lstspeacker

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstspeacker"
)

// AlibabaLstSpeakerStatusGet 音箱设备在线状态
// alibaba.lst.speaker.status.get
//
// 音箱设备在线状态查询
func AlibabaLstSpeakerStatusGet(clt *core.SDKClient, req *lstspeacker.AlibabaLstSpeakerStatusGetAPIRequest, resp *lstspeacker.AlibabaLstSpeakerStatusGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

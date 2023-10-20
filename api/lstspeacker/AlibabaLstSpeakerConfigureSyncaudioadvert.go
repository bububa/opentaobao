package lstspeacker

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstspeacker"
)

// AlibabaLstSpeakerConfigureSyncaudioadvert 同步广告
// alibaba.lst.speaker.configure.syncaudioadvert
//
// 如意音箱广告同步
func AlibabaLstSpeakerConfigureSyncaudioadvert(clt *core.SDKClient, req *lstspeacker.AlibabaLstSpeakerConfigureSyncaudioadvertAPIRequest, resp *lstspeacker.AlibabaLstSpeakerConfigureSyncaudioadvertAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

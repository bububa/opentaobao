package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Alibabaailabsaligenieopencontentscenepush 音频场景接入接口
// alibaba.ailabs.aligenie.opencontent.scenepush
//
// 天猫精灵音频挂靠场景接入
func Alibabaailabsaligenieopencontentscenepush(clt *core.SDKClient, req *tmallgenie.AlibabaailabsaligenieopencontentscenepushAPIRequest, session string) (*tmallgenie.AlibabaailabsaligenieopencontentscenepushAPIResponse, error) {
	var resp tmallgenie.AlibabaailabsaligenieopencontentscenepushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

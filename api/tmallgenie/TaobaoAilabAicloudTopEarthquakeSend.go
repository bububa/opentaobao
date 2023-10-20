package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// TaobaoAilabAicloudTopEarthquakeSend 地震局发送地震消息
// taobao.ailab.aicloud.top.earthquake.send
//
// 地震局发送地震消息给天猫精灵，天猫精灵根据地震消息判断发送地震消息给危险区域用户
func TaobaoAilabAicloudTopEarthquakeSend(clt *core.SDKClient, req *tmallgenie.TaobaoAilabAicloudTopEarthquakeSendAPIRequest, resp *tmallgenie.TaobaoAilabAicloudTopEarthquakeSendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

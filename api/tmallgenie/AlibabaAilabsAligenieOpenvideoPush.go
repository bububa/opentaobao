package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsAligenieOpenvideoPush 天猫精灵内容库视频分集数据推送接口
// alibaba.ailabs.aligenie.openvideo.push
//
// 天猫精灵内容库视频分集数据推送接口
func AlibabaAilabsAligenieOpenvideoPush(clt *core.SDKClient, req *tmallgenie.AlibabaAilabsAligenieOpenvideoPushAPIRequest, resp *tmallgenie.AlibabaAilabsAligenieOpenvideoPushAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

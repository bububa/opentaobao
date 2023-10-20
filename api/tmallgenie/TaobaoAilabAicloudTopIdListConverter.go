package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// TaobaoAilabAicloudTopIdListConverter 将淘宝openId或者设备id/用户id转换为openId
// taobao.ailab.aicloud.top.id.list.converter
//
// 将淘宝openId或者设备id/用户id转换为openId
func TaobaoAilabAicloudTopIdListConverter(clt *core.SDKClient, req *tmallgenie.TaobaoAilabAicloudTopIdListConverterAPIRequest, resp *tmallgenie.TaobaoAilabAicloudTopIdListConverterAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

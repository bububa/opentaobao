package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Taobaoailabaicloudtopearthquakesend 地震局发送地震消息
// taobao.ailab.aicloud.top.earthquake.send
//
// 地震局发送地震消息给天猫精灵，天猫精灵根据地震消息判断发送地震消息给危险区域用户
func Taobaoailabaicloudtopearthquakesend(clt *core.SDKClient, req *tmallgenie.TaobaoailabaicloudtopearthquakesendAPIRequest, session string) (*tmallgenie.TaobaoailabaicloudtopearthquakesendAPIResponse, error) {
	var resp tmallgenie.TaobaoailabaicloudtopearthquakesendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

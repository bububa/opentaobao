package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Taobaoailabaicloudtopmemoalarmlist 天猫精灵闹钟查询
// taobao.ailab.aicloud.top.memo.alarm.list
//
// 查询天猫精灵用户设置的所有闹钟
func Taobaoailabaicloudtopmemoalarmlist(clt *core.SDKClient, req *tmallgenie.TaobaoailabaicloudtopmemoalarmlistAPIRequest, session string) (*tmallgenie.TaobaoailabaicloudtopmemoalarmlistAPIResponse, error) {
	var resp tmallgenie.TaobaoailabaicloudtopmemoalarmlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

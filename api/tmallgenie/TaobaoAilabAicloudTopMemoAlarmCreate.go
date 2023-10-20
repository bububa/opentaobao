package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Taobaoailabaicloudtopmemoalarmcreate 天猫精灵闹钟创建
// taobao.ailab.aicloud.top.memo.alarm.create
//
// 天猫精灵闹钟创建
func Taobaoailabaicloudtopmemoalarmcreate(clt *core.SDKClient, req *tmallgenie.TaobaoailabaicloudtopmemoalarmcreateAPIRequest, session string) (*tmallgenie.TaobaoailabaicloudtopmemoalarmcreateAPIResponse, error) {
	var resp tmallgenie.TaobaoailabaicloudtopmemoalarmcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

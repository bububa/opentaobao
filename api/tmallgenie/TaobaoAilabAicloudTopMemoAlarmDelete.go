package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Taobaoailabaicloudtopmemoalarmdelete 天猫精灵闹钟删除
// taobao.ailab.aicloud.top.memo.alarm.delete
//
// 天猫精灵闹钟删除
func Taobaoailabaicloudtopmemoalarmdelete(clt *core.SDKClient, req *tmallgenie.TaobaoailabaicloudtopmemoalarmdeleteAPIRequest, session string) (*tmallgenie.TaobaoailabaicloudtopmemoalarmdeleteAPIResponse, error) {
	var resp tmallgenie.TaobaoailabaicloudtopmemoalarmdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

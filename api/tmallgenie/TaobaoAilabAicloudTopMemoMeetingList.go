package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Taobaoailabaicloudtopmemomeetinglist 天猫精灵会议查询
// taobao.ailab.aicloud.top.memo.meeting.list
//
// 查询天猫精灵用户设置的所有会议
func Taobaoailabaicloudtopmemomeetinglist(clt *core.SDKClient, req *tmallgenie.TaobaoailabaicloudtopmemomeetinglistAPIRequest, session string) (*tmallgenie.TaobaoailabaicloudtopmemomeetinglistAPIResponse, error) {
	var resp tmallgenie.TaobaoailabaicloudtopmemomeetinglistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Taobaoailabaicloudtopmemomeetingdelete 天猫精灵会议删除
// taobao.ailab.aicloud.top.memo.meeting.delete
//
// 天猫精灵会议删除
func Taobaoailabaicloudtopmemomeetingdelete(clt *core.SDKClient, req *tmallgenie.TaobaoailabaicloudtopmemomeetingdeleteAPIRequest, session string) (*tmallgenie.TaobaoailabaicloudtopmemomeetingdeleteAPIResponse, error) {
	var resp tmallgenie.TaobaoailabaicloudtopmemomeetingdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

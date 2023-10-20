package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// TaobaoAilabAicloudTopMemoMeetingDelete 天猫精灵会议删除
// taobao.ailab.aicloud.top.memo.meeting.delete
//
// 天猫精灵会议删除
func TaobaoAilabAicloudTopMemoMeetingDelete(clt *core.SDKClient, req *tmallgenie.TaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest, resp *tmallgenie.TaobaoAilabAicloudTopMemoMeetingDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

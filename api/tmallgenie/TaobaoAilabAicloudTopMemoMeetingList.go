package tmallgenie

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// TaobaoAilabAicloudTopMemoMeetingList 天猫精灵会议查询
// taobao.ailab.aicloud.top.memo.meeting.list
//
// 查询天猫精灵用户设置的所有会议
func TaobaoAilabAicloudTopMemoMeetingList(ctx context.Context, clt *core.SDKClient, req *tmallgenie.TaobaoAilabAicloudTopMemoMeetingListAPIRequest, resp *tmallgenie.TaobaoAilabAicloudTopMemoMeetingListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

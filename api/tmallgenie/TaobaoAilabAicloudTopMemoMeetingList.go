package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// TaobaoAilabAicloudTopMemoMeetingList 天猫精灵会议查询
// taobao.ailab.aicloud.top.memo.meeting.list
//
// 查询天猫精灵用户设置的所有会议
func TaobaoAilabAicloudTopMemoMeetingList(clt *core.SDKClient, req *tmallgenie.TaobaoAilabAicloudTopMemoMeetingListAPIRequest, session string) (*tmallgenie.TaobaoAilabAicloudTopMemoMeetingListAPIResponse, error) {
	var resp tmallgenie.TaobaoAilabAicloudTopMemoMeetingListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

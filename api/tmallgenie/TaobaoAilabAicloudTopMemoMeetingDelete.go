package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

/* TaobaoAilabAicloudTopMemoMeetingDelete
天猫精灵会议删除
taobao.ailab.aicloud.top.memo.meeting.delete

天猫精灵会议删除 */
func TaobaoAilabAicloudTopMemoMeetingDelete(clt *core.SDKClient, req *tmallgenie.TaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest, session string) (*tmallgenie.TaobaoAilabAicloudTopMemoMeetingDeleteAPIResponse, error) {
	var resp tmallgenie.TaobaoAilabAicloudTopMemoMeetingDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

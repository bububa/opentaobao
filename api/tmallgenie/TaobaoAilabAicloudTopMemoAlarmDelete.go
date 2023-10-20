package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// TaobaoAilabAicloudTopMemoAlarmDelete 天猫精灵闹钟删除
// taobao.ailab.aicloud.top.memo.alarm.delete
//
// 天猫精灵闹钟删除
func TaobaoAilabAicloudTopMemoAlarmDelete(clt *core.SDKClient, req *tmallgenie.TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest, session string) (*tmallgenie.TaobaoAilabAicloudTopMemoAlarmDeleteAPIResponse, error) {
	var resp tmallgenie.TaobaoAilabAicloudTopMemoAlarmDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// TaobaoAilabAicloudTopMemoAlarmDelete 天猫精灵闹钟删除
// taobao.ailab.aicloud.top.memo.alarm.delete
//
// 天猫精灵闹钟删除
func TaobaoAilabAicloudTopMemoAlarmDelete(clt *core.SDKClient, req *tmallgenie.TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest, resp *tmallgenie.TaobaoAilabAicloudTopMemoAlarmDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

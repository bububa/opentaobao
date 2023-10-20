package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// TaobaoAilabAicloudTopMemoAlarmList 天猫精灵闹钟查询
// taobao.ailab.aicloud.top.memo.alarm.list
//
// 查询天猫精灵用户设置的所有闹钟
func TaobaoAilabAicloudTopMemoAlarmList(clt *core.SDKClient, req *tmallgenie.TaobaoAilabAicloudTopMemoAlarmListAPIRequest, resp *tmallgenie.TaobaoAilabAicloudTopMemoAlarmListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

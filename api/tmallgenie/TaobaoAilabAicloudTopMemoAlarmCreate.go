package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// TaobaoAilabAicloudTopMemoAlarmCreate 天猫精灵闹钟创建
// taobao.ailab.aicloud.top.memo.alarm.create
//
// 天猫精灵闹钟创建
func TaobaoAilabAicloudTopMemoAlarmCreate(clt *core.SDKClient, req *tmallgenie.TaobaoAilabAicloudTopMemoAlarmCreateAPIRequest, session string) (*tmallgenie.TaobaoAilabAicloudTopMemoAlarmCreateAPIResponse, error) {
	var resp tmallgenie.TaobaoAilabAicloudTopMemoAlarmCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

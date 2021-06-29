package tmallgenie

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallgenie"
)

/* 
天猫精灵闹钟查询 
taobao.ailab.aicloud.top.memo.alarm.list

查询天猫精灵用户设置的所有闹钟
*/
func TaobaoAilabAicloudTopMemoAlarmList(clt *core.SDKClient, req *tmallgenie.TaobaoAilabAicloudTopMemoAlarmListRequest, session string) (*tmallgenie.TaobaoAilabAicloudTopMemoAlarmListAPIResponse, error) {
    var resp tmallgenie.TaobaoAilabAicloudTopMemoAlarmListAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

package pentraprism

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pentraprism"
)

// TaobaoPentaprismTaskTriggerFrom 任务进度推进（根据fromtoken）
// taobao.pentaprism.task.trigger.from
//
// 外网用户推进单条五棱镜任务进度
func TaobaoPentaprismTaskTriggerFrom(clt *core.SDKClient, req *pentraprism.TaobaoPentaprismTaskTriggerFromAPIRequest, session string) (*pentraprism.TaobaoPentaprismTaskTriggerFromAPIResponse, error) {
	var resp pentraprism.TaobaoPentaprismTaskTriggerFromAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

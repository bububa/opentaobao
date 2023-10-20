package pentraprism

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pentraprism"
)

// TaobaoPentaprismTaskTriggerFrom 任务进度推进（根据fromtoken）
// taobao.pentaprism.task.trigger.from
//
// 外网用户推进单条五棱镜任务进度
func TaobaoPentaprismTaskTriggerFrom(clt *core.SDKClient, req *pentraprism.TaobaoPentaprismTaskTriggerFromAPIRequest, resp *pentraprism.TaobaoPentaprismTaskTriggerFromAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

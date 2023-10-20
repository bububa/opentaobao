package pentraprism

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pentraprism"
)

// TaobaoPentaprismTaskTrigger 推进单条任务进度
// taobao.pentaprism.task.trigger
//
// 外网用户推进单条五棱镜任务进度
func TaobaoPentaprismTaskTrigger(clt *core.SDKClient, req *pentraprism.TaobaoPentaprismTaskTriggerAPIRequest, resp *pentraprism.TaobaoPentaprismTaskTriggerAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

package pentraprism

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pentraprism"
)

/* TaobaoPentaprismTaskTrigger
推进单条任务进度
taobao.pentaprism.task.trigger

外网用户推进单条五棱镜任务进度 */
func TaobaoPentaprismTaskTrigger(clt *core.SDKClient, req *pentraprism.TaobaoPentaprismTaskTriggerAPIRequest, session string) (*pentraprism.TaobaoPentaprismTaskTriggerAPIResponse, error) {
	var resp pentraprism.TaobaoPentaprismTaskTriggerAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

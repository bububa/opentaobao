package pentraprism

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPentaprismTaskTriggerAPIRequest
推进单条任务进度 API请求
taobao.pentaprism.task.trigger

外网用户推进单条五棱镜任务进度 */
type TaobaoPentaprismTaskTriggerAPIRequest struct {
	model.Params
	// TOP接口标准入参
	_openPo *OpenTaskPo
}

// New

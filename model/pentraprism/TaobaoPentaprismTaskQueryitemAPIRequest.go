package pentraprism

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPentaprismTaskQueryitemAPIRequest
查询任务当前进度 API请求
taobao.pentaprism.task.queryitem

外网用户查询五棱镜任务系统当前进度 */
type TaobaoPentaprismTaskQueryitemAPIRequest struct {
	model.Params
	// TOP接口标准入参
	_openPo *OpenTaskPo
}

// New

package itpolicy

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripItFareQuerytaskAPIRequest
【国际机票自有政策】批量操作结果查询 API请求
taobao.alitrip.it.fare.querytask

批量操作同步返回任务id，后台生成异步任务，通过此接口查询批量操作的执行结果 */
type TaobaoAlitripItFareQuerytaskAPIRequest struct {
	model.Params
	// json格式的字符串，扩展属性，预留
	_extendAttributes string
	// 任务id
	_taskId int64
}

// New

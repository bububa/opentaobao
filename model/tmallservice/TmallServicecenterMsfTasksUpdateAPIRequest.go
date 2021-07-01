package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterMsfTasksUpdateAPIRequest
喵师傅工人任务批量完成接口 API请求
tmall.servicecenter.msf.tasks.update

喵师傅工人任务批量完成接口 */
type TmallServicecenterMsfTasksUpdateAPIRequest struct {
	model.Params
	// 工人手机号
	_workerMobile int64
	// 服务编码
	_serviceCode string
	// 调用来源
	_source string
	// 子订单号列表。最多100个
	_bizOrderIds []int64
	// 周期序号。必须与子订单列表对应
	_seqs []int64
}

// New

package tmallfcbox

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallFcboxNotifyAPIRequest
丰巢通知接口 API请求
tmall.fcbox.notify

tmax接收丰巢快递通知 */
type TmallFcboxNotifyAPIRequest struct {
	model.Params
	// 申请接口返回的申请标识
	_applyId string
	// 状态备注
	_stateRemark string
	// 申请记录当前状态
	_state string
	// 变更时间
	_operateTime string
	// 变更操作
	_operate string
}

// New

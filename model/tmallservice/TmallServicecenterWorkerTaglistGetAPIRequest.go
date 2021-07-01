package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterWorkerTaglistGetAPIRequest
获取工人标签 API请求
tmall.servicecenter.worker.taglist.get

服务商获取对应工人的标签 */
type TmallServicecenterWorkerTaglistGetAPIRequest struct {
	model.Params
	// 工人注册勤鸽时的身份证
	_idNumber string
	// 工人注册勤鸽时的手机号码
	_mobile string
}

// New

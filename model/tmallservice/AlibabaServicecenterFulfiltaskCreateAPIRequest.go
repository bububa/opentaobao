package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaServicecenterFulfiltaskCreateAPIRequest
合单生成核销单 API请求
alibaba.servicecenter.fulfiltask.create

服务对工单进行合单，合单的结果是生成核销单 */
type AlibabaServicecenterFulfiltaskCreateAPIRequest struct {
	model.Params
	// 工单id列表
	_workcardIds []int64
	// 外部单号
	_outerId string
}

// New

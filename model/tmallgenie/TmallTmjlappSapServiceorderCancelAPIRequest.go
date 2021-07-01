package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallTmjlappSapServiceorderCancelAPIRequest
取消售后服务单 API请求
tmall.tmjlapp.sap.serviceorder.cancel

SAP跟天猫精灵app接口对接，用户在app取消sap售后服务工单 */
type TmallTmjlappSapServiceorderCancelAPIRequest struct {
	model.Params
	// 取消服务单请求
	_cancelRequest *Dtcancelrequest
}

// New

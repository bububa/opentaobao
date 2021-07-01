package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTclsAelophyBillVerificateCallbackAPIRequest
翱象ERP核销回调 API请求
alibaba.tcls.aelophy.bill.verificate.callback

翱象ERP核销回调 */
type AlibabaTclsAelophyBillVerificateCallbackAPIRequest struct {
	model.Params
	// 回调对象
	_module *VerificateCallbackDto
}

// New

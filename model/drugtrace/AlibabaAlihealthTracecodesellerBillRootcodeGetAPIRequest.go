package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest
获取最外层包装码 API请求
alibaba.alihealth.tracecodeseller.bill.rootcode.get

获取最外层包装码 */
type AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest struct {
	model.Params
	// 用户身份认证
	_appCode string
	// 码
	_code string
}

// New

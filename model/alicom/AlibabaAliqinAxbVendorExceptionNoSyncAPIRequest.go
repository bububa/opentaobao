package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest
中心化供应商异常号码状态同步接口 API请求
alibaba.aliqin.axb.vendor.exception.no.sync

用于中心化供应商同步异常号码 */
type AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest struct {
	model.Params
	// 异常的中间号码
	_secretNo string
	// 异常的原因
	_exceptionMsg string
	// 0-异常状态 1-可恢复正常使用
	_status int64
	// 供应商KEY
	_vendorKey string
}

// New

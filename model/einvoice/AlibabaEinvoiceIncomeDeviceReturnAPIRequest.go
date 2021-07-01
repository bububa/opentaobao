package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceIncomeDeviceReturnAPIRequest
服务商回传客户端设备列表 API请求
alibaba.einvoice.income.device.return

服务商回传客户端agent所处环境的设备列表，比如扫描仪 */
type AlibabaEinvoiceIncomeDeviceReturnAPIRequest struct {
	model.Params
	// 设备列表，success=true时必填
	_deviceList []string
	// 错误码，success=false时必填
	_errorCode string
	// 错误信息，success=false时必填
	_errorMessage string
	// 请求标识
	_reqIndex string
	// 查询设备是否成功，true=成功，false=失败
	_success bool
}

// New

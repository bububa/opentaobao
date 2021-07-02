package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceIncomeDeviceReturnAPIRequest 服务商回传客户端设备列表 API请求
// alibaba.einvoice.income.device.return
//
// 服务商回传客户端agent所处环境的设备列表，比如扫描仪
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

// NewAlibabaEinvoiceIncomeDeviceReturnRequest 初始化AlibabaEinvoiceIncomeDeviceReturnAPIRequest对象
func NewAlibabaEinvoiceIncomeDeviceReturnRequest() *AlibabaEinvoiceIncomeDeviceReturnAPIRequest {
	return &AlibabaEinvoiceIncomeDeviceReturnAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceIncomeDeviceReturnAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.income.device.return"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceIncomeDeviceReturnAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is DeviceList Setter
// 设备列表，success=true时必填
func (r *AlibabaEinvoiceIncomeDeviceReturnAPIRequest) SetDeviceList(_deviceList []string) error {
	r._deviceList = _deviceList
	r.Set("device_list", _deviceList)
	return nil
}

// Get DeviceList Getter
func (r AlibabaEinvoiceIncomeDeviceReturnAPIRequest) GetDeviceList() []string {
	return r._deviceList
}

// Set is ErrorCode Setter
// 错误码，success=false时必填
func (r *AlibabaEinvoiceIncomeDeviceReturnAPIRequest) SetErrorCode(_errorCode string) error {
	r._errorCode = _errorCode
	r.Set("error_code", _errorCode)
	return nil
}

// Get ErrorCode Getter
func (r AlibabaEinvoiceIncomeDeviceReturnAPIRequest) GetErrorCode() string {
	return r._errorCode
}

// Set is ErrorMessage Setter
// 错误信息，success=false时必填
func (r *AlibabaEinvoiceIncomeDeviceReturnAPIRequest) SetErrorMessage(_errorMessage string) error {
	r._errorMessage = _errorMessage
	r.Set("error_message", _errorMessage)
	return nil
}

// Get ErrorMessage Getter
func (r AlibabaEinvoiceIncomeDeviceReturnAPIRequest) GetErrorMessage() string {
	return r._errorMessage
}

// Set is ReqIndex Setter
// 请求标识
func (r *AlibabaEinvoiceIncomeDeviceReturnAPIRequest) SetReqIndex(_reqIndex string) error {
	r._reqIndex = _reqIndex
	r.Set("req_index", _reqIndex)
	return nil
}

// Get ReqIndex Getter
func (r AlibabaEinvoiceIncomeDeviceReturnAPIRequest) GetReqIndex() string {
	return r._reqIndex
}

// Set is Success Setter
// 查询设备是否成功，true=成功，false=失败
func (r *AlibabaEinvoiceIncomeDeviceReturnAPIRequest) SetSuccess(_success bool) error {
	r._success = _success
	r.Set("success", _success)
	return nil
}

// Get Success Getter
func (r AlibabaEinvoiceIncomeDeviceReturnAPIRequest) GetSuccess() bool {
	return r._success
}

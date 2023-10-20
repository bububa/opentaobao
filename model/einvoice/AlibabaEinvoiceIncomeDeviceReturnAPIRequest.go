package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoiceincomedevicereturnAPIRequest 服务商回传客户端设备列表 API请求
// alibaba.einvoice.income.device.return
//
// 服务商回传客户端agent所处环境的设备列表，比如扫描仪
type AlibabaeinvoiceincomedevicereturnAPIRequest struct {
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

// NewAlibabaeinvoiceincomedevicereturnRequest 初始化AlibabaeinvoiceincomedevicereturnAPIRequest对象
func NewAlibabaeinvoiceincomedevicereturnRequest() *AlibabaeinvoiceincomedevicereturnAPIRequest {
	return &AlibabaeinvoiceincomedevicereturnAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoiceincomedevicereturnAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.income.device.return"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoiceincomedevicereturnAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoiceincomedevicereturnAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceList is DeviceList Setter
// 设备列表，success=true时必填
func (r *AlibabaeinvoiceincomedevicereturnAPIRequest) SetDeviceList(_deviceList []string) error {
	r._deviceList = _deviceList
	r.Set("device_list", _deviceList)
	return nil
}

// GetDeviceList DeviceList Getter
func (r AlibabaeinvoiceincomedevicereturnAPIRequest) GetDeviceList() []string {
	return r._deviceList
}

// SetErrorCode is ErrorCode Setter
// 错误码，success=false时必填
func (r *AlibabaeinvoiceincomedevicereturnAPIRequest) SetErrorCode(_errorCode string) error {
	r._errorCode = _errorCode
	r.Set("error_code", _errorCode)
	return nil
}

// GetErrorCode ErrorCode Getter
func (r AlibabaeinvoiceincomedevicereturnAPIRequest) GetErrorCode() string {
	return r._errorCode
}

// SetErrorMessage is ErrorMessage Setter
// 错误信息，success=false时必填
func (r *AlibabaeinvoiceincomedevicereturnAPIRequest) SetErrorMessage(_errorMessage string) error {
	r._errorMessage = _errorMessage
	r.Set("error_message", _errorMessage)
	return nil
}

// GetErrorMessage ErrorMessage Getter
func (r AlibabaeinvoiceincomedevicereturnAPIRequest) GetErrorMessage() string {
	return r._errorMessage
}

// SetReqIndex is ReqIndex Setter
// 请求标识
func (r *AlibabaeinvoiceincomedevicereturnAPIRequest) SetReqIndex(_reqIndex string) error {
	r._reqIndex = _reqIndex
	r.Set("req_index", _reqIndex)
	return nil
}

// GetReqIndex ReqIndex Getter
func (r AlibabaeinvoiceincomedevicereturnAPIRequest) GetReqIndex() string {
	return r._reqIndex
}

// SetSuccess is Success Setter
// 查询设备是否成功，true=成功，false=失败
func (r *AlibabaeinvoiceincomedevicereturnAPIRequest) SetSuccess(_success bool) error {
	r._success = _success
	r.Set("success", _success)
	return nil
}

// GetSuccess Success Getter
func (r AlibabaeinvoiceincomedevicereturnAPIRequest) GetSuccess() bool {
	return r._success
}

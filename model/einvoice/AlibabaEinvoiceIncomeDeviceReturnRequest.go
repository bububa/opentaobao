package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商回传客户端设备列表 API请求
alibaba.einvoice.income.device.return

服务商回传客户端agent所处环境的设备列表，比如扫描仪
*/
type AlibabaEinvoiceIncomeDeviceReturnRequest struct {
    model.Params
    // 设备列表，success=true时必填
    deviceList   []string
    // 错误码，success=false时必填
    errorCode   string
    // 错误信息，success=false时必填
    errorMessage   string
    // 请求标识
    reqIndex   string
    // 查询设备是否成功，true=成功，false=失败
    success   bool
}

// 初始化AlibabaEinvoiceIncomeDeviceReturnRequest对象
func NewAlibabaEinvoiceIncomeDeviceReturnRequest() *AlibabaEinvoiceIncomeDeviceReturnRequest{
    return &AlibabaEinvoiceIncomeDeviceReturnRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceIncomeDeviceReturnRequest) GetApiMethodName() string {
    return "alibaba.einvoice.income.device.return"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceIncomeDeviceReturnRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceList Setter
// 设备列表，success=true时必填
func (r *AlibabaEinvoiceIncomeDeviceReturnRequest) SetDeviceList(deviceList []string) error {
    r.deviceList = deviceList
    r.Set("device_list", deviceList)
    return nil
}

// DeviceList Getter
func (r AlibabaEinvoiceIncomeDeviceReturnRequest) GetDeviceList() []string {
    return r.deviceList
}
// ErrorCode Setter
// 错误码，success=false时必填
func (r *AlibabaEinvoiceIncomeDeviceReturnRequest) SetErrorCode(errorCode string) error {
    r.errorCode = errorCode
    r.Set("error_code", errorCode)
    return nil
}

// ErrorCode Getter
func (r AlibabaEinvoiceIncomeDeviceReturnRequest) GetErrorCode() string {
    return r.errorCode
}
// ErrorMessage Setter
// 错误信息，success=false时必填
func (r *AlibabaEinvoiceIncomeDeviceReturnRequest) SetErrorMessage(errorMessage string) error {
    r.errorMessage = errorMessage
    r.Set("error_message", errorMessage)
    return nil
}

// ErrorMessage Getter
func (r AlibabaEinvoiceIncomeDeviceReturnRequest) GetErrorMessage() string {
    return r.errorMessage
}
// ReqIndex Setter
// 请求标识
func (r *AlibabaEinvoiceIncomeDeviceReturnRequest) SetReqIndex(reqIndex string) error {
    r.reqIndex = reqIndex
    r.Set("req_index", reqIndex)
    return nil
}

// ReqIndex Getter
func (r AlibabaEinvoiceIncomeDeviceReturnRequest) GetReqIndex() string {
    return r.reqIndex
}
// Success Setter
// 查询设备是否成功，true=成功，false=失败
func (r *AlibabaEinvoiceIncomeDeviceReturnRequest) SetSuccess(success bool) error {
    r.success = success
    r.Set("success", success)
    return nil
}

// Success Getter
func (r AlibabaEinvoiceIncomeDeviceReturnRequest) GetSuccess() bool {
    return r.success
}

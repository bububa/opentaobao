package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询扫码信息 API请求
alibaba.alihealth.drugcode.scan

查询扫码信息
*/
type AlibabaAlihealthDrugcodeScanRequest struct {
    model.Params
    // 20位码
    _code   string
    // 渠道
    _queryAppName   string
    // 用户ip
    _clientId   string
    // 设备标识
    _deviceUtdid   string
    // 用户ID
    _userId   string
}

// 初始化AlibabaAlihealthDrugcodeScanRequest对象
func NewAlibabaAlihealthDrugcodeScanRequest() *AlibabaAlihealthDrugcodeScanRequest{
    return &AlibabaAlihealthDrugcodeScanRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeScanRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugcode.scan"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeScanRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Code Setter
// 20位码
func (r *AlibabaAlihealthDrugcodeScanRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r AlibabaAlihealthDrugcodeScanRequest) GetCode() string {
    return r._code
}
// QueryAppName Setter
// 渠道
func (r *AlibabaAlihealthDrugcodeScanRequest) SetQueryAppName(_queryAppName string) error {
    r._queryAppName = _queryAppName
    r.Set("query_app_name", _queryAppName)
    return nil
}

// QueryAppName Getter
func (r AlibabaAlihealthDrugcodeScanRequest) GetQueryAppName() string {
    return r._queryAppName
}
// ClientId Setter
// 用户ip
func (r *AlibabaAlihealthDrugcodeScanRequest) SetClientId(_clientId string) error {
    r._clientId = _clientId
    r.Set("client_id", _clientId)
    return nil
}

// ClientId Getter
func (r AlibabaAlihealthDrugcodeScanRequest) GetClientId() string {
    return r._clientId
}
// DeviceUtdid Setter
// 设备标识
func (r *AlibabaAlihealthDrugcodeScanRequest) SetDeviceUtdid(_deviceUtdid string) error {
    r._deviceUtdid = _deviceUtdid
    r.Set("device_utdid", _deviceUtdid)
    return nil
}

// DeviceUtdid Getter
func (r AlibabaAlihealthDrugcodeScanRequest) GetDeviceUtdid() string {
    return r._deviceUtdid
}
// UserId Setter
// 用户ID
func (r *AlibabaAlihealthDrugcodeScanRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaAlihealthDrugcodeScanRequest) GetUserId() string {
    return r._userId
}

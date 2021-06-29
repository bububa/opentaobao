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
    code   string
    // 渠道
    queryAppName   string
    // 用户ip
    clientId   string
    // 设备标识
    deviceUtdid   string
    // 用户ID
    userId   string
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
func (r *AlibabaAlihealthDrugcodeScanRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

// Code Getter
func (r AlibabaAlihealthDrugcodeScanRequest) GetCode() string {
    return r.code
}
// QueryAppName Setter
// 渠道
func (r *AlibabaAlihealthDrugcodeScanRequest) SetQueryAppName(queryAppName string) error {
    r.queryAppName = queryAppName
    r.Set("query_app_name", queryAppName)
    return nil
}

// QueryAppName Getter
func (r AlibabaAlihealthDrugcodeScanRequest) GetQueryAppName() string {
    return r.queryAppName
}
// ClientId Setter
// 用户ip
func (r *AlibabaAlihealthDrugcodeScanRequest) SetClientId(clientId string) error {
    r.clientId = clientId
    r.Set("client_id", clientId)
    return nil
}

// ClientId Getter
func (r AlibabaAlihealthDrugcodeScanRequest) GetClientId() string {
    return r.clientId
}
// DeviceUtdid Setter
// 设备标识
func (r *AlibabaAlihealthDrugcodeScanRequest) SetDeviceUtdid(deviceUtdid string) error {
    r.deviceUtdid = deviceUtdid
    r.Set("device_utdid", deviceUtdid)
    return nil
}

// DeviceUtdid Getter
func (r AlibabaAlihealthDrugcodeScanRequest) GetDeviceUtdid() string {
    return r.deviceUtdid
}
// UserId Setter
// 用户ID
func (r *AlibabaAlihealthDrugcodeScanRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaAlihealthDrugcodeScanRequest) GetUserId() string {
    return r.userId
}

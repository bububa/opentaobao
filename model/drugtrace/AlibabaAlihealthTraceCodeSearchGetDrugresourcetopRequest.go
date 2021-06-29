package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据码获取码信息 API请求
alibaba.alihealth.trace.code.search.get.drugresourcetop

根据码获取码信息
*/
type AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest struct {
    model.Params
    // 追溯码
    _code   string
    // 校验值
    _token   string
    // 查询app名称
    _queryAppName   string
    // 客户端ip
    _clientId   string
    // 用户id
    _tbUserId   int64
    // 设备号
    _deviceUtdid   string
}

// 初始化AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest对象
func NewAlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest() *AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest{
    return &AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest) GetApiMethodName() string {
    return "alibaba.alihealth.trace.code.search.get.drugresourcetop"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Code Setter
// 追溯码
func (r *AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest) GetCode() string {
    return r._code
}
// Token Setter
// 校验值
func (r *AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest) GetToken() string {
    return r._token
}
// QueryAppName Setter
// 查询app名称
func (r *AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest) SetQueryAppName(_queryAppName string) error {
    r._queryAppName = _queryAppName
    r.Set("query_app_name", _queryAppName)
    return nil
}

// QueryAppName Getter
func (r AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest) GetQueryAppName() string {
    return r._queryAppName
}
// ClientId Setter
// 客户端ip
func (r *AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest) SetClientId(_clientId string) error {
    r._clientId = _clientId
    r.Set("client_id", _clientId)
    return nil
}

// ClientId Getter
func (r AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest) GetClientId() string {
    return r._clientId
}
// TbUserId Setter
// 用户id
func (r *AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest) SetTbUserId(_tbUserId int64) error {
    r._tbUserId = _tbUserId
    r.Set("tb_user_id", _tbUserId)
    return nil
}

// TbUserId Getter
func (r AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest) GetTbUserId() int64 {
    return r._tbUserId
}
// DeviceUtdid Setter
// 设备号
func (r *AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest) SetDeviceUtdid(_deviceUtdid string) error {
    r._deviceUtdid = _deviceUtdid
    r.Set("device_utdid", _deviceUtdid)
    return nil
}

// DeviceUtdid Getter
func (r AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest) GetDeviceUtdid() string {
    return r._deviceUtdid
}

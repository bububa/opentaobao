package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据码获取码信息 APIRequest
alibaba.alihealth.trace.code.search.get.drugresourcetop

根据码获取码信息
*/
type AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest struct {
    model.Params

    // 追溯码
    code   string 

    // 校验值
    token   string 

    // 查询app名称
    queryAppName   string 

    // 客户端ip
    clientId   string 

    // 用户id
    tbUserId   int64 

    // 设备号
    deviceUtdid   string 

}

func NewAlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest() *AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest{
    return &AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest) GetApiMethodName() string {
    return "alibaba.alihealth.trace.code.search.get.drugresourcetop"
}

func (r AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest) GetCode() string {
    return r.code
}

func (r *AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest) GetToken() string {
    return r.token
}

func (r *AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest) SetQueryAppName(queryAppName string) error {
    r.queryAppName = queryAppName
    r.Set("query_app_name", queryAppName)
    return nil
}

func (r AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest) GetQueryAppName() string {
    return r.queryAppName
}

func (r *AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest) SetClientId(clientId string) error {
    r.clientId = clientId
    r.Set("client_id", clientId)
    return nil
}

func (r AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest) GetClientId() string {
    return r.clientId
}

func (r *AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest) SetTbUserId(tbUserId int64) error {
    r.tbUserId = tbUserId
    r.Set("tb_user_id", tbUserId)
    return nil
}

func (r AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest) GetTbUserId() int64 {
    return r.tbUserId
}

func (r *AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest) SetDeviceUtdid(deviceUtdid string) error {
    r.deviceUtdid = deviceUtdid
    r.Set("device_utdid", deviceUtdid)
    return nil
}

func (r AlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest) GetDeviceUtdid() string {
    return r.deviceUtdid
}


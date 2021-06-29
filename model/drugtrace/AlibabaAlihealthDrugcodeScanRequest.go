package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询扫码信息 APIRequest
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

func NewAlibabaAlihealthDrugcodeScanRequest() *AlibabaAlihealthDrugcodeScanRequest{
    return &AlibabaAlihealthDrugcodeScanRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugcodeScanRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugcode.scan"
}

func (r AlibabaAlihealthDrugcodeScanRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugcodeScanRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r AlibabaAlihealthDrugcodeScanRequest) GetCode() string {
    return r.code
}

func (r *AlibabaAlihealthDrugcodeScanRequest) SetQueryAppName(queryAppName string) error {
    r.queryAppName = queryAppName
    r.Set("query_app_name", queryAppName)
    return nil
}

func (r AlibabaAlihealthDrugcodeScanRequest) GetQueryAppName() string {
    return r.queryAppName
}

func (r *AlibabaAlihealthDrugcodeScanRequest) SetClientId(clientId string) error {
    r.clientId = clientId
    r.Set("client_id", clientId)
    return nil
}

func (r AlibabaAlihealthDrugcodeScanRequest) GetClientId() string {
    return r.clientId
}

func (r *AlibabaAlihealthDrugcodeScanRequest) SetDeviceUtdid(deviceUtdid string) error {
    r.deviceUtdid = deviceUtdid
    r.Set("device_utdid", deviceUtdid)
    return nil
}

func (r AlibabaAlihealthDrugcodeScanRequest) GetDeviceUtdid() string {
    return r.deviceUtdid
}

func (r *AlibabaAlihealthDrugcodeScanRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaAlihealthDrugcodeScanRequest) GetUserId() string {
    return r.userId
}


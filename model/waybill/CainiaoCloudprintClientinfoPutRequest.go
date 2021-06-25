package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
云打印客户端监控信息收集 APIRequest
cainiao.cloudprint.clientinfo.put

云打印客户端监控信息收集
*/
type CainiaoCloudprintClientinfoPutRequest struct {
    model.Params

    // 客户端上传json数据
    jsonData   string 

}

func NewCainiaoCloudprintClientinfoPutRequest() *CainiaoCloudprintClientinfoPutRequest{
    return &CainiaoCloudprintClientinfoPutRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoCloudprintClientinfoPutRequest) GetApiMethodName() string {
    return "cainiao.cloudprint.clientinfo.put"
}

func (r CainiaoCloudprintClientinfoPutRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoCloudprintClientinfoPutRequest) SetJsonData(jsonData string) error {
    r.jsonData = jsonData
    r.Set("json_data", jsonData)
    return nil
}

func (r CainiaoCloudprintClientinfoPutRequest) GetJsonData() string {
    return r.jsonData
}


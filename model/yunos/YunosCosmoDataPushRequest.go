package yunos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
COSMO-PUSH模式数据接入 APIRequest
yunos.cosmo.data.push

YunOS提供外部数据源接入，并输出到多端设备上，该接口提供了PUSH模式的数据接入
*/
type YunosCosmoDataPushRequest struct {
    model.Params

    // 业务方数据源唯一标识，由COSMO平台颁发
    appId   string 

    // 业务方推送数据，List结构的JSON序列化字符串
    jsonModel   string 

}

func NewYunosCosmoDataPushRequest() *YunosCosmoDataPushRequest{
    return &YunosCosmoDataPushRequest{
        Params: model.NewParams(),
    }
}

func (r YunosCosmoDataPushRequest) GetApiMethodName() string {
    return "yunos.cosmo.data.push"
}

func (r YunosCosmoDataPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosCosmoDataPushRequest) SetAppId(appId string) error {
    r.appId = appId
    r.Set("app_id", appId)
    return nil
}

func (r YunosCosmoDataPushRequest) GetAppId() string {
    return r.appId
}

func (r *YunosCosmoDataPushRequest) SetJsonModel(jsonModel string) error {
    r.jsonModel = jsonModel
    r.Set("json_model", jsonModel)
    return nil
}

func (r YunosCosmoDataPushRequest) GetJsonModel() string {
    return r.jsonModel
}


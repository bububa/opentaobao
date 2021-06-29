package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建应用升级任务 APIRequest
yunos.osupdate.appversion.create

创建应用升级任务
*/
type YunosOsupdateAppversionCreateRequest struct {
    model.Params

    // 应用版本信息
    appVersion   *TvAppVersion 

}

func NewYunosOsupdateAppversionCreateRequest() *YunosOsupdateAppversionCreateRequest{
    return &YunosOsupdateAppversionCreateRequest{
        Params: model.NewParams(),
    }
}

func (r YunosOsupdateAppversionCreateRequest) GetApiMethodName() string {
    return "yunos.osupdate.appversion.create"
}

func (r YunosOsupdateAppversionCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosOsupdateAppversionCreateRequest) SetAppVersion(appVersion *TvAppVersion) error {
    r.appVersion = appVersion
    r.Set("app_version", appVersion)
    return nil
}

func (r YunosOsupdateAppversionCreateRequest) GetAppVersion() *TvAppVersion {
    return r.appVersion
}


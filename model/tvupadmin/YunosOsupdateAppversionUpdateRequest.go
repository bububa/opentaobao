package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
应用升级任务更新 APIRequest
yunos.osupdate.appversion.update

应用升级任务更新
*/
type YunosOsupdateAppversionUpdateRequest struct {
    model.Params

    // 应用版本升级信息
    appVersion   *TvAppVersion 

}

func NewYunosOsupdateAppversionUpdateRequest() *YunosOsupdateAppversionUpdateRequest{
    return &YunosOsupdateAppversionUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r YunosOsupdateAppversionUpdateRequest) GetApiMethodName() string {
    return "yunos.osupdate.appversion.update"
}

func (r YunosOsupdateAppversionUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosOsupdateAppversionUpdateRequest) SetAppVersion(appVersion *TvAppVersion) error {
    r.appVersion = appVersion
    r.Set("app_version", appVersion)
    return nil
}

func (r YunosOsupdateAppversionUpdateRequest) GetAppVersion() *TvAppVersion {
    return r.appVersion
}


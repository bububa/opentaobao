package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
应用升级任务更新 API请求
yunos.osupdate.appversion.update

应用升级任务更新
*/
type YunosOsupdateAppversionUpdateRequest struct {
    model.Params
    // 应用版本升级信息
    appVersion   *TvAppVersion
}

// 初始化YunosOsupdateAppversionUpdateRequest对象
func NewYunosOsupdateAppversionUpdateRequest() *YunosOsupdateAppversionUpdateRequest{
    return &YunosOsupdateAppversionUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosOsupdateAppversionUpdateRequest) GetApiMethodName() string {
    return "yunos.osupdate.appversion.update"
}

// IRequest interface 方法, 获取API参数
func (r YunosOsupdateAppversionUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppVersion Setter
// 应用版本升级信息
func (r *YunosOsupdateAppversionUpdateRequest) SetAppVersion(appVersion *TvAppVersion) error {
    r.appVersion = appVersion
    r.Set("app_version", appVersion)
    return nil
}

// AppVersion Getter
func (r YunosOsupdateAppversionUpdateRequest) GetAppVersion() *TvAppVersion {
    return r.appVersion
}

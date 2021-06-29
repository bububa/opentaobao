package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建应用升级任务 API请求
yunos.osupdate.appversion.create

创建应用升级任务
*/
type YunosOsupdateAppversionCreateRequest struct {
    model.Params
    // 应用版本信息
    appVersion   *TvAppVersion
}

// 初始化YunosOsupdateAppversionCreateRequest对象
func NewYunosOsupdateAppversionCreateRequest() *YunosOsupdateAppversionCreateRequest{
    return &YunosOsupdateAppversionCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosOsupdateAppversionCreateRequest) GetApiMethodName() string {
    return "yunos.osupdate.appversion.create"
}

// IRequest interface 方法, 获取API参数
func (r YunosOsupdateAppversionCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppVersion Setter
// 应用版本信息
func (r *YunosOsupdateAppversionCreateRequest) SetAppVersion(appVersion *TvAppVersion) error {
    r.appVersion = appVersion
    r.Set("app_version", appVersion)
    return nil
}

// AppVersion Getter
func (r YunosOsupdateAppversionCreateRequest) GetAppVersion() *TvAppVersion {
    return r.appVersion
}

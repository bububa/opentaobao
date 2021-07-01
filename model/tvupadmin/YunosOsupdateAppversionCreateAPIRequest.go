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
type YunosOsupdateAppversionCreateAPIRequest struct {
    model.Params
    // 应用版本信息
    _appVersion   *TvAppVersion
}

// 初始化YunosOsupdateAppversionCreateAPIRequest对象
func NewYunosOsupdateAppversionCreateRequest() *YunosOsupdateAppversionCreateAPIRequest{
    return &YunosOsupdateAppversionCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosOsupdateAppversionCreateAPIRequest) GetApiMethodName() string {
    return "yunos.osupdate.appversion.create"
}

// IRequest interface 方法, 获取API参数
func (r YunosOsupdateAppversionCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppVersion Setter
// 应用版本信息
func (r *YunosOsupdateAppversionCreateAPIRequest) SetAppVersion(_appVersion *TvAppVersion) error {
    r._appVersion = _appVersion
    r.Set("app_version", _appVersion)
    return nil
}

// AppVersion Getter
func (r YunosOsupdateAppversionCreateAPIRequest) GetAppVersion() *TvAppVersion {
    return r._appVersion
}

package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加系统升级任务 API请求
yunos.osupdate.osfota.add

添加osupdate系统升级任务
*/
type YunosOsupdateOsfotaAddRequest struct {
    model.Params
    // 系统升级任务json格式
    osFotaJson   string
}

// 初始化YunosOsupdateOsfotaAddRequest对象
func NewYunosOsupdateOsfotaAddRequest() *YunosOsupdateOsfotaAddRequest{
    return &YunosOsupdateOsfotaAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosOsupdateOsfotaAddRequest) GetApiMethodName() string {
    return "yunos.osupdate.osfota.add"
}

// IRequest interface 方法, 获取API参数
func (r YunosOsupdateOsfotaAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OsFotaJson Setter
// 系统升级任务json格式
func (r *YunosOsupdateOsfotaAddRequest) SetOsFotaJson(osFotaJson string) error {
    r.osFotaJson = osFotaJson
    r.Set("os_fota_json", osFotaJson)
    return nil
}

// OsFotaJson Getter
func (r YunosOsupdateOsfotaAddRequest) GetOsFotaJson() string {
    return r.osFotaJson
}

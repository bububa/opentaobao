package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
系统升级发布 API请求
yunos.osupdate.osfota.publish

发布osupdate系统升级任务
*/
type YunosOsupdateOsfotaPublishRequest struct {
    model.Params
    // 入参json格式
    _publishJson   string
}

// 初始化YunosOsupdateOsfotaPublishRequest对象
func NewYunosOsupdateOsfotaPublishRequest() *YunosOsupdateOsfotaPublishRequest{
    return &YunosOsupdateOsfotaPublishRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosOsupdateOsfotaPublishRequest) GetApiMethodName() string {
    return "yunos.osupdate.osfota.publish"
}

// IRequest interface 方法, 获取API参数
func (r YunosOsupdateOsfotaPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PublishJson Setter
// 入参json格式
func (r *YunosOsupdateOsfotaPublishRequest) SetPublishJson(_publishJson string) error {
    r._publishJson = _publishJson
    r.Set("publish_json", _publishJson)
    return nil
}

// PublishJson Getter
func (r YunosOsupdateOsfotaPublishRequest) GetPublishJson() string {
    return r._publishJson
}

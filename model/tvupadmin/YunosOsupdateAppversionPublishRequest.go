package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发布应用升级 API请求
yunos.osupdate.appversion.publish

发布应用升级任务
*/
type YunosOsupdateAppversionPublishRequest struct {
    model.Params
    // 发布应用升级入参json
    publishJson   string
}

// 初始化YunosOsupdateAppversionPublishRequest对象
func NewYunosOsupdateAppversionPublishRequest() *YunosOsupdateAppversionPublishRequest{
    return &YunosOsupdateAppversionPublishRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosOsupdateAppversionPublishRequest) GetApiMethodName() string {
    return "yunos.osupdate.appversion.publish"
}

// IRequest interface 方法, 获取API参数
func (r YunosOsupdateAppversionPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PublishJson Setter
// 发布应用升级入参json
func (r *YunosOsupdateAppversionPublishRequest) SetPublishJson(publishJson string) error {
    r.publishJson = publishJson
    r.Set("publish_json", publishJson)
    return nil
}

// PublishJson Getter
func (r YunosOsupdateAppversionPublishRequest) GetPublishJson() string {
    return r.publishJson
}

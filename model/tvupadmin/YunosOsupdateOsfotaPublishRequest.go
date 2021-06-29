package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
系统升级发布 APIRequest
yunos.osupdate.osfota.publish

发布osupdate系统升级任务
*/
type YunosOsupdateOsfotaPublishRequest struct {
    model.Params

    // 入参json格式
    publishJson   string 

}

func NewYunosOsupdateOsfotaPublishRequest() *YunosOsupdateOsfotaPublishRequest{
    return &YunosOsupdateOsfotaPublishRequest{
        Params: model.NewParams(),
    }
}

func (r YunosOsupdateOsfotaPublishRequest) GetApiMethodName() string {
    return "yunos.osupdate.osfota.publish"
}

func (r YunosOsupdateOsfotaPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosOsupdateOsfotaPublishRequest) SetPublishJson(publishJson string) error {
    r.publishJson = publishJson
    r.Set("publish_json", publishJson)
    return nil
}

func (r YunosOsupdateOsfotaPublishRequest) GetPublishJson() string {
    return r.publishJson
}


package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发布应用升级 APIRequest
yunos.osupdate.appversion.publish

发布应用升级任务
*/
type YunosOsupdateAppversionPublishRequest struct {
    model.Params

    // 发布应用升级入参json
    publishJson   string 

}

func NewYunosOsupdateAppversionPublishRequest() *YunosOsupdateAppversionPublishRequest{
    return &YunosOsupdateAppversionPublishRequest{
        Params: model.NewParams(),
    }
}

func (r YunosOsupdateAppversionPublishRequest) GetApiMethodName() string {
    return "yunos.osupdate.appversion.publish"
}

func (r YunosOsupdateAppversionPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosOsupdateAppversionPublishRequest) SetPublishJson(publishJson string) error {
    r.publishJson = publishJson
    r.Set("publish_json", publishJson)
    return nil
}

func (r YunosOsupdateAppversionPublishRequest) GetPublishJson() string {
    return r.publishJson
}


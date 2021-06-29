package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
机型检索 APIRequest
yunos.osupdate.model.search

机型检索
*/
type YunosOsupdateModelSearchRequest struct {
    model.Params

    // 应用ID
    appId   int64 

    // 关键词
    name   string 

}

func NewYunosOsupdateModelSearchRequest() *YunosOsupdateModelSearchRequest{
    return &YunosOsupdateModelSearchRequest{
        Params: model.NewParams(),
    }
}

func (r YunosOsupdateModelSearchRequest) GetApiMethodName() string {
    return "yunos.osupdate.model.search"
}

func (r YunosOsupdateModelSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosOsupdateModelSearchRequest) SetAppId(appId int64) error {
    r.appId = appId
    r.Set("app_id", appId)
    return nil
}

func (r YunosOsupdateModelSearchRequest) GetAppId() int64 {
    return r.appId
}

func (r *YunosOsupdateModelSearchRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r YunosOsupdateModelSearchRequest) GetName() string {
    return r.name
}


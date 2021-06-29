package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取应用升级详情 APIRequest
yunos.osupdate.appversion.info

获取应用升级详情
*/
type YunosOsupdateAppversionInfoRequest struct {
    model.Params

    // 升级任务ID
    id   int64 

}

func NewYunosOsupdateAppversionInfoRequest() *YunosOsupdateAppversionInfoRequest{
    return &YunosOsupdateAppversionInfoRequest{
        Params: model.NewParams(),
    }
}

func (r YunosOsupdateAppversionInfoRequest) GetApiMethodName() string {
    return "yunos.osupdate.appversion.info"
}

func (r YunosOsupdateAppversionInfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosOsupdateAppversionInfoRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r YunosOsupdateAppversionInfoRequest) GetId() int64 {
    return r.id
}


package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新应用升级状态 APIRequest
yunos.osupdate.versionstatus.update

更新应用升级状态
*/
type YunosOsupdateVersionstatusUpdateRequest struct {
    model.Params

    // 升级任务ID
    id   int64 

    // 状态值
    status   string 

}

func NewYunosOsupdateVersionstatusUpdateRequest() *YunosOsupdateVersionstatusUpdateRequest{
    return &YunosOsupdateVersionstatusUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r YunosOsupdateVersionstatusUpdateRequest) GetApiMethodName() string {
    return "yunos.osupdate.versionstatus.update"
}

func (r YunosOsupdateVersionstatusUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosOsupdateVersionstatusUpdateRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r YunosOsupdateVersionstatusUpdateRequest) GetId() int64 {
    return r.id
}

func (r *YunosOsupdateVersionstatusUpdateRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r YunosOsupdateVersionstatusUpdateRequest) GetStatus() string {
    return r.status
}


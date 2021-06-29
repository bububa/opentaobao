package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新应用升级状态 API请求
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

// 初始化YunosOsupdateVersionstatusUpdateRequest对象
func NewYunosOsupdateVersionstatusUpdateRequest() *YunosOsupdateVersionstatusUpdateRequest{
    return &YunosOsupdateVersionstatusUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosOsupdateVersionstatusUpdateRequest) GetApiMethodName() string {
    return "yunos.osupdate.versionstatus.update"
}

// IRequest interface 方法, 获取API参数
func (r YunosOsupdateVersionstatusUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 升级任务ID
func (r *YunosOsupdateVersionstatusUpdateRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r YunosOsupdateVersionstatusUpdateRequest) GetId() int64 {
    return r.id
}
// Status Setter
// 状态值
func (r *YunosOsupdateVersionstatusUpdateRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r YunosOsupdateVersionstatusUpdateRequest) GetStatus() string {
    return r.status
}

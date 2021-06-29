package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
切换用户 APIRequest
alibaba.ailabs.tmallgenie.auth.switchuser

设备切换授权用户
*/
type AlibabaAilabsTmallgenieAuthSwitchuserRequest struct {
    model.Params

    // client_id
    clientId   string 

    // 目标用户openId
    newUserOpenId   string 

    // 当前拥有设备权限的用户openId
    oldUserOpenId   string 

    // 设备uuid
    uuid   string 

}

func NewAlibabaAilabsTmallgenieAuthSwitchuserRequest() *AlibabaAilabsTmallgenieAuthSwitchuserRequest{
    return &AlibabaAilabsTmallgenieAuthSwitchuserRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsTmallgenieAuthSwitchuserRequest) GetApiMethodName() string {
    return "alibaba.ailabs.tmallgenie.auth.switchuser"
}

func (r AlibabaAilabsTmallgenieAuthSwitchuserRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsTmallgenieAuthSwitchuserRequest) SetClientId(clientId string) error {
    r.clientId = clientId
    r.Set("client_id", clientId)
    return nil
}

func (r AlibabaAilabsTmallgenieAuthSwitchuserRequest) GetClientId() string {
    return r.clientId
}

func (r *AlibabaAilabsTmallgenieAuthSwitchuserRequest) SetNewUserOpenId(newUserOpenId string) error {
    r.newUserOpenId = newUserOpenId
    r.Set("new_user_open_id", newUserOpenId)
    return nil
}

func (r AlibabaAilabsTmallgenieAuthSwitchuserRequest) GetNewUserOpenId() string {
    return r.newUserOpenId
}

func (r *AlibabaAilabsTmallgenieAuthSwitchuserRequest) SetOldUserOpenId(oldUserOpenId string) error {
    r.oldUserOpenId = oldUserOpenId
    r.Set("old_user_open_id", oldUserOpenId)
    return nil
}

func (r AlibabaAilabsTmallgenieAuthSwitchuserRequest) GetOldUserOpenId() string {
    return r.oldUserOpenId
}

func (r *AlibabaAilabsTmallgenieAuthSwitchuserRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r AlibabaAilabsTmallgenieAuthSwitchuserRequest) GetUuid() string {
    return r.uuid
}


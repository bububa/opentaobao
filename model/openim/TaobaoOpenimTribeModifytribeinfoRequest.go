package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群信息修改 APIRequest
taobao.openim.tribe.modifytribeinfo

OPENIM群信息修改
*/
type TaobaoOpenimTribeModifytribeinfoRequest struct {
    model.Params

    // 用户信息
    user   *OpenImUser 

    // 群名称
    tribeName   string 

    // 群公告
    notice   string 

    // 群id
    tribeId   int64 

}

func NewTaobaoOpenimTribeModifytribeinfoRequest() *TaobaoOpenimTribeModifytribeinfoRequest{
    return &TaobaoOpenimTribeModifytribeinfoRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenimTribeModifytribeinfoRequest) GetApiMethodName() string {
    return "taobao.openim.tribe.modifytribeinfo"
}

func (r TaobaoOpenimTribeModifytribeinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenimTribeModifytribeinfoRequest) SetUser(user *OpenImUser) error {
    r.user = user
    r.Set("user", user)
    return nil
}

func (r TaobaoOpenimTribeModifytribeinfoRequest) GetUser() *OpenImUser {
    return r.user
}

func (r *TaobaoOpenimTribeModifytribeinfoRequest) SetTribeName(tribeName string) error {
    r.tribeName = tribeName
    r.Set("tribe_name", tribeName)
    return nil
}

func (r TaobaoOpenimTribeModifytribeinfoRequest) GetTribeName() string {
    return r.tribeName
}

func (r *TaobaoOpenimTribeModifytribeinfoRequest) SetNotice(notice string) error {
    r.notice = notice
    r.Set("notice", notice)
    return nil
}

func (r TaobaoOpenimTribeModifytribeinfoRequest) GetNotice() string {
    return r.notice
}

func (r *TaobaoOpenimTribeModifytribeinfoRequest) SetTribeId(tribeId int64) error {
    r.tribeId = tribeId
    r.Set("tribe_id", tribeId)
    return nil
}

func (r TaobaoOpenimTribeModifytribeinfoRequest) GetTribeId() int64 {
    return r.tribeId
}


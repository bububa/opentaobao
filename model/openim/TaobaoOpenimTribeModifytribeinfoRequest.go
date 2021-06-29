package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群信息修改 API请求
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

// 初始化TaobaoOpenimTribeModifytribeinfoRequest对象
func NewTaobaoOpenimTribeModifytribeinfoRequest() *TaobaoOpenimTribeModifytribeinfoRequest{
    return &TaobaoOpenimTribeModifytribeinfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeModifytribeinfoRequest) GetApiMethodName() string {
    return "taobao.openim.tribe.modifytribeinfo"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeModifytribeinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// User Setter
// 用户信息
func (r *TaobaoOpenimTribeModifytribeinfoRequest) SetUser(user *OpenImUser) error {
    r.user = user
    r.Set("user", user)
    return nil
}

// User Getter
func (r TaobaoOpenimTribeModifytribeinfoRequest) GetUser() *OpenImUser {
    return r.user
}
// TribeName Setter
// 群名称
func (r *TaobaoOpenimTribeModifytribeinfoRequest) SetTribeName(tribeName string) error {
    r.tribeName = tribeName
    r.Set("tribe_name", tribeName)
    return nil
}

// TribeName Getter
func (r TaobaoOpenimTribeModifytribeinfoRequest) GetTribeName() string {
    return r.tribeName
}
// Notice Setter
// 群公告
func (r *TaobaoOpenimTribeModifytribeinfoRequest) SetNotice(notice string) error {
    r.notice = notice
    r.Set("notice", notice)
    return nil
}

// Notice Getter
func (r TaobaoOpenimTribeModifytribeinfoRequest) GetNotice() string {
    return r.notice
}
// TribeId Setter
// 群id
func (r *TaobaoOpenimTribeModifytribeinfoRequest) SetTribeId(tribeId int64) error {
    r.tribeId = tribeId
    r.Set("tribe_id", tribeId)
    return nil
}

// TribeId Getter
func (r TaobaoOpenimTribeModifytribeinfoRequest) GetTribeId() int64 {
    return r.tribeId
}

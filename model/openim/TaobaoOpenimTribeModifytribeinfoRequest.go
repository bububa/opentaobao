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
    _user   *OpenImUser
    // 群名称
    _tribeName   string
    // 群公告
    _notice   string
    // 群id
    _tribeId   int64
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
func (r *TaobaoOpenimTribeModifytribeinfoRequest) SetUser(_user *OpenImUser) error {
    r._user = _user
    r.Set("user", _user)
    return nil
}

// User Getter
func (r TaobaoOpenimTribeModifytribeinfoRequest) GetUser() *OpenImUser {
    return r._user
}
// TribeName Setter
// 群名称
func (r *TaobaoOpenimTribeModifytribeinfoRequest) SetTribeName(_tribeName string) error {
    r._tribeName = _tribeName
    r.Set("tribe_name", _tribeName)
    return nil
}

// TribeName Getter
func (r TaobaoOpenimTribeModifytribeinfoRequest) GetTribeName() string {
    return r._tribeName
}
// Notice Setter
// 群公告
func (r *TaobaoOpenimTribeModifytribeinfoRequest) SetNotice(_notice string) error {
    r._notice = _notice
    r.Set("notice", _notice)
    return nil
}

// Notice Getter
func (r TaobaoOpenimTribeModifytribeinfoRequest) GetNotice() string {
    return r._notice
}
// TribeId Setter
// 群id
func (r *TaobaoOpenimTribeModifytribeinfoRequest) SetTribeId(_tribeId int64) error {
    r._tribeId = _tribeId
    r.Set("tribe_id", _tribeId)
    return nil
}

// TribeId Getter
func (r TaobaoOpenimTribeModifytribeinfoRequest) GetTribeId() int64 {
    return r._tribeId
}

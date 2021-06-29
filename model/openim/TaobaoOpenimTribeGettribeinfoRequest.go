package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取群信息 API请求
taobao.openim.tribe.gettribeinfo

获取群信息
*/
type TaobaoOpenimTribeGettribeinfoRequest struct {
    model.Params
    // 用户信息
    _user   *OpenImUser
    // 群ID
    _tribeId   int64
}

// 初始化TaobaoOpenimTribeGettribeinfoRequest对象
func NewTaobaoOpenimTribeGettribeinfoRequest() *TaobaoOpenimTribeGettribeinfoRequest{
    return &TaobaoOpenimTribeGettribeinfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeGettribeinfoRequest) GetApiMethodName() string {
    return "taobao.openim.tribe.gettribeinfo"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeGettribeinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// User Setter
// 用户信息
func (r *TaobaoOpenimTribeGettribeinfoRequest) SetUser(_user *OpenImUser) error {
    r._user = _user
    r.Set("user", _user)
    return nil
}

// User Getter
func (r TaobaoOpenimTribeGettribeinfoRequest) GetUser() *OpenImUser {
    return r._user
}
// TribeId Setter
// 群ID
func (r *TaobaoOpenimTribeGettribeinfoRequest) SetTribeId(_tribeId int64) error {
    r._tribeId = _tribeId
    r.Set("tribe_id", _tribeId)
    return nil
}

// TribeId Getter
func (r TaobaoOpenimTribeGettribeinfoRequest) GetTribeId() int64 {
    return r._tribeId
}

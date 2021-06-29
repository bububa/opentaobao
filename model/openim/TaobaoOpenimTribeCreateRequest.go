package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建群 API请求
taobao.openim.tribe.create

创建一个openim的群
*/
type TaobaoOpenimTribeCreateRequest struct {
    model.Params
    // 用户信息
    user   *OpenImUser
    // 群名称
    tribeName   string
    // 群公告
    notice   string
    // 群类型有两种tribe_type = 0  普通群  普通群有管理员角色，对成员加入有权限控制tribe_type = 1  讨论组  讨论组没有管理员，不能解散
    tribeType   int64
    // 创建群时候拉入群的成员tribe_type = 1（即为讨论组类型)时  该参数为必选tribe_type = 0 (即为普通群类型)时，改参数无效，可不填
    members   []OpenImUser
}

// 初始化TaobaoOpenimTribeCreateRequest对象
func NewTaobaoOpenimTribeCreateRequest() *TaobaoOpenimTribeCreateRequest{
    return &TaobaoOpenimTribeCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeCreateRequest) GetApiMethodName() string {
    return "taobao.openim.tribe.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// User Setter
// 用户信息
func (r *TaobaoOpenimTribeCreateRequest) SetUser(user *OpenImUser) error {
    r.user = user
    r.Set("user", user)
    return nil
}

// User Getter
func (r TaobaoOpenimTribeCreateRequest) GetUser() *OpenImUser {
    return r.user
}
// TribeName Setter
// 群名称
func (r *TaobaoOpenimTribeCreateRequest) SetTribeName(tribeName string) error {
    r.tribeName = tribeName
    r.Set("tribe_name", tribeName)
    return nil
}

// TribeName Getter
func (r TaobaoOpenimTribeCreateRequest) GetTribeName() string {
    return r.tribeName
}
// Notice Setter
// 群公告
func (r *TaobaoOpenimTribeCreateRequest) SetNotice(notice string) error {
    r.notice = notice
    r.Set("notice", notice)
    return nil
}

// Notice Getter
func (r TaobaoOpenimTribeCreateRequest) GetNotice() string {
    return r.notice
}
// TribeType Setter
// 群类型有两种tribe_type = 0  普通群  普通群有管理员角色，对成员加入有权限控制tribe_type = 1  讨论组  讨论组没有管理员，不能解散
func (r *TaobaoOpenimTribeCreateRequest) SetTribeType(tribeType int64) error {
    r.tribeType = tribeType
    r.Set("tribe_type", tribeType)
    return nil
}

// TribeType Getter
func (r TaobaoOpenimTribeCreateRequest) GetTribeType() int64 {
    return r.tribeType
}
// Members Setter
// 创建群时候拉入群的成员tribe_type = 1（即为讨论组类型)时  该参数为必选tribe_type = 0 (即为普通群类型)时，改参数无效，可不填
func (r *TaobaoOpenimTribeCreateRequest) SetMembers(members []OpenImUser) error {
    r.members = members
    r.Set("members", members)
    return nil
}

// Members Getter
func (r TaobaoOpenimTribeCreateRequest) GetMembers() []OpenImUser {
    return r.members
}

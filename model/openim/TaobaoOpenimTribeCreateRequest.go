package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建群 APIRequest
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

func NewTaobaoOpenimTribeCreateRequest() *TaobaoOpenimTribeCreateRequest{
    return &TaobaoOpenimTribeCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenimTribeCreateRequest) GetApiMethodName() string {
    return "taobao.openim.tribe.create"
}

func (r TaobaoOpenimTribeCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenimTribeCreateRequest) SetUser(user *OpenImUser) error {
    r.user = user
    r.Set("user", user)
    return nil
}

func (r TaobaoOpenimTribeCreateRequest) GetUser() *OpenImUser {
    return r.user
}

func (r *TaobaoOpenimTribeCreateRequest) SetTribeName(tribeName string) error {
    r.tribeName = tribeName
    r.Set("tribe_name", tribeName)
    return nil
}

func (r TaobaoOpenimTribeCreateRequest) GetTribeName() string {
    return r.tribeName
}

func (r *TaobaoOpenimTribeCreateRequest) SetNotice(notice string) error {
    r.notice = notice
    r.Set("notice", notice)
    return nil
}

func (r TaobaoOpenimTribeCreateRequest) GetNotice() string {
    return r.notice
}

func (r *TaobaoOpenimTribeCreateRequest) SetTribeType(tribeType int64) error {
    r.tribeType = tribeType
    r.Set("tribe_type", tribeType)
    return nil
}

func (r TaobaoOpenimTribeCreateRequest) GetTribeType() int64 {
    return r.tribeType
}

func (r *TaobaoOpenimTribeCreateRequest) SetMembers(members []OpenImUser) error {
    r.members = members
    r.Set("members", members)
    return nil
}

func (r TaobaoOpenimTribeCreateRequest) GetMembers() []OpenImUser {
    return r.members
}


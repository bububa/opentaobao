package subuser

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取指定账户子账号的详细信息 APIRequest
taobao.subuser.fullinfo.get

获取指定账户子账号的详细信息，其中包括子账号的账号信息以及员工、部门、职务信息（只能通过主账号登陆并查询属于该主账号下的某个子账号详细信息）
*/
type TaobaoSubuserFullinfoGetRequest struct {
    model.Params

    // 子账号ID（传参中sub_id和sub_nick至少需要其中一个，若sub_id与sub_nick同时传入并且合法，那么sub_nick优先，以sub_nick查询子账号）
    subId   int64 

    // 传入所需要的参数信息（若不需要获取子账号或主账号的企业邮箱地址，则无需传入该参数；若需要获取子账号或主账号的企业邮箱地址，则需要传入fields；可选参数值为subuser_email和user_email，传入其他参数值均无效；两个参数都需要则以逗号隔开传入即可，例如：subuser_email,user_email）
    fields   string 

    // 子账号用户名（传参中sub_id和sub_nick至少需要其中一个，若sub_id与sub_nick同时传入并且合法，那么sub_nick优先，以sub_nick查询子账号）
    subNick   string 

}

func NewTaobaoSubuserFullinfoGetRequest() *TaobaoSubuserFullinfoGetRequest{
    return &TaobaoSubuserFullinfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSubuserFullinfoGetRequest) GetApiMethodName() string {
    return "taobao.subuser.fullinfo.get"
}

func (r TaobaoSubuserFullinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSubuserFullinfoGetRequest) SetSubId(subId int64) error {
    r.subId = subId
    r.Set("sub_id", subId)
    return nil
}

func (r TaobaoSubuserFullinfoGetRequest) GetSubId() int64 {
    return r.subId
}

func (r *TaobaoSubuserFullinfoGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoSubuserFullinfoGetRequest) GetFields() string {
    return r.fields
}

func (r *TaobaoSubuserFullinfoGetRequest) SetSubNick(subNick string) error {
    r.subNick = subNick
    r.Set("sub_nick", subNick)
    return nil
}

func (r TaobaoSubuserFullinfoGetRequest) GetSubNick() string {
    return r.subNick
}


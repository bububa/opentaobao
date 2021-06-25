package subuser

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改指定账户子账号的基本信息 APIRequest
taobao.subuser.info.update

修改指定账户子账号的基本信息（通过主账号登陆只能修改属于该主账号的子账号基本信息）
*/
type TaobaoSubuserInfoUpdateRequest struct {
    model.Params

    // 是否停用子账号 true:表示停用该子账号false:表示开启该子账号
    isDisableSubaccount   bool 

    // 子账号是否参与分流 true:参与分流 false:不参与分流
    isDispatch   bool 

    // 子账号ID
    subId   int64 

}

func NewTaobaoSubuserInfoUpdateRequest() *TaobaoSubuserInfoUpdateRequest{
    return &TaobaoSubuserInfoUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSubuserInfoUpdateRequest) GetApiMethodName() string {
    return "taobao.subuser.info.update"
}

func (r TaobaoSubuserInfoUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSubuserInfoUpdateRequest) SetIsDisableSubaccount(isDisableSubaccount bool) error {
    r.isDisableSubaccount = isDisableSubaccount
    r.Set("is_disable_subaccount", isDisableSubaccount)
    return nil
}

func (r TaobaoSubuserInfoUpdateRequest) GetIsDisableSubaccount() bool {
    return r.isDisableSubaccount
}

func (r *TaobaoSubuserInfoUpdateRequest) SetIsDispatch(isDispatch bool) error {
    r.isDispatch = isDispatch
    r.Set("is_dispatch", isDispatch)
    return nil
}

func (r TaobaoSubuserInfoUpdateRequest) GetIsDispatch() bool {
    return r.isDispatch
}

func (r *TaobaoSubuserInfoUpdateRequest) SetSubId(subId int64) error {
    r.subId = subId
    r.Set("sub_id", subId)
    return nil
}

func (r TaobaoSubuserInfoUpdateRequest) GetSubId() int64 {
    return r.subId
}


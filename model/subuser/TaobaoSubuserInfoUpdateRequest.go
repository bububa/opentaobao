package subuser

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改指定账户子账号的基本信息 API请求
taobao.subuser.info.update

修改指定账户子账号的基本信息（通过主账号登陆只能修改属于该主账号的子账号基本信息）
*/
type TaobaoSubuserInfoUpdateRequest struct {
    model.Params
    // 是否停用子账号 true:表示停用该子账号false:表示开启该子账号
    _isDisableSubaccount   bool
    // 子账号是否参与分流 true:参与分流 false:不参与分流
    _isDispatch   bool
    // 子账号ID
    _subId   int64
}

// 初始化TaobaoSubuserInfoUpdateRequest对象
func NewTaobaoSubuserInfoUpdateRequest() *TaobaoSubuserInfoUpdateRequest{
    return &TaobaoSubuserInfoUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSubuserInfoUpdateRequest) GetApiMethodName() string {
    return "taobao.subuser.info.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSubuserInfoUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IsDisableSubaccount Setter
// 是否停用子账号 true:表示停用该子账号false:表示开启该子账号
func (r *TaobaoSubuserInfoUpdateRequest) SetIsDisableSubaccount(_isDisableSubaccount bool) error {
    r._isDisableSubaccount = _isDisableSubaccount
    r.Set("is_disable_subaccount", _isDisableSubaccount)
    return nil
}

// IsDisableSubaccount Getter
func (r TaobaoSubuserInfoUpdateRequest) GetIsDisableSubaccount() bool {
    return r._isDisableSubaccount
}
// IsDispatch Setter
// 子账号是否参与分流 true:参与分流 false:不参与分流
func (r *TaobaoSubuserInfoUpdateRequest) SetIsDispatch(_isDispatch bool) error {
    r._isDispatch = _isDispatch
    r.Set("is_dispatch", _isDispatch)
    return nil
}

// IsDispatch Getter
func (r TaobaoSubuserInfoUpdateRequest) GetIsDispatch() bool {
    return r._isDispatch
}
// SubId Setter
// 子账号ID
func (r *TaobaoSubuserInfoUpdateRequest) SetSubId(_subId int64) error {
    r._subId = _subId
    r.Set("sub_id", _subId)
    return nil
}

// SubId Getter
func (r TaobaoSubuserInfoUpdateRequest) GetSubId() int64 {
    return r._subId
}

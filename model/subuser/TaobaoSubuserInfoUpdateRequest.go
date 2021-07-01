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
type TaobaoSubuserInfoUpdateAPIRequest struct {
    model.Params
    // 是否停用子账号 true:表示停用该子账号false:表示开启该子账号
    _isDisableSubaccount   bool
    // 子账号是否参与分流 true:参与分流 false:不参与分流
    _isDispatch   bool
    // 子账号ID
    _subId   int64
}

// 初始化TaobaoSubuserInfoUpdateAPIRequest对象
func NewTaobaoSubuserInfoUpdateRequest() *TaobaoSubuserInfoUpdateAPIRequest{
    return &TaobaoSubuserInfoUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSubuserInfoUpdateAPIRequest) GetApiMethodName() string {
    return "taobao.subuser.info.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSubuserInfoUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IsDisableSubaccount Setter
// 是否停用子账号 true:表示停用该子账号false:表示开启该子账号
func (r *TaobaoSubuserInfoUpdateAPIRequest) SetIsDisableSubaccount(_isDisableSubaccount bool) error {
    r._isDisableSubaccount = _isDisableSubaccount
    r.Set("is_disable_subaccount", _isDisableSubaccount)
    return nil
}

// IsDisableSubaccount Getter
func (r TaobaoSubuserInfoUpdateAPIRequest) GetIsDisableSubaccount() bool {
    return r._isDisableSubaccount
}
// IsDispatch Setter
// 子账号是否参与分流 true:参与分流 false:不参与分流
func (r *TaobaoSubuserInfoUpdateAPIRequest) SetIsDispatch(_isDispatch bool) error {
    r._isDispatch = _isDispatch
    r.Set("is_dispatch", _isDispatch)
    return nil
}

// IsDispatch Getter
func (r TaobaoSubuserInfoUpdateAPIRequest) GetIsDispatch() bool {
    return r._isDispatch
}
// SubId Setter
// 子账号ID
func (r *TaobaoSubuserInfoUpdateAPIRequest) SetSubId(_subId int64) error {
    r._subId = _subId
    r.Set("sub_id", _subId)
    return nil
}

// SubId Getter
func (r TaobaoSubuserInfoUpdateAPIRequest) GetSubId() int64 {
    return r._subId
}

package subuser

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosubuserinfoupdateAPIRequest 修改指定账户子账号的基本信息 API请求
// taobao.subuser.info.update
//
// 修改指定账户子账号的基本信息（通过主账号登陆只能修改属于该主账号的子账号基本信息）
type TaobaosubuserinfoupdateAPIRequest struct {
	model.Params
	// 子账号ID
	_subId int64
	// 子账号是否参与分流 true:参与分流 false:不参与分流
	_isDispatch bool
	// 是否停用子账号 true:表示停用该子账号false:表示开启该子账号
	_isDisableSubaccount bool
}

// NewTaobaosubuserinfoupdateRequest 初始化TaobaosubuserinfoupdateAPIRequest对象
func NewTaobaosubuserinfoupdateRequest() *TaobaosubuserinfoupdateAPIRequest {
	return &TaobaosubuserinfoupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosubuserinfoupdateAPIRequest) GetApiMethodName() string {
	return "taobao.subuser.info.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosubuserinfoupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosubuserinfoupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubId is SubId Setter
// 子账号ID
func (r *TaobaosubuserinfoupdateAPIRequest) SetSubId(_subId int64) error {
	r._subId = _subId
	r.Set("sub_id", _subId)
	return nil
}

// GetSubId SubId Getter
func (r TaobaosubuserinfoupdateAPIRequest) GetSubId() int64 {
	return r._subId
}

// SetIsDispatch is IsDispatch Setter
// 子账号是否参与分流 true:参与分流 false:不参与分流
func (r *TaobaosubuserinfoupdateAPIRequest) SetIsDispatch(_isDispatch bool) error {
	r._isDispatch = _isDispatch
	r.Set("is_dispatch", _isDispatch)
	return nil
}

// GetIsDispatch IsDispatch Getter
func (r TaobaosubuserinfoupdateAPIRequest) GetIsDispatch() bool {
	return r._isDispatch
}

// SetIsDisableSubaccount is IsDisableSubaccount Setter
// 是否停用子账号 true:表示停用该子账号false:表示开启该子账号
func (r *TaobaosubuserinfoupdateAPIRequest) SetIsDisableSubaccount(_isDisableSubaccount bool) error {
	r._isDisableSubaccount = _isDisableSubaccount
	r.Set("is_disable_subaccount", _isDisableSubaccount)
	return nil
}

// GetIsDisableSubaccount IsDisableSubaccount Getter
func (r TaobaosubuserinfoupdateAPIRequest) GetIsDisableSubaccount() bool {
	return r._isDisableSubaccount
}

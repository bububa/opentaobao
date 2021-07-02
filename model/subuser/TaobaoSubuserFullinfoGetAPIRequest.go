package subuser

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubuserFullinfoGetAPIRequest 获取指定账户子账号的详细信息 API请求
// taobao.subuser.fullinfo.get
//
// 获取指定账户子账号的详细信息，其中包括子账号的账号信息以及员工、部门、职务信息（只能通过主账号登陆并查询属于该主账号下的某个子账号详细信息）
type TaobaoSubuserFullinfoGetAPIRequest struct {
	model.Params
	// 子账号ID（传参中sub_id和sub_nick至少需要其中一个，若sub_id与sub_nick同时传入并且合法，那么sub_nick优先，以sub_nick查询子账号）
	_subId int64
	// 传入所需要的参数信息（若不需要获取子账号或主账号的企业邮箱地址，则无需传入该参数；若需要获取子账号或主账号的企业邮箱地址，则需要传入fields；可选参数值为subuser_email和user_email，传入其他参数值均无效；两个参数都需要则以逗号隔开传入即可，例如：subuser_email,user_email）
	_fields string
	// 子账号用户名（传参中sub_id和sub_nick至少需要其中一个，若sub_id与sub_nick同时传入并且合法，那么sub_nick优先，以sub_nick查询子账号）
	_subNick string
}

// NewTaobaoSubuserFullinfoGetRequest 初始化TaobaoSubuserFullinfoGetAPIRequest对象
func NewTaobaoSubuserFullinfoGetRequest() *TaobaoSubuserFullinfoGetAPIRequest {
	return &TaobaoSubuserFullinfoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubuserFullinfoGetAPIRequest) GetApiMethodName() string {
	return "taobao.subuser.fullinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubuserFullinfoGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSubId is SubId Setter
// 子账号ID（传参中sub_id和sub_nick至少需要其中一个，若sub_id与sub_nick同时传入并且合法，那么sub_nick优先，以sub_nick查询子账号）
func (r *TaobaoSubuserFullinfoGetAPIRequest) SetSubId(_subId int64) error {
	r._subId = _subId
	r.Set("sub_id", _subId)
	return nil
}

// GetSubId SubId Getter
func (r TaobaoSubuserFullinfoGetAPIRequest) GetSubId() int64 {
	return r._subId
}

// SetFields is Fields Setter
// 传入所需要的参数信息（若不需要获取子账号或主账号的企业邮箱地址，则无需传入该参数；若需要获取子账号或主账号的企业邮箱地址，则需要传入fields；可选参数值为subuser_email和user_email，传入其他参数值均无效；两个参数都需要则以逗号隔开传入即可，例如：subuser_email,user_email）
func (r *TaobaoSubuserFullinfoGetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoSubuserFullinfoGetAPIRequest) GetFields() string {
	return r._fields
}

// SetSubNick is SubNick Setter
// 子账号用户名（传参中sub_id和sub_nick至少需要其中一个，若sub_id与sub_nick同时传入并且合法，那么sub_nick优先，以sub_nick查询子账号）
func (r *TaobaoSubuserFullinfoGetAPIRequest) SetSubNick(_subNick string) error {
	r._subNick = _subNick
	r.Set("sub_nick", _subNick)
	return nil
}

// GetSubNick SubNick Getter
func (r TaobaoSubuserFullinfoGetAPIRequest) GetSubNick() string {
	return r._subNick
}

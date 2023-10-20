package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkscucrowdmemberaddAPIRequest 淘宝客-服务商-上传人群数据 API请求
// taobao.tbk.sc.ucrowd.member.add
//
// 服务商使用。支持淘宝客上传人群标签id对应的会员列表，支持全量和增量多种数据更新方式。
type TaobaotbkscucrowdmemberaddAPIRequest struct {
	model.Params
	// 成员列表
	_accountlist []string
	// 成员账号类型，1:sid
	_accounttype int64
	// 人群id
	_ucrowdid int64
	// 更新方式，1:增量更新，2:全量更新
	_edittype int64
}

// NewTaobaotbkscucrowdmemberaddRequest 初始化TaobaotbkscucrowdmemberaddAPIRequest对象
func NewTaobaotbkscucrowdmemberaddRequest() *TaobaotbkscucrowdmemberaddAPIRequest {
	return &TaobaotbkscucrowdmemberaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkscucrowdmemberaddAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.ucrowd.member.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkscucrowdmemberaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkscucrowdmemberaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccountlist is Accountlist Setter
// 成员列表
func (r *TaobaotbkscucrowdmemberaddAPIRequest) SetAccountlist(_accountlist []string) error {
	r._accountlist = _accountlist
	r.Set("account_list", _accountlist)
	return nil
}

// GetAccountlist Accountlist Getter
func (r TaobaotbkscucrowdmemberaddAPIRequest) GetAccountlist() []string {
	return r._accountlist
}

// SetAccounttype is Accounttype Setter
// 成员账号类型，1:sid
func (r *TaobaotbkscucrowdmemberaddAPIRequest) SetAccounttype(_accounttype int64) error {
	r._accounttype = _accounttype
	r.Set("account_type", _accounttype)
	return nil
}

// GetAccounttype Accounttype Getter
func (r TaobaotbkscucrowdmemberaddAPIRequest) GetAccounttype() int64 {
	return r._accounttype
}

// SetUcrowdid is Ucrowdid Setter
// 人群id
func (r *TaobaotbkscucrowdmemberaddAPIRequest) SetUcrowdid(_ucrowdid int64) error {
	r._ucrowdid = _ucrowdid
	r.Set("ucrowd_id", _ucrowdid)
	return nil
}

// GetUcrowdid Ucrowdid Getter
func (r TaobaotbkscucrowdmemberaddAPIRequest) GetUcrowdid() int64 {
	return r._ucrowdid
}

// SetEdittype is Edittype Setter
// 更新方式，1:增量更新，2:全量更新
func (r *TaobaotbkscucrowdmemberaddAPIRequest) SetEdittype(_edittype int64) error {
	r._edittype = _edittype
	r.Set("edit_type", _edittype)
	return nil
}

// GetEdittype Edittype Getter
func (r TaobaotbkscucrowdmemberaddAPIRequest) GetEdittype() int64 {
	return r._edittype
}

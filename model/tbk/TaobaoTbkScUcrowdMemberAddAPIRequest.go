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
	_accountList []string
	// 成员账号类型，1:sid
	_accountType int64
	// 人群id
	_ucrowdId int64
	// 更新方式，1:增量更新，2:全量更新
	_editType int64
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

// SetAccountList is AccountList Setter
// 成员列表
func (r *TaobaotbkscucrowdmemberaddAPIRequest) SetAccountList(_accountList []string) error {
	r._accountList = _accountList
	r.Set("account_list", _accountList)
	return nil
}

// GetAccountList AccountList Getter
func (r TaobaotbkscucrowdmemberaddAPIRequest) GetAccountList() []string {
	return r._accountList
}

// SetAccountType is AccountType Setter
// 成员账号类型，1:sid
func (r *TaobaotbkscucrowdmemberaddAPIRequest) SetAccountType(_accountType int64) error {
	r._accountType = _accountType
	r.Set("account_type", _accountType)
	return nil
}

// GetAccountType AccountType Getter
func (r TaobaotbkscucrowdmemberaddAPIRequest) GetAccountType() int64 {
	return r._accountType
}

// SetUcrowdId is UcrowdId Setter
// 人群id
func (r *TaobaotbkscucrowdmemberaddAPIRequest) SetUcrowdId(_ucrowdId int64) error {
	r._ucrowdId = _ucrowdId
	r.Set("ucrowd_id", _ucrowdId)
	return nil
}

// GetUcrowdId UcrowdId Getter
func (r TaobaotbkscucrowdmemberaddAPIRequest) GetUcrowdId() int64 {
	return r._ucrowdId
}

// SetEditType is EditType Setter
// 更新方式，1:增量更新，2:全量更新
func (r *TaobaotbkscucrowdmemberaddAPIRequest) SetEditType(_editType int64) error {
	r._editType = _editType
	r.Set("edit_type", _editType)
	return nil
}

// GetEditType EditType Getter
func (r TaobaotbkscucrowdmemberaddAPIRequest) GetEditType() int64 {
	return r._editType
}

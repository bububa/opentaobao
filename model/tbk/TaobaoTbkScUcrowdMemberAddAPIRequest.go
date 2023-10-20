package tbk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScUcrowdMemberAddAPIRequest 淘宝客-服务商-上传人群数据 API请求
// taobao.tbk.sc.ucrowd.member.add
//
// 服务商使用。支持淘宝客上传人群标签id对应的会员列表，支持全量和增量多种数据更新方式。
type TaobaoTbkScUcrowdMemberAddAPIRequest struct {
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

// NewTaobaoTbkScUcrowdMemberAddRequest 初始化TaobaoTbkScUcrowdMemberAddAPIRequest对象
func NewTaobaoTbkScUcrowdMemberAddRequest() *TaobaoTbkScUcrowdMemberAddAPIRequest {
	return &TaobaoTbkScUcrowdMemberAddAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTbkScUcrowdMemberAddAPIRequest) Reset() {
	r._accountList = r._accountList[:0]
	r._accountType = 0
	r._ucrowdId = 0
	r._editType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkScUcrowdMemberAddAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.ucrowd.member.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkScUcrowdMemberAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkScUcrowdMemberAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccountList is AccountList Setter
// 成员列表
func (r *TaobaoTbkScUcrowdMemberAddAPIRequest) SetAccountList(_accountList []string) error {
	r._accountList = _accountList
	r.Set("account_list", _accountList)
	return nil
}

// GetAccountList AccountList Getter
func (r TaobaoTbkScUcrowdMemberAddAPIRequest) GetAccountList() []string {
	return r._accountList
}

// SetAccountType is AccountType Setter
// 成员账号类型，1:sid
func (r *TaobaoTbkScUcrowdMemberAddAPIRequest) SetAccountType(_accountType int64) error {
	r._accountType = _accountType
	r.Set("account_type", _accountType)
	return nil
}

// GetAccountType AccountType Getter
func (r TaobaoTbkScUcrowdMemberAddAPIRequest) GetAccountType() int64 {
	return r._accountType
}

// SetUcrowdId is UcrowdId Setter
// 人群id
func (r *TaobaoTbkScUcrowdMemberAddAPIRequest) SetUcrowdId(_ucrowdId int64) error {
	r._ucrowdId = _ucrowdId
	r.Set("ucrowd_id", _ucrowdId)
	return nil
}

// GetUcrowdId UcrowdId Getter
func (r TaobaoTbkScUcrowdMemberAddAPIRequest) GetUcrowdId() int64 {
	return r._ucrowdId
}

// SetEditType is EditType Setter
// 更新方式，1:增量更新，2:全量更新
func (r *TaobaoTbkScUcrowdMemberAddAPIRequest) SetEditType(_editType int64) error {
	r._editType = _editType
	r.Set("edit_type", _editType)
	return nil
}

// GetEditType EditType Getter
func (r TaobaoTbkScUcrowdMemberAddAPIRequest) GetEditType() int64 {
	return r._editType
}

var poolTaobaoTbkScUcrowdMemberAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTbkScUcrowdMemberAddRequest()
	},
}

// GetTaobaoTbkScUcrowdMemberAddRequest 从 sync.Pool 获取 TaobaoTbkScUcrowdMemberAddAPIRequest
func GetTaobaoTbkScUcrowdMemberAddAPIRequest() *TaobaoTbkScUcrowdMemberAddAPIRequest {
	return poolTaobaoTbkScUcrowdMemberAddAPIRequest.Get().(*TaobaoTbkScUcrowdMemberAddAPIRequest)
}

// ReleaseTaobaoTbkScUcrowdMemberAddAPIRequest 将 TaobaoTbkScUcrowdMemberAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoTbkScUcrowdMemberAddAPIRequest(v *TaobaoTbkScUcrowdMemberAddAPIRequest) {
	v.Reset()
	poolTaobaoTbkScUcrowdMemberAddAPIRequest.Put(v)
}

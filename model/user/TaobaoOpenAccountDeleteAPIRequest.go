package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenaccountdeleteAPIRequest OpenAccount删除数据 API请求
// taobao.open.account.delete
//
// OpenAccount删除数据
type TaobaoopenaccountdeleteAPIRequest struct {
	model.Params
	// Open Account的id列表
	_openAccountIds []int64
	// ISV自己账号的id列表
	_isvAccountIds []string
}

// NewTaobaoopenaccountdeleteRequest 初始化TaobaoopenaccountdeleteAPIRequest对象
func NewTaobaoopenaccountdeleteRequest() *TaobaoopenaccountdeleteAPIRequest {
	return &TaobaoopenaccountdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenaccountdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.open.account.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenaccountdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenaccountdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenAccountIds is OpenAccountIds Setter
// Open Account的id列表
func (r *TaobaoopenaccountdeleteAPIRequest) SetOpenAccountIds(_openAccountIds []int64) error {
	r._openAccountIds = _openAccountIds
	r.Set("open_account_ids", _openAccountIds)
	return nil
}

// GetOpenAccountIds OpenAccountIds Getter
func (r TaobaoopenaccountdeleteAPIRequest) GetOpenAccountIds() []int64 {
	return r._openAccountIds
}

// SetIsvAccountIds is IsvAccountIds Setter
// ISV自己账号的id列表
func (r *TaobaoopenaccountdeleteAPIRequest) SetIsvAccountIds(_isvAccountIds []string) error {
	r._isvAccountIds = _isvAccountIds
	r.Set("isv_account_ids", _isvAccountIds)
	return nil
}

// GetIsvAccountIds IsvAccountIds Getter
func (r TaobaoopenaccountdeleteAPIRequest) GetIsvAccountIds() []string {
	return r._isvAccountIds
}

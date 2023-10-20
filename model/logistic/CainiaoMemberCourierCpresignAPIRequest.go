package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaomembercouriercpresignAPIRequest cp清理离职用户信息 API请求
// cainiao.member.courier.cpresign
//
// CP清理内部离职的用户信息
type CainiaomembercouriercpresignAPIRequest struct {
	model.Params
	// 菜鸟用户id
	_accountId int64
}

// NewCainiaomembercouriercpresignRequest 初始化CainiaomembercouriercpresignAPIRequest对象
func NewCainiaomembercouriercpresignRequest() *CainiaomembercouriercpresignAPIRequest {
	return &CainiaomembercouriercpresignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaomembercouriercpresignAPIRequest) GetApiMethodName() string {
	return "cainiao.member.courier.cpresign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaomembercouriercpresignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaomembercouriercpresignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccountId is AccountId Setter
// 菜鸟用户id
func (r *CainiaomembercouriercpresignAPIRequest) SetAccountId(_accountId int64) error {
	r._accountId = _accountId
	r.Set("account_id", _accountId)
	return nil
}

// GetAccountId AccountId Getter
func (r CainiaomembercouriercpresignAPIRequest) GetAccountId() int64 {
	return r._accountId
}

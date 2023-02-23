package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoMemberCourierCpresignAPIRequest cp清理离职用户信息 API请求
// cainiao.member.courier.cpresign
//
// CP清理内部离职的用户信息
type CainiaoMemberCourierCpresignAPIRequest struct {
	model.Params
	// 菜鸟用户id
	_accountId int64
}

// NewCainiaoMemberCourierCpresignRequest 初始化CainiaoMemberCourierCpresignAPIRequest对象
func NewCainiaoMemberCourierCpresignRequest() *CainiaoMemberCourierCpresignAPIRequest {
	return &CainiaoMemberCourierCpresignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoMemberCourierCpresignAPIRequest) GetApiMethodName() string {
	return "cainiao.member.courier.cpresign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoMemberCourierCpresignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoMemberCourierCpresignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccountId is AccountId Setter
// 菜鸟用户id
func (r *CainiaoMemberCourierCpresignAPIRequest) SetAccountId(_accountId int64) error {
	r._accountId = _accountId
	r.Set("account_id", _accountId)
	return nil
}

// GetAccountId AccountId Getter
func (r CainiaoMemberCourierCpresignAPIRequest) GetAccountId() int64 {
	return r._accountId
}

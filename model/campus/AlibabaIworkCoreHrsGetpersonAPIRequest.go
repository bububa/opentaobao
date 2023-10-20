package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaiworkcorehrsgetpersonAPIRequest 获取神鲸用户基本信息 API请求
// alibaba.iwork.core.hrs.getperson
//
// 神鲸用户的基本信息查询，根据PERSON_ID或者用户ACCOUNT_ID查询
type AlibabaiworkcorehrsgetpersonAPIRequest struct {
	model.Params
	// 用户ACCOUNT_ID
	_accountId string
	// 应用ID
	_appId string
	// 操作人ID
	_operatorId string
	// 用户ID
	_personId int64
}

// NewAlibabaiworkcorehrsgetpersonRequest 初始化AlibabaiworkcorehrsgetpersonAPIRequest对象
func NewAlibabaiworkcorehrsgetpersonRequest() *AlibabaiworkcorehrsgetpersonAPIRequest {
	return &AlibabaiworkcorehrsgetpersonAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaiworkcorehrsgetpersonAPIRequest) GetApiMethodName() string {
	return "alibaba.iwork.core.hrs.getperson"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaiworkcorehrsgetpersonAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaiworkcorehrsgetpersonAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccountId is AccountId Setter
// 用户ACCOUNT_ID
func (r *AlibabaiworkcorehrsgetpersonAPIRequest) SetAccountId(_accountId string) error {
	r._accountId = _accountId
	r.Set("account_id", _accountId)
	return nil
}

// GetAccountId AccountId Getter
func (r AlibabaiworkcorehrsgetpersonAPIRequest) GetAccountId() string {
	return r._accountId
}

// SetAppId is AppId Setter
// 应用ID
func (r *AlibabaiworkcorehrsgetpersonAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r AlibabaiworkcorehrsgetpersonAPIRequest) GetAppId() string {
	return r._appId
}

// SetOperatorId is OperatorId Setter
// 操作人ID
func (r *AlibabaiworkcorehrsgetpersonAPIRequest) SetOperatorId(_operatorId string) error {
	r._operatorId = _operatorId
	r.Set("operator_id", _operatorId)
	return nil
}

// GetOperatorId OperatorId Getter
func (r AlibabaiworkcorehrsgetpersonAPIRequest) GetOperatorId() string {
	return r._operatorId
}

// SetPersonId is PersonId Setter
// 用户ID
func (r *AlibabaiworkcorehrsgetpersonAPIRequest) SetPersonId(_personId int64) error {
	r._personId = _personId
	r.Set("person_id", _personId)
	return nil
}

// GetPersonId PersonId Getter
func (r AlibabaiworkcorehrsgetpersonAPIRequest) GetPersonId() int64 {
	return r._personId
}

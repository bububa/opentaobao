package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIworkCoreHrsGetpersonAPIRequest 获取神鲸用户基本信息 API请求
// alibaba.iwork.core.hrs.getperson
//
// 神鲸用户的基本信息查询，根据PERSON_ID或者用户ACCOUNT_ID查询
type AlibabaIworkCoreHrsGetpersonAPIRequest struct {
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

// NewAlibabaIworkCoreHrsGetpersonRequest 初始化AlibabaIworkCoreHrsGetpersonAPIRequest对象
func NewAlibabaIworkCoreHrsGetpersonRequest() *AlibabaIworkCoreHrsGetpersonAPIRequest {
	return &AlibabaIworkCoreHrsGetpersonAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIworkCoreHrsGetpersonAPIRequest) Reset() {
	r._accountId = ""
	r._appId = ""
	r._operatorId = ""
	r._personId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIworkCoreHrsGetpersonAPIRequest) GetApiMethodName() string {
	return "alibaba.iwork.core.hrs.getperson"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIworkCoreHrsGetpersonAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIworkCoreHrsGetpersonAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccountId is AccountId Setter
// 用户ACCOUNT_ID
func (r *AlibabaIworkCoreHrsGetpersonAPIRequest) SetAccountId(_accountId string) error {
	r._accountId = _accountId
	r.Set("account_id", _accountId)
	return nil
}

// GetAccountId AccountId Getter
func (r AlibabaIworkCoreHrsGetpersonAPIRequest) GetAccountId() string {
	return r._accountId
}

// SetAppId is AppId Setter
// 应用ID
func (r *AlibabaIworkCoreHrsGetpersonAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r AlibabaIworkCoreHrsGetpersonAPIRequest) GetAppId() string {
	return r._appId
}

// SetOperatorId is OperatorId Setter
// 操作人ID
func (r *AlibabaIworkCoreHrsGetpersonAPIRequest) SetOperatorId(_operatorId string) error {
	r._operatorId = _operatorId
	r.Set("operator_id", _operatorId)
	return nil
}

// GetOperatorId OperatorId Getter
func (r AlibabaIworkCoreHrsGetpersonAPIRequest) GetOperatorId() string {
	return r._operatorId
}

// SetPersonId is PersonId Setter
// 用户ID
func (r *AlibabaIworkCoreHrsGetpersonAPIRequest) SetPersonId(_personId int64) error {
	r._personId = _personId
	r.Set("person_id", _personId)
	return nil
}

// GetPersonId PersonId Getter
func (r AlibabaIworkCoreHrsGetpersonAPIRequest) GetPersonId() int64 {
	return r._personId
}

var poolAlibabaIworkCoreHrsGetpersonAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIworkCoreHrsGetpersonRequest()
	},
}

// GetAlibabaIworkCoreHrsGetpersonRequest 从 sync.Pool 获取 AlibabaIworkCoreHrsGetpersonAPIRequest
func GetAlibabaIworkCoreHrsGetpersonAPIRequest() *AlibabaIworkCoreHrsGetpersonAPIRequest {
	return poolAlibabaIworkCoreHrsGetpersonAPIRequest.Get().(*AlibabaIworkCoreHrsGetpersonAPIRequest)
}

// ReleaseAlibabaIworkCoreHrsGetpersonAPIRequest 将 AlibabaIworkCoreHrsGetpersonAPIRequest 放入 sync.Pool
func ReleaseAlibabaIworkCoreHrsGetpersonAPIRequest(v *AlibabaIworkCoreHrsGetpersonAPIRequest) {
	v.Reset()
	poolAlibabaIworkCoreHrsGetpersonAPIRequest.Put(v)
}

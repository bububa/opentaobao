package fundplatform

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFundplatformAccountQueryInfoAPIRequest 查询账户信息 API请求
// alibaba.fundplatform.account.query.info
//
// 外部查询资金平台用户账户信息
type AlibabaFundplatformAccountQueryInfoAPIRequest struct {
	model.Params
	// 账户ID
	_accountId int64
}

// NewAlibabaFundplatformAccountQueryInfoRequest 初始化AlibabaFundplatformAccountQueryInfoAPIRequest对象
func NewAlibabaFundplatformAccountQueryInfoRequest() *AlibabaFundplatformAccountQueryInfoAPIRequest {
	return &AlibabaFundplatformAccountQueryInfoAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaFundplatformAccountQueryInfoAPIRequest) Reset() {
	r._accountId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformAccountQueryInfoAPIRequest) GetApiMethodName() string {
	return "alibaba.fundplatform.account.query.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformAccountQueryInfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaFundplatformAccountQueryInfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccountId is AccountId Setter
// 账户ID
func (r *AlibabaFundplatformAccountQueryInfoAPIRequest) SetAccountId(_accountId int64) error {
	r._accountId = _accountId
	r.Set("account_id", _accountId)
	return nil
}

// GetAccountId AccountId Getter
func (r AlibabaFundplatformAccountQueryInfoAPIRequest) GetAccountId() int64 {
	return r._accountId
}

var poolAlibabaFundplatformAccountQueryInfoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaFundplatformAccountQueryInfoRequest()
	},
}

// GetAlibabaFundplatformAccountQueryInfoRequest 从 sync.Pool 获取 AlibabaFundplatformAccountQueryInfoAPIRequest
func GetAlibabaFundplatformAccountQueryInfoAPIRequest() *AlibabaFundplatformAccountQueryInfoAPIRequest {
	return poolAlibabaFundplatformAccountQueryInfoAPIRequest.Get().(*AlibabaFundplatformAccountQueryInfoAPIRequest)
}

// ReleaseAlibabaFundplatformAccountQueryInfoAPIRequest 将 AlibabaFundplatformAccountQueryInfoAPIRequest 放入 sync.Pool
func ReleaseAlibabaFundplatformAccountQueryInfoAPIRequest(v *AlibabaFundplatformAccountQueryInfoAPIRequest) {
	v.Reset()
	poolAlibabaFundplatformAccountQueryInfoAPIRequest.Put(v)
}

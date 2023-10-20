package charity

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaValueCoinIssueAPIRequest 爱豆发放 API请求
// alibaba.value.coin.issue
//
// 爱豆发放
type AlibabaValueCoinIssueAPIRequest struct {
	model.Params
	// 爱豆发放参数
	_exCoinIssueParam *ExCoinIssueParam
}

// NewAlibabaValueCoinIssueRequest 初始化AlibabaValueCoinIssueAPIRequest对象
func NewAlibabaValueCoinIssueRequest() *AlibabaValueCoinIssueAPIRequest {
	return &AlibabaValueCoinIssueAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaValueCoinIssueAPIRequest) Reset() {
	r._exCoinIssueParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaValueCoinIssueAPIRequest) GetApiMethodName() string {
	return "alibaba.value.coin.issue"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaValueCoinIssueAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaValueCoinIssueAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExCoinIssueParam is ExCoinIssueParam Setter
// 爱豆发放参数
func (r *AlibabaValueCoinIssueAPIRequest) SetExCoinIssueParam(_exCoinIssueParam *ExCoinIssueParam) error {
	r._exCoinIssueParam = _exCoinIssueParam
	r.Set("ex_coin_issue_param", _exCoinIssueParam)
	return nil
}

// GetExCoinIssueParam ExCoinIssueParam Getter
func (r AlibabaValueCoinIssueAPIRequest) GetExCoinIssueParam() *ExCoinIssueParam {
	return r._exCoinIssueParam
}

var poolAlibabaValueCoinIssueAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaValueCoinIssueRequest()
	},
}

// GetAlibabaValueCoinIssueRequest 从 sync.Pool 获取 AlibabaValueCoinIssueAPIRequest
func GetAlibabaValueCoinIssueAPIRequest() *AlibabaValueCoinIssueAPIRequest {
	return poolAlibabaValueCoinIssueAPIRequest.Get().(*AlibabaValueCoinIssueAPIRequest)
}

// ReleaseAlibabaValueCoinIssueAPIRequest 将 AlibabaValueCoinIssueAPIRequest 放入 sync.Pool
func ReleaseAlibabaValueCoinIssueAPIRequest(v *AlibabaValueCoinIssueAPIRequest) {
	v.Reset()
	poolAlibabaValueCoinIssueAPIRequest.Put(v)
}

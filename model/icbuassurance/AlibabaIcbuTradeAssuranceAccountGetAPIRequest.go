package icbuassurance

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuTradeAssuranceAccountGetAPIRequest icbu信保账户信息 API请求
// alibaba.icbu.trade.assurance.account.get
//
// icbu交易信用保障开通状态&amp;额度信息查询
type AlibabaIcbuTradeAssuranceAccountGetAPIRequest struct {
	model.Params
}

// NewAlibabaIcbuTradeAssuranceAccountGetRequest 初始化AlibabaIcbuTradeAssuranceAccountGetAPIRequest对象
func NewAlibabaIcbuTradeAssuranceAccountGetRequest() *AlibabaIcbuTradeAssuranceAccountGetAPIRequest {
	return &AlibabaIcbuTradeAssuranceAccountGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIcbuTradeAssuranceAccountGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuTradeAssuranceAccountGetAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.trade.assurance.account.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuTradeAssuranceAccountGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIcbuTradeAssuranceAccountGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaIcbuTradeAssuranceAccountGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIcbuTradeAssuranceAccountGetRequest()
	},
}

// GetAlibabaIcbuTradeAssuranceAccountGetRequest 从 sync.Pool 获取 AlibabaIcbuTradeAssuranceAccountGetAPIRequest
func GetAlibabaIcbuTradeAssuranceAccountGetAPIRequest() *AlibabaIcbuTradeAssuranceAccountGetAPIRequest {
	return poolAlibabaIcbuTradeAssuranceAccountGetAPIRequest.Get().(*AlibabaIcbuTradeAssuranceAccountGetAPIRequest)
}

// ReleaseAlibabaIcbuTradeAssuranceAccountGetAPIRequest 将 AlibabaIcbuTradeAssuranceAccountGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaIcbuTradeAssuranceAccountGetAPIRequest(v *AlibabaIcbuTradeAssuranceAccountGetAPIRequest) {
	v.Reset()
	poolAlibabaIcbuTradeAssuranceAccountGetAPIRequest.Put(v)
}

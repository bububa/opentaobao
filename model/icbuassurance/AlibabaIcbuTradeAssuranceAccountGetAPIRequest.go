package icbuassurance

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuTradeAssuranceAccountGetAPIRequest icbu信保账户信息 API请求
// alibaba.icbu.trade.assurance.account.get
//
// icbu交易信用保障开通状态&额度信息查询
type AlibabaIcbuTradeAssuranceAccountGetAPIRequest struct {
	model.Params
}

// NewAlibabaIcbuTradeAssuranceAccountGetRequest 初始化AlibabaIcbuTradeAssuranceAccountGetAPIRequest对象
func NewAlibabaIcbuTradeAssuranceAccountGetRequest() *AlibabaIcbuTradeAssuranceAccountGetAPIRequest {
	return &AlibabaIcbuTradeAssuranceAccountGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuTradeAssuranceAccountGetAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.trade.assurance.account.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuTradeAssuranceAccountGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

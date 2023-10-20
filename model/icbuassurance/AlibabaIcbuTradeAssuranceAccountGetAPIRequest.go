package icbuassurance

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbutradeassuranceaccountgetAPIRequest icbu信保账户信息 API请求
// alibaba.icbu.trade.assurance.account.get
//
// icbu交易信用保障开通状态&amp;额度信息查询
type AlibabaicbutradeassuranceaccountgetAPIRequest struct {
	model.Params
}

// NewAlibabaicbutradeassuranceaccountgetRequest 初始化AlibabaicbutradeassuranceaccountgetAPIRequest对象
func NewAlibabaicbutradeassuranceaccountgetRequest() *AlibabaicbutradeassuranceaccountgetAPIRequest {
	return &AlibabaicbutradeassuranceaccountgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicbutradeassuranceaccountgetAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.trade.assurance.account.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicbutradeassuranceaccountgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicbutradeassuranceaccountgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

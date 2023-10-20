package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadaccountbalancegetAPIRequest 查询账户余额 API请求
// alibaba.scbp.ad.account.balance.get
//
// 查询推广账户余额
type AlibabascbpadaccountbalancegetAPIRequest struct {
	model.Params
}

// NewAlibabascbpadaccountbalancegetRequest 初始化AlibabascbpadaccountbalancegetAPIRequest对象
func NewAlibabascbpadaccountbalancegetRequest() *AlibabascbpadaccountbalancegetAPIRequest {
	return &AlibabascbpadaccountbalancegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadaccountbalancegetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.account.balance.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadaccountbalancegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadaccountbalancegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

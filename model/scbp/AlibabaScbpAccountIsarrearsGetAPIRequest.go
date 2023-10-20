package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpaccountisarrearsgetAPIRequest 查询关键词推广账户是否欠款 API请求
// alibaba.scbp.account.isarrears.get
//
// 查询关键词推广账户是否欠款
type AlibabascbpaccountisarrearsgetAPIRequest struct {
	model.Params
}

// NewAlibabascbpaccountisarrearsgetRequest 初始化AlibabascbpaccountisarrearsgetAPIRequest对象
func NewAlibabascbpaccountisarrearsgetRequest() *AlibabascbpaccountisarrearsgetAPIRequest {
	return &AlibabascbpaccountisarrearsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpaccountisarrearsgetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.account.isarrears.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpaccountisarrearsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpaccountisarrearsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

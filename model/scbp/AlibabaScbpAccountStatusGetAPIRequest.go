package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpaccountstatusgetAPIRequest 查询账户级别关键词推广状态 API请求
// alibaba.scbp.account.status.get
//
// 查询账户级别关键词推广状态
type AlibabascbpaccountstatusgetAPIRequest struct {
	model.Params
}

// NewAlibabascbpaccountstatusgetRequest 初始化AlibabascbpaccountstatusgetAPIRequest对象
func NewAlibabascbpaccountstatusgetRequest() *AlibabascbpaccountstatusgetAPIRequest {
	return &AlibabascbpaccountstatusgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpaccountstatusgetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.account.status.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpaccountstatusgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpaccountstatusgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

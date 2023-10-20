package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpeffectaccountdategetAPIRequest 获取最近报表生成时间 API请求
// alibaba.scbp.effect.account.date.get
//
// 获取最近报表生成时间,格式为yyyy-MM-dd
type AlibabascbpeffectaccountdategetAPIRequest struct {
	model.Params
}

// NewAlibabascbpeffectaccountdategetRequest 初始化AlibabascbpeffectaccountdategetAPIRequest对象
func NewAlibabascbpeffectaccountdategetRequest() *AlibabascbpeffectaccountdategetAPIRequest {
	return &AlibabascbpeffectaccountdategetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpeffectaccountdategetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.effect.account.date.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpeffectaccountdategetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpeffectaccountdategetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

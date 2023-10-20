package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpaccountdaycostgetAPIRequest 查询今日消耗 API请求
// alibaba.scbp.account.daycost.get
//
// 查询今日消耗
type AlibabascbpaccountdaycostgetAPIRequest struct {
	model.Params
}

// NewAlibabascbpaccountdaycostgetRequest 初始化AlibabascbpaccountdaycostgetAPIRequest对象
func NewAlibabascbpaccountdaycostgetRequest() *AlibabascbpaccountdaycostgetAPIRequest {
	return &AlibabascbpaccountdaycostgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpaccountdaycostgetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.account.daycost.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpaccountdaycostgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpaccountdaycostgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

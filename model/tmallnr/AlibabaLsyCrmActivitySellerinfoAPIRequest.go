package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalsycrmactivitysellerinfoAPIRequest 商家信息推送 API请求
// alibaba.lsy.crm.activity.sellerinfo
//
// 本地团商家信息推送
type AlibabalsycrmactivitysellerinfoAPIRequest struct {
	model.Params
}

// NewAlibabalsycrmactivitysellerinfoRequest 初始化AlibabalsycrmactivitysellerinfoAPIRequest对象
func NewAlibabalsycrmactivitysellerinfoRequest() *AlibabalsycrmactivitysellerinfoAPIRequest {
	return &AlibabalsycrmactivitysellerinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalsycrmactivitysellerinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.lsy.crm.activity.sellerinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalsycrmactivitysellerinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalsycrmactivitysellerinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

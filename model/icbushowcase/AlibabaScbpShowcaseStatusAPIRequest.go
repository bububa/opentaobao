package icbushowcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpshowcasestatusAPIRequest 橱窗状态 API请求
// alibaba.scbp.showcase.status
//
// 查询橱窗状态，如总数、可用数量
type AlibabascbpshowcasestatusAPIRequest struct {
	model.Params
}

// NewAlibabascbpshowcasestatusRequest 初始化AlibabascbpshowcasestatusAPIRequest对象
func NewAlibabascbpshowcasestatusRequest() *AlibabascbpshowcasestatusAPIRequest {
	return &AlibabascbpshowcasestatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpshowcasestatusAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.showcase.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpshowcasestatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpshowcasestatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

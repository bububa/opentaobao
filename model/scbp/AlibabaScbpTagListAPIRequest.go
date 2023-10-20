package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbptaglistAPIRequest 查询所有分组 API请求
// alibaba.scbp.tag.list
//
// 查询所有分组
type AlibabascbptaglistAPIRequest struct {
	model.Params
}

// NewAlibabascbptaglistRequest 初始化AlibabascbptaglistAPIRequest对象
func NewAlibabascbptaglistRequest() *AlibabascbptaglistAPIRequest {
	return &AlibabascbptaglistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbptaglistAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.tag.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbptaglistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbptaglistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

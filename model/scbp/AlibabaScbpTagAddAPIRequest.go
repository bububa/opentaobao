package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbptagaddAPIRequest 创建关键词分组 API请求
// alibaba.scbp.tag.add
//
// 创建关键词分组
type AlibabascbptagaddAPIRequest struct {
	model.Params
	// 分组名称，最多允许创建100个
	_tagName string
}

// NewAlibabascbptagaddRequest 初始化AlibabascbptagaddAPIRequest对象
func NewAlibabascbptagaddRequest() *AlibabascbptagaddAPIRequest {
	return &AlibabascbptagaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbptagaddAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.tag.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbptagaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbptagaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTagName is TagName Setter
// 分组名称，最多允许创建100个
func (r *AlibabascbptagaddAPIRequest) SetTagName(_tagName string) error {
	r._tagName = _tagName
	r.Set("tag_name", _tagName)
	return nil
}

// GetTagName TagName Getter
func (r AlibabascbptagaddAPIRequest) GetTagName() string {
	return r._tagName
}

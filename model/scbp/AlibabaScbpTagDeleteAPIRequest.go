package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbptagdeleteAPIRequest 删除关键词分组 API请求
// alibaba.scbp.tag.delete
//
// 删除关键词分组
type AlibabascbptagdeleteAPIRequest struct {
	model.Params
	// 关键词分组名
	_tagName string
}

// NewAlibabascbptagdeleteRequest 初始化AlibabascbptagdeleteAPIRequest对象
func NewAlibabascbptagdeleteRequest() *AlibabascbptagdeleteAPIRequest {
	return &AlibabascbptagdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbptagdeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.tag.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbptagdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbptagdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTagName is TagName Setter
// 关键词分组名
func (r *AlibabascbptagdeleteAPIRequest) SetTagName(_tagName string) error {
	r._tagName = _tagName
	r.Set("tag_name", _tagName)
	return nil
}

// GetTagName TagName Getter
func (r AlibabascbptagdeleteAPIRequest) GetTagName() string {
	return r._tagName
}

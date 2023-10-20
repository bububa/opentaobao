package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbptagrenameAPIRequest 重命名关键词分组 API请求
// alibaba.scbp.tag.rename
//
// 重命名关键词分组
type AlibabascbptagrenameAPIRequest struct {
	model.Params
	// 需要重命名的关键词分组名
	_tagName string
	// 新分组名
	_newTagName string
}

// NewAlibabascbptagrenameRequest 初始化AlibabascbptagrenameAPIRequest对象
func NewAlibabascbptagrenameRequest() *AlibabascbptagrenameAPIRequest {
	return &AlibabascbptagrenameAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbptagrenameAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.tag.rename"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbptagrenameAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbptagrenameAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTagName is TagName Setter
// 需要重命名的关键词分组名
func (r *AlibabascbptagrenameAPIRequest) SetTagName(_tagName string) error {
	r._tagName = _tagName
	r.Set("tag_name", _tagName)
	return nil
}

// GetTagName TagName Getter
func (r AlibabascbptagrenameAPIRequest) GetTagName() string {
	return r._tagName
}

// SetNewTagName is NewTagName Setter
// 新分组名
func (r *AlibabascbptagrenameAPIRequest) SetNewTagName(_newTagName string) error {
	r._newTagName = _newTagName
	r.Set("new_tag_name", _newTagName)
	return nil
}

// GetNewTagName NewTagName Getter
func (r AlibabascbptagrenameAPIRequest) GetNewTagName() string {
	return r._newTagName
}

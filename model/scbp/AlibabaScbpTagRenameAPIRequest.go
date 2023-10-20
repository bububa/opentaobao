package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTagRenameAPIRequest 重命名关键词分组 API请求
// alibaba.scbp.tag.rename
//
// 重命名关键词分组
type AlibabaScbpTagRenameAPIRequest struct {
	model.Params
	// 需要重命名的关键词分组名
	_tagName string
	// 新分组名
	_newTagName string
}

// NewAlibabaScbpTagRenameRequest 初始化AlibabaScbpTagRenameAPIRequest对象
func NewAlibabaScbpTagRenameRequest() *AlibabaScbpTagRenameAPIRequest {
	return &AlibabaScbpTagRenameAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpTagRenameAPIRequest) Reset() {
	r._tagName = ""
	r._newTagName = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpTagRenameAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.tag.rename"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpTagRenameAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpTagRenameAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTagName is TagName Setter
// 需要重命名的关键词分组名
func (r *AlibabaScbpTagRenameAPIRequest) SetTagName(_tagName string) error {
	r._tagName = _tagName
	r.Set("tag_name", _tagName)
	return nil
}

// GetTagName TagName Getter
func (r AlibabaScbpTagRenameAPIRequest) GetTagName() string {
	return r._tagName
}

// SetNewTagName is NewTagName Setter
// 新分组名
func (r *AlibabaScbpTagRenameAPIRequest) SetNewTagName(_newTagName string) error {
	r._newTagName = _newTagName
	r.Set("new_tag_name", _newTagName)
	return nil
}

// GetNewTagName NewTagName Getter
func (r AlibabaScbpTagRenameAPIRequest) GetNewTagName() string {
	return r._newTagName
}

var poolAlibabaScbpTagRenameAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpTagRenameRequest()
	},
}

// GetAlibabaScbpTagRenameRequest 从 sync.Pool 获取 AlibabaScbpTagRenameAPIRequest
func GetAlibabaScbpTagRenameAPIRequest() *AlibabaScbpTagRenameAPIRequest {
	return poolAlibabaScbpTagRenameAPIRequest.Get().(*AlibabaScbpTagRenameAPIRequest)
}

// ReleaseAlibabaScbpTagRenameAPIRequest 将 AlibabaScbpTagRenameAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpTagRenameAPIRequest(v *AlibabaScbpTagRenameAPIRequest) {
	v.Reset()
	poolAlibabaScbpTagRenameAPIRequest.Put(v)
}

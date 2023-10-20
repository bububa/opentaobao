package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTagDeleteAPIRequest 删除关键词分组 API请求
// alibaba.scbp.tag.delete
//
// 删除关键词分组
type AlibabaScbpTagDeleteAPIRequest struct {
	model.Params
	// 关键词分组名
	_tagName string
}

// NewAlibabaScbpTagDeleteRequest 初始化AlibabaScbpTagDeleteAPIRequest对象
func NewAlibabaScbpTagDeleteRequest() *AlibabaScbpTagDeleteAPIRequest {
	return &AlibabaScbpTagDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpTagDeleteAPIRequest) Reset() {
	r._tagName = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpTagDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.tag.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpTagDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpTagDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTagName is TagName Setter
// 关键词分组名
func (r *AlibabaScbpTagDeleteAPIRequest) SetTagName(_tagName string) error {
	r._tagName = _tagName
	r.Set("tag_name", _tagName)
	return nil
}

// GetTagName TagName Getter
func (r AlibabaScbpTagDeleteAPIRequest) GetTagName() string {
	return r._tagName
}

var poolAlibabaScbpTagDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpTagDeleteRequest()
	},
}

// GetAlibabaScbpTagDeleteRequest 从 sync.Pool 获取 AlibabaScbpTagDeleteAPIRequest
func GetAlibabaScbpTagDeleteAPIRequest() *AlibabaScbpTagDeleteAPIRequest {
	return poolAlibabaScbpTagDeleteAPIRequest.Get().(*AlibabaScbpTagDeleteAPIRequest)
}

// ReleaseAlibabaScbpTagDeleteAPIRequest 将 AlibabaScbpTagDeleteAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpTagDeleteAPIRequest(v *AlibabaScbpTagDeleteAPIRequest) {
	v.Reset()
	poolAlibabaScbpTagDeleteAPIRequest.Put(v)
}

package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTagAddAPIRequest 创建关键词分组 API请求
// alibaba.scbp.tag.add
//
// 创建关键词分组
type AlibabaScbpTagAddAPIRequest struct {
	model.Params
	// 分组名称，最多允许创建100个
	_tagName string
}

// NewAlibabaScbpTagAddRequest 初始化AlibabaScbpTagAddAPIRequest对象
func NewAlibabaScbpTagAddRequest() *AlibabaScbpTagAddAPIRequest {
	return &AlibabaScbpTagAddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpTagAddAPIRequest) Reset() {
	r._tagName = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpTagAddAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.tag.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpTagAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpTagAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTagName is TagName Setter
// 分组名称，最多允许创建100个
func (r *AlibabaScbpTagAddAPIRequest) SetTagName(_tagName string) error {
	r._tagName = _tagName
	r.Set("tag_name", _tagName)
	return nil
}

// GetTagName TagName Getter
func (r AlibabaScbpTagAddAPIRequest) GetTagName() string {
	return r._tagName
}

var poolAlibabaScbpTagAddAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpTagAddRequest()
	},
}

// GetAlibabaScbpTagAddRequest 从 sync.Pool 获取 AlibabaScbpTagAddAPIRequest
func GetAlibabaScbpTagAddAPIRequest() *AlibabaScbpTagAddAPIRequest {
	return poolAlibabaScbpTagAddAPIRequest.Get().(*AlibabaScbpTagAddAPIRequest)
}

// ReleaseAlibabaScbpTagAddAPIRequest 将 AlibabaScbpTagAddAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpTagAddAPIRequest(v *AlibabaScbpTagAddAPIRequest) {
	v.Reset()
	poolAlibabaScbpTagAddAPIRequest.Put(v)
}

package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkChannelCommentCreateAPIRequest 差评导入 API请求
// alibaba.wdk.channel.comment.create
//
// 差评导入
type AlibabaWdkChannelCommentCreateAPIRequest struct {
	model.Params
	// 差评信息
	_commentCreateInfo *CommentCreateInfo
}

// NewAlibabaWdkChannelCommentCreateRequest 初始化AlibabaWdkChannelCommentCreateAPIRequest对象
func NewAlibabaWdkChannelCommentCreateRequest() *AlibabaWdkChannelCommentCreateAPIRequest {
	return &AlibabaWdkChannelCommentCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkChannelCommentCreateAPIRequest) Reset() {
	r._commentCreateInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkChannelCommentCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.channel.comment.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkChannelCommentCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkChannelCommentCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCommentCreateInfo is CommentCreateInfo Setter
// 差评信息
func (r *AlibabaWdkChannelCommentCreateAPIRequest) SetCommentCreateInfo(_commentCreateInfo *CommentCreateInfo) error {
	r._commentCreateInfo = _commentCreateInfo
	r.Set("comment_create_info", _commentCreateInfo)
	return nil
}

// GetCommentCreateInfo CommentCreateInfo Getter
func (r AlibabaWdkChannelCommentCreateAPIRequest) GetCommentCreateInfo() *CommentCreateInfo {
	return r._commentCreateInfo
}

var poolAlibabaWdkChannelCommentCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkChannelCommentCreateRequest()
	},
}

// GetAlibabaWdkChannelCommentCreateRequest 从 sync.Pool 获取 AlibabaWdkChannelCommentCreateAPIRequest
func GetAlibabaWdkChannelCommentCreateAPIRequest() *AlibabaWdkChannelCommentCreateAPIRequest {
	return poolAlibabaWdkChannelCommentCreateAPIRequest.Get().(*AlibabaWdkChannelCommentCreateAPIRequest)
}

// ReleaseAlibabaWdkChannelCommentCreateAPIRequest 将 AlibabaWdkChannelCommentCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkChannelCommentCreateAPIRequest(v *AlibabaWdkChannelCommentCreateAPIRequest) {
	v.Reset()
	poolAlibabaWdkChannelCommentCreateAPIRequest.Put(v)
}

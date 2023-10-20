package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkchannelcommentcreateAPIRequest 差评导入 API请求
// alibaba.wdk.channel.comment.create
//
// 差评导入
type AlibabawdkchannelcommentcreateAPIRequest struct {
	model.Params
	// 差评信息
	_commentCreateInfo *CommentCreateInfo
}

// NewAlibabawdkchannelcommentcreateRequest 初始化AlibabawdkchannelcommentcreateAPIRequest对象
func NewAlibabawdkchannelcommentcreateRequest() *AlibabawdkchannelcommentcreateAPIRequest {
	return &AlibabawdkchannelcommentcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkchannelcommentcreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.channel.comment.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkchannelcommentcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkchannelcommentcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCommentCreateInfo is CommentCreateInfo Setter
// 差评信息
func (r *AlibabawdkchannelcommentcreateAPIRequest) SetCommentCreateInfo(_commentCreateInfo *CommentCreateInfo) error {
	r._commentCreateInfo = _commentCreateInfo
	r.Set("comment_create_info", _commentCreateInfo)
	return nil
}

// GetCommentCreateInfo CommentCreateInfo Getter
func (r AlibabawdkchannelcommentcreateAPIRequest) GetCommentCreateInfo() *CommentCreateInfo {
	return r._commentCreateInfo
}
